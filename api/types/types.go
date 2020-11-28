package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// VCS contains the information common amongst most OAuth and OAuth2 providers.
// All of the "raw" datafrom the provider can be found in the `RawData` field.
type VCS struct {
	ID           string    `json:"id" bson:"id,omitempty"`
	Name         string    `json:"name" bson:"name,omitempty"`
	Kind         string    `json:"kind" bson:"kind,omitempty"`
	Link         string    `json:"link" bson:"link,omitempty"`
	Source       string    `json:"source" bson:"source"`
	OwnerType    string    `json:"owner_type" bson:"owner_type,omitempty"`
	AvatarURL    string    `json:"avatar" bson:"avatar,omitempty"`
	AccessCode   string    `json:"access_code" bson:"access_code,omitempty"`
	AccessToken  string    `json:"access_token" bson:"access_token,omitempty"`
	RefreshToken string    `json:"refresh_token" bson:"refresh_token,omitempty"`
	TokenExpiry  time.Time `json:"token_expiry" bson:"token_expiry,omitempty"`
	SecretID     string    `json:"-" bson:"secret_id"`
}

// VCSSysConf ..(sysconf)
type VCSSysConf struct {
	// common fields for any sys config
	ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name string        `bson:"name,omitempty" json:"name"`
	Kind string        `bson:"kind,omitempty" json:"kind"`

	Key         string `bson:"key,omitempty" json:"key"`
	Secret      string `bson:"secret,omitempty" json:"secret"`
	CallbackURL string `bson:"callback_url,omitempty" json:"callback_url"`
	HookURL     string `bson:"hook_url,omitempty" json:"hook_url"`
}
