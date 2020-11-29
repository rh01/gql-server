package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go/request"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/store"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"report/internal/graphql/models"
	"strings"
)

type key string

// CurrentUserKey for middleware
const CurrentUserKey key = "currentUser"

// AuthMiddleware to authenticate graphql users
func AuthMiddleware(repo store.Report) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := parseToken(r)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				next.ServeHTTP(w, r)
				return
			}

			var user = models.User{}
			q := bson.M{}
			q["username"] = claims["jti"].(string)
			if err := repo.Query.GetStore("user").FindOne(q, &user); err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), CurrentUserKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromToken,
}

var graphqlHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Sec-Websocket-Protocol"},
	Filter:    stripWebSocketHeader,
}

func stripBearerPrefixFromToken(token string) (string, error) {
	bearer := "BEARER"

	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}

	return token, nil
}

func stripWebSocketHeader(token string) (string, error) {
	head := "GRAPHQL-WS,"

	if len(token) > len(head) && strings.ToUpper(token[0:len(head)]) == head {
		return token[len(head)+1:], nil
	}

	return token, nil
}

var authExtractor = &request.MultiExtractor{
	authHeaderExtractor,
	graphqlHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func parseToken(r *http.Request) (*jwt.Token, error) {
	jwtToken, err := request.ParseFromRequest(r, authExtractor, func(token *jwt.Token) (interface{}, error) {
		t := []byte(os.Getenv("JWT_SECRET"))
		return t, nil
	})

	return jwtToken, errors.Wrap(err, "parseToken error: ")
}

// GetCurrentUserFromCTX from context
func GetCurrentUserFromCTX(ctx context.Context) (*models.User, error) {
	errNoUserInContext := errors.New("no user in context")
	if ctx.Value(CurrentUserKey) == nil {
		return nil, errNoUserInContext
	}

	user, ok := ctx.Value(CurrentUserKey).(*models.User)
	if !ok || user.ID == "" {
		return nil, errNoUserInContext
	}

	return user, nil
}
