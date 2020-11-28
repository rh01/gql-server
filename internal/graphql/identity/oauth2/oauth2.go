package oauth2

import (
	"context"
	"encoding/base64"
	"net/http"
	"report/internal/graphql"
	"strconv"

	"golang.org/x/oauth2"

	"net/url"

	"fmt"

	"bytes"
	"time"

	"strings"

	oidc "github.com/coreos/go-oidc"
)

// Constants for performing encode decode
const (
	EQUAL        = "="
	DOUBLEEQUALS = "=="
	DOT0         = ".0"
	DOT1         = ".1"
	DOT2         = ".2"
	SLASH        = "/"
	SEMICOLON    = ";"
)

// App ...
type App struct {
	clientID     string
	clientSecret string
	redirectURI  string

	verifier *oidc.IDTokenVerifier
	provider *oidc.Provider

	// Does the provider use "offline_access" scope to request a refresh token
	// or does it use "access_type=offline" (e.g. Google)?
	offlineAsScope bool

	ctx context.Context
}

// Auth ..
type authServer struct {
	app          App
	issuerURL    url.URL
	listen       string
	tlsCert      string
	tlsKey       string
	rootCAs      string
	debug        bool
	oauth2Config oauth2.Config
}

var (
	idTokenName = "__idt"
)

// NewAuthServer ..
func NewAuthServer(ctx context.Context, r *http.ServeMux, c graphql.Config) error {

	authServ := authServer{}
	a := App{}
	a.ctx = ctx
	a.clientID = c.Identity.ID
	a.clientSecret = c.Identity.Secret
	a.redirectURI = c.Identity.RedirectURI

	redirectURI, err := url.Parse(c.Identity.RedirectURI)
	if err != nil {
		return fmt.Errorf("Erorr when parsing dex redirect url %s : %v", c.Identity.RedirectURI, err)
	}

	issuer, err := url.Parse(c.Identity.Issuer)
	if err != nil {
		return fmt.Errorf("Erorr when parsing dex url %s : %v", c.Identity.Issuer, err)
	}
	authServ.issuerURL = *issuer

	provider, err := oidc.NewProvider(ctx, c.Identity.Issuer)
	if err != nil {
		return fmt.Errorf("Failed to query provider %q: %v", c.Identity.Issuer, err)
	}

	var s struct {
		// What scopes does a provider support?
		//
		// See: https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata
		ScopesSupported []string `json:"scopes_supported"`
	}
	if err := provider.Claims(&s); err != nil {
		return fmt.Errorf("Failed to parse provider scopes_supported: %v", err)
	}

	if len(s.ScopesSupported) == 0 {
		// scopes_supported is a "RECOMMENDED" discovery claim, not a required
		// one. If missing, assume that the provider follows the spec and has
		// an "offline_access" scope.
		a.offlineAsScope = true
	} else {
		// See if scopes_supported has the "offline_access" scope.
		a.offlineAsScope = func() bool {
			for _, scope := range s.ScopesSupported {
				if scope == oidc.ScopeOfflineAccess {
					return true
				}
			}
			return false
		}()
	}

	a.provider = provider

	// Configure an OpenID Connect aware OAuth2 client.
	authServ.oauth2Config = oauth2.Config{
		ClientID:     a.clientID,
		ClientSecret: a.clientSecret,
		RedirectURL:  a.redirectURI,

		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),

		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}

	a.verifier = provider.Verifier(&oidc.Config{
		ClientID: a.clientID,
		// SkipNonceCheck: true,
	})

	authServ.app = a

	r.HandleFunc("/login", authServ.login)
	r.HandleFunc(redirectURI.Path, authServ.handleOAuth2Callback)

	return nil
}

func (a *authServer) login(w http.ResponseWriter, r *http.Request) {

	// opts := oauth2.SetAuthURLParam("id", encode(r.Host))
	// fmt.Println(encode(r.Host))
	var buf bytes.Buffer
	buf.WriteString(r.URL.Scheme)
	buf.WriteString(SEMICOLON)
	buf.WriteString(r.Host)
	buf.WriteString(SEMICOLON)
	buf.WriteString(strconv.Itoa(time.Now().Nanosecond()))
	state := encode(buf.String())

	authURL := a.oauth2Config.AuthCodeURL(state)
	http.Redirect(w, r, authURL, http.StatusFound)
}

func (a *authServer) handleOAuth2Callback(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	oauth2Token, err := a.oauth2Config.Exchange(a.app.ctx, params.Get("code"))
	if err != nil {
		// handle error
	}

	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		// handle missing token
	}

	// Parse and verify ID Token payload.
	idToken, err := a.app.verifier.Verify(a.app.ctx, rawIDToken)
	if err != nil {
		// handle error
	}

	// Extract custom claims
	var claims struct {
		Email    string `json:"email"`
		Verified bool   `json:"email_verified"`
	}

	if err := idToken.Claims(&claims); err != nil {
		// handle error
	}

	state := decode(params.Get("state"))
	states := strings.Split(state, SEMICOLON)
	fmt.Println(states[0])
	cookie := &http.Cookie{
		Name:     idTokenName,
		Value:    rawIDToken,
		HttpOnly: true,
		Path:     "/",
		//Secure : true, // TODO enable this to ensure the cookie is passed only with https
	}
	http.SetCookie(w, cookie)

	homeURL := states[0] + "//" + states[1] + "/vcs"
	http.Redirect(w, r, homeURL, http.StatusMovedPermanently)
}

func encode(id string) string {

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

func decode(id string) string {

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
