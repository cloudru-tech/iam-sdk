package sdk

import (
	"context"
	"fmt"
	"time"

	authV1 "github.com/cloudru-tech/iam-sdk/api/auth/v1"
	"github.com/cloudru-tech/iam-sdk/token"
)

// AccessToken is an IAM access token.
type AccessToken struct {
	// AccessToken is the access token, is used to authenticate the user.
	AccessToken string
	// ExpiresAt is the token expiration date.
	ExpiresAt time.Time
}

// GetToken returns the IAM access token token.
func (k *Kit) GetToken(ctx context.Context) (*token.Token, error) {
	res, err := k.client.AuthService.GetToken(ctx, &authV1.GetTokenRequest{
		KeyId:  k.conf.AccessKey.KeyID,
		Secret: k.conf.AccessKey.Secret,
	})
	if err != nil {
		return nil, err
	}

	t, err := token.ParseToken(res.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("internal error, retrieved token could not be parsed: %w", err)
	}
	return t, nil
}
