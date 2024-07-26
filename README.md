# Cloud.ru IAM SDK

  
---  

iam-sdk is a Go client library for accessing the Cloud.ru IAM API.

## Installation

iam-sdk is compatible with modern Go releases in module mode, with Go installed:

```bash  
go get github.com/cloudru-tech/iam-sdk@latest
```  

will resolve the latest version of SDK and add the package to the current module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go  
import "github.com/cloudru-tech/iam-sdk"  
```  

and then run `go get` without arguments.

## Usage

Here's a simple example to get you started:

```go  
package main

import (
	"context"
	"fmt"
	"github.com/cloudru-tech/iam-sdk"
)

func main() {
	kit, err := sdk.New(
		sdk.WithAccessKey("key_id", "key_secret"),
	)

	// ...
}  
```  

The option `WithAccessKey` **is required** and used to provide the credentials to the SDK client.

AccessKey can be obtained from
the [Cloud.ru IAM console](https://cloud.ru/docs/console_api/ug/topics/guides__service_accounts_key.html#id1).

### Retrieve an access token

```go  
package main

import (
	"context"
	"log"
	
	"github.com/cloudru-tech/iam-sdk"
)

func main() {
	kit, err := sdk.New(
		sdk.WithAccessKey("key_id", "key_secret"),
	)
	if err != nil {
		log.Fatalf("failed to create iam kit: %v", err)
	}

	token, err := kit.GetToken(context.Background())
	if err != nil {
		log.Fatalf("failed to get the token: %v", err)
	}

	// ...
}

```

### Using the custom gRPC client

Alternatively you can use the custom gRPC client to interact with the IAM API, for example if you want to provide
additional gRPC options:

```go
package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	
	"github.com/cloudru-tech/iam-sdk"
	"github.com/cloudru-tech/iam-sdk/api"
)

func main() {
	client, err := api.New(
		&api.Config{},
		// any of your gRPC dial option
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                time.Second * 30,
			Timeout:             time.Second * 5,
			PermitWithoutStream: false,
		}),
		grpc.WithUserAgent("your-application-name"),
	)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	kit, err := sdk.New(
		sdk.WithAccessKey("key_id", "key_secret"),
		sdk.WithCustomClient(client),
	)

	// ...
}
```