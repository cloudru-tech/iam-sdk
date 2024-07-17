package sdk

import "github.com/cloudru-tech/iam-sdk/api"

// Option is a functional option for the SDK.
type Option func(k *Kit)

// Config is the configuration for the SDK.
type Config struct {
	AccessKey *AccessKey
}

// AccessKey is a pair of key_id and secret
// used to authenticate the client.
type AccessKey struct {
	KeyID  string
	Secret string
}

// WithAccessKey sets the access key for the SDK configuration.
func WithAccessKey(kid, secret string) Option {
	return func(k *Kit) {
		k.conf.AccessKey = &AccessKey{KeyID: kid, Secret: secret}
	}
}

// WithCustomClient is used to override default cloud.ru client.
// If this option won't be provided, the SDK will create a new one with
// default options and connects to the automatically discovered IAM server.
func WithCustomClient(client *api.Client) Option {
	return func(k *Kit) {
		k.client = client
	}
}
