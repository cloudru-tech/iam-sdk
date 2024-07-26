package api

import (
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	authV1 "github.com/cloudru-tech/iam-sdk/api/auth/v1"
	customersV1 "github.com/cloudru-tech/iam-sdk/api/customers/v1"
	permissionsV1 "github.com/cloudru-tech/iam-sdk/api/permissions/v1"
	projectsV1 "github.com/cloudru-tech/iam-sdk/api/projects/v1"
	saV1 "github.com/cloudru-tech/iam-sdk/api/service_accounts/v1"
)

// Config is the gRPC client connection configuration.
type Config struct {
	// Address is the IAM gRPC server address.
	// Property is optional. If this field is empty,
	// the client will try to get the IAM address from
	// the discovery service automatically.
	//  Example:
	//   "iam.api.cloud.ru:443"
	Address string
	// Insecure disables SSL verification.
	// This option will be ignored if the Address field is empty.
	Insecure bool
}

// Client is the IAM gRPC client with all available API services.
type Client struct {
	conn *grpc.ClientConn

	AuthService            authV1.AuthServiceClient
	CustomersService       customersV1.CustomersServiceClient
	PermissionsService     permissionsV1.PermissionsServiceClient
	ProjectsService        projectsV1.ProjectsServiceClient
	ServiceAccountsService saV1.ServiceAccountServiceClient
}

// New возвращает новый клиент к gRPC серверу
func New(conf *Config, opts ...grpc.DialOption) (*Client, error) {
	if conf == nil {
		conf = &Config{}
	}

	target := conf.Address
	if target != "" {
		if conf.Insecure {
			opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		} else {
			opts = append(opts,
				grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{MinVersion: tls.VersionTLS13})),
			)
		}
	} else {
		var err error
		target, err = DiscoverEndpoint()
		if err != nil {
			return nil, fmt.Errorf("discover cloud.ru IAM endpoint: %w", err)
		}
	}

	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,

		AuthService:            authV1.NewAuthServiceClient(conn),
		CustomersService:       customersV1.NewCustomersServiceClient(conn),
		PermissionsService:     permissionsV1.NewPermissionsServiceClient(conn),
		ProjectsService:        projectsV1.NewProjectsServiceClient(conn),
		ServiceAccountsService: saV1.NewServiceAccountServiceClient(conn),
	}, nil
}

// Close закрывает соединение с сервером
func (c *Client) Close() error {
	err := c.conn.Close()
	return err
}
