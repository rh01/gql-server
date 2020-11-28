package dispatch

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
)

// Methods
const (
	GET  = "GET"
	POST = "POST"
	PUT  = "PUT"
)

// Content type
const (
	JSON       = "application/json"
	URLENCODED = "application/x-www-form-urlencoded"
)

var (
	pathParamIndicator = ":"
	urlSeparator       = "/"

	errPathParamConflift = errors.New("Path parameter placeholder and actual vaues doesn't match")
	errMethodUnknown     = errors.New("Unknown request METHOD")
	errNoLoggerSet       = errors.New("Ensure the logger is set.")
	errCannotSetBody     = "Cannot set body for %s request"
)

// RequestMaker ...
type RequestMaker struct {
	method              string
	url                 string
	headers             map[string]string
	pathParams          []string
	queryParams         map[string]string
	unescapeQueryParams bool
	body                interface{}
	response            interface{}
	username            string
	password            string
	contentType         string
	logger              *logrus.Entry
	verbose             bool
}

// NewGetRequestMaker ..
// Create new request maker of GET request
func NewGetRequestMaker(url string) *RequestMaker {

	return &RequestMaker{
		method:      GET,
		url:         url,
		headers:     make(map[string]string),
		queryParams: make(map[string]string),
		pathParams:  make([]string, 0),
	}
}

// NewPostRequestMaker ..
// Create new request maker of GET request
func NewPostRequestMaker(url string) *RequestMaker {

	return &RequestMaker{
		method:      POST,
		url:         url,
		headers:     make(map[string]string),
		queryParams: make(map[string]string),
		pathParams:  make([]string, 0),
	}
}

// PathParams ..
// Inject params to the url path described with :
// Ex: http://elasticshift.com/api/users/:name
func (r *RequestMaker) PathParams(params ...string) *RequestMaker {
	r.pathParams = params
	return r
}

// QueryParam ..
// Set a query paramter to a request
func (r *RequestMaker) QueryParam(key, value string) *RequestMaker {
	r.queryParams[key] = value
	return r
}

//EncodeQueryParams
func (r *RequestMaker) UnescapeQueryParams(unescape bool) *RequestMaker {
	r.unescapeQueryParams = unescape
	return r
}

// Header ..
// Set a header value to a request
func (r *RequestMaker) Header(key, value string) *RequestMaker {
	r.headers[key] = value
	return r
}

// Body ..
// Set the request struct which will be converted to json during request
func (r *RequestMaker) Body(request interface{}) *RequestMaker {
	r.body = request
	return r
}

// Scan ..
// Maps the response to response struct
func (r *RequestMaker) Scan(response interface{}) *RequestMaker {
	r.response = response
	return r
}

// SetBasicAuth ..
// Set the base64 auth token in header
func (r *RequestMaker) SetBasicAuth(username, password string) *RequestMaker {
	r.username = username
	r.password = password
	return r
}

// SetContentType ..
// Set the content type of the request
func (r *RequestMaker) SetContentType(contentType string) *RequestMaker {
	r.contentType = contentType
	return r
}

// SetLogger ..
// Set the logger
func (r *RequestMaker) SetLogger(logger *logrus.Entry) *RequestMaker {
	r.logger = logger
	return r
}

// SetVerbose ..
// Displays the information about the request
func (r *RequestMaker) Verbose(verbose bool) *RequestMaker {
	r.verbose = verbose
	return r
}

// Dispatch ..
// This is where actuall request made to destination
func (r *RequestMaker) Dispatch() error {

	if strings.Contains(r.url, "/:") {

		// Set the path params
		splits := strings.Split(r.url, urlSeparator)
		var idx int
		for i, s := range splits {

			if strings.HasPrefix(s, pathParamIndicator) {
				splits[i] = r.pathParams[idx]
				idx++
			}
		}

		// Verify all the path params are set
		if idx != len(r.pathParams) {
			return errPathParamConflift
		}

		// sets the final url after injecting path params
		r.url = strings.Join(splits, urlSeparator)
	}

	if r.method == "" {
		return errMethodUnknown
	}

	var req *http.Request
	var err error
	// Set the body
	if r.body != nil {

		if r.method != POST {
			return fmt.Errorf(errCannotSetBody, r.method)
		}

		if URLENCODED == r.contentType {

			// create a request
			req, err = http.NewRequest(r.method, r.url, bytes.NewBufferString(r.body.(url.Values).Encode()))

		} else if JSON == r.contentType {

			bits, err := json.Marshal(r.body)
			if err != nil {
				return fmt.Errorf("Failed to marshall json request: ", err)
			}

			// create a request
			req, err = http.NewRequest(r.method, r.url, bytes.NewBuffer(bits))
		}

	} else {

		// create a request
		req, err = http.NewRequest(r.method, r.url, nil)
	}

	// checks for http request creation errors if any
	if err != nil {
		return fmt.Errorf("Can't create request ", err)
	}

	// Sets the basic auth header
	if r.username != "" || r.password != "" {
		req.Header.Add("Authorization", "Basic "+basicAuth(r.username, r.password))
	}

	// Sets the content type
	if r.contentType != "" {
		req.Header.Add("Content-Type", r.contentType)
	}

	// Sets the header
	if len(r.headers) > 0 {

		for k, v := range r.headers {
			req.Header.Add(k, v)
		}
	}

	// Set the query params
	if len(r.queryParams) > 0 {

		q := req.URL.Query()
		for k, v := range r.queryParams {
			q.Add(k, v)
		}

		params := q.Encode()
		if r.unescapeQueryParams {
			params, err = url.QueryUnescape(params)
			if err != nil {
				return fmt.Errorf("Query unescape failed:", err)
			}
		}

		req.URL.RawQuery = params
	}

	if r.verbose {
		r.logger.Infoln("Making request to = ", req.URL.String())
	}

	// dispatch the request
	res, err := http.DefaultClient.Do(req)

	// Scans the response
	if err != nil {
		if res != nil {
			res.Body.Close()
		}
		return fmt.Errorf("Dispatching request failed :", err)
	}
	defer res.Body.Close()

	if r.response != nil {

		// read the response body
		bits, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		// r.logger.Infoln("Response = ", string(bits[:]))
		// decode to response type
		err = json.NewDecoder(bytes.NewBuffer(bits)).Decode(r.response)
		if err != nil {
			return fmt.Errorf("Failed to decode respose : ", err)
		}
	}
	return nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
