package vcs

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"report/internal/graphql/identity/oauth2/providers"
	"report/internal/graphql/store"
	"report/internal/pkg/logger"
	"strings"

	"bytes"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (

	// VCS
	errNoProviderFound         = "No provider found for %s"
	errGetUpdatedFokenFailed   = "Failed to get updated token %s"
	errGettingRepositories     = "Failed to get repositories for %s"
	errVCSAccountAlreadyLinked = "VCS account already linked"
)

// True or False
const (
	True  = 1
	False = 0
)

// Constants for performing encode decode
const (
	EQUAL        = "="
	DOUBLEEQUALS = "=="
	DOT0         = ".0"
	DOT1         = ".1"
	DOT2         = ".2"
)

// Common constants
const (
	SLASH     = "/"
	SEMICOLON = ";"
)

type service struct {
	store     store.Vcs
	logger    *logrus.Entry
	providers providers.Providers
}

// Service ..
type Service interface {
	Authorize(w http.ResponseWriter, r *http.Request)
	Authorized(w http.ResponseWriter, r *http.Request)
}

// NewService ..
func NewService(loggr logger.Loggr, d store.Database, providers providers.Providers, s store.Report) Service {

	l := loggr.GetLogger("service/vcs")
	return &service{
		store:     s.Vcs,
		logger:    l,
		providers: providers,
	}
}

func (s service) Authorize(w http.ResponseWriter, r *http.Request) {
	provider := mux.Vars(r)["provider"]
	// get provider ...
	p, err := s.providers.Get(provider)
	if err != nil {
		http.Error(w, fmt.Sprintf("Getting provider %s failed: %v", provider, err), http.StatusBadRequest)
		return
	}

	// source := vcs.GetSource(provider)
	var buf bytes.Buffer
	// buf.WriteString(source)
	buf.WriteString(SEMICOLON)
	buf.WriteString(SEMICOLON)
	buf.WriteString(SLASH)
	buf.WriteString(SLASH)
	buf.WriteString(r.Host)
	url := p.Authorize(s.encode(buf.String()))
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Authorized ..
// Invoked when authorization finished by oauth app
func (s service) Authorized(w http.ResponseWriter, r *http.Request) {

	provider := mux.Vars(r)["provider"]
	p, err := s.providers.Get(provider)
	if err != nil {
		http.Error(w, fmt.Sprintf("Getting provider %s failed: %v", provider, err), http.StatusBadRequest)
	}

	id := r.FormValue("id")
	code := r.FormValue("code")
	u, err := p.Authorized(id, code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Finalize the authorization failed : %v", err), http.StatusBadRequest)
	}

	unescID := s.decode(id)
	escID := strings.Split(unescID, SEMICOLON)

	// persist user
	source := escID[0]
	u.Source = source

	// url := escID[2] + "/integrations"
	http.Redirect(w, r, "/integrations", http.StatusTemporaryRedirect)
}

func (s service) encode(id string) string {

	eid := base64.URLEncoding.EncodeToString([]byte(id))
	if strings.Contains(eid, DOUBLEEQUALS) {
		eid = strings.TrimRight(eid, DOUBLEEQUALS) + DOT2
	} else if strings.Contains(eid, EQUAL) {
		eid = strings.TrimRight(eid, EQUAL) + DOT1
	} else {
		eid = eid + DOT0
	}
	return eid
}

func (s service) decode(id string) string {

	if strings.Contains(id, DOT2) {
		id = strings.TrimRight(id, DOT2) + DOUBLEEQUALS
	} else if strings.Contains(id, DOT1) {
		id = strings.TrimRight(id, DOT1) + EQUAL
	} else {
		id = strings.TrimRight(id, DOT0)
	}
	did, _ := base64.URLEncoding.DecodeString(id)
	return string(did[:])
}
