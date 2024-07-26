package sdk

import (
	"fmt"

	"github.com/cloudru-tech/iam-sdk/api"
)

// Kit is a collection of tools for working
// with the Cloud.ru IAM service.
type Kit struct {
	client *api.Client
	conf   *Config
}

// New creates a new IAM SDK instance.
func New(opts ...Option) (*Kit, error) {
	k := &Kit{conf: &Config{}}
	for _, opt := range opts {
		opt(k)
	}

	switch {
	case k.conf.AccessKey == nil, k.conf.AccessKey.KeyID == "", k.conf.AccessKey.Secret == "":
		return nil, fmt.Errorf("AccessKey must be provided")

	}

	if k.client == nil {
		client, err := api.New(&api.Config{})
		if err != nil {
			return nil, fmt.Errorf("instantiate cloud.ru IAM gRPC client: %w", err)
		}

		k.client = client
	}

	return k, nil
}

// Close closes the IAM SDK client connection.
func (k *Kit) Close() error { return k.client.Close() }
