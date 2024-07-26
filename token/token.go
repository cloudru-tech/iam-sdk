package token

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// SubjectType is the type of the subject.
type SubjectType string

const (
	// SubjectTypeUser is the user subject type.
	SubjectTypeUser SubjectType = "user"
	// SubjectTypeServiceAccount is the service account subject type.
	SubjectTypeServiceAccount SubjectType = "service_account"
)

// Aud is the list of applications to which the subject has access.
type Aud []string

// Token is an IAM access token.
type Token struct {
	// Aud - `aud` parameter of the jwt token
	Aud Aud `json:"aud"`
	// Issuer is the token issuer.
	Issuer string `json:"iss"`
	// Username is the subject username property
	// For the SubType=user it will contain the user username.
	// For the SubType=service_account it will contain the service account name.
	Username string `json:"preferred_username"`
	// Email is the subject email address.
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	// SubID is the subject identifier.
	SubID uuid.UUID `json:"sub"`
	// SubType is the subject type
	SubType SubjectType `json:"sub_type"`

	// AccessToken is the access token, is used to authenticate the user.
	AccessToken string `json:"-"`
	// ExpiresAt is the token expiration date.
	// The time is in UTC. Use the time.Location() method to get the local time.
	ExpiresAt time.Time `json:"-"`
	// Claims original claims from the token.
	Claims jwt.MapClaims `json:"-"`
}

// ParseToken parses the JWT token and returns the Token object.
//
// WARNING: Don't use this method unless you know what you're doing.
//
// It's only ever useful in cases where you know the signature is valid (since it has already
// been or will be checked elsewhere in the stack) and you want to extract values from it.
func ParseToken(accessToken string) (*Token, error) {
	t := &Token{
		AccessToken: accessToken,
		Claims:      jwt.MapClaims{},
	}
	parser := new(jwt.Parser)
	accessToken = strings.TrimSpace(accessToken)

	_, _, err := parser.ParseUnverified(accessToken, &t.Claims)
	if err != nil {
		return nil, fmt.Errorf("parse token: %w", err)
	}

	data, _ := json.Marshal(t.Claims)
	_ = json.Unmarshal(data, t)

	exp, ok := t.Claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("token expiration date is missing")
	}

	t.ExpiresAt = time.Unix(int64(exp), 0)
	return t, nil
}

func (a *Aud) UnmarshalJSON(data []byte) error {
	if data[0] == 34 { // QuoteByte
		var v string
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		*a = []string{v}
	} else {
		var v []string
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		*a = v
	}

	return nil
}

func (a *Aud) Has(value string) bool {
	for _, v := range *a {
		if v == value {
			return true
		}
	}

	return false
}
