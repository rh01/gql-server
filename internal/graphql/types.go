package graphql

import (
	"time"
)

// Token ..
type Token struct {
	Username string    `json:"user_name"`
	Subject  string    `json:"sub"`
	Issuer   string    `json:"iss"`
	Audience string    `json:"aud"`
	Scope    []string  `json:"scope"`
	AuthTime time.Time `json:"auth_time"`
	Expiry   time.Time `json:"exp"`
	IssuedAt time.Time `json:"iat"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone_number"`
}

// AuthRequest ..
type AuthRequest struct {
	ID            string
	ClientID      string
	Scopes        []string
	RedirectURI   string
	ResponseTypes []string
	State         string
	Nonce         string
}
