package helloworld

import (
	"context"

	"github.com/go-shana/core/config"
	"github.com/go-shana/core/errors"
	"github.com/go-shana/core/rpc"
	"github.com/go-shana/core/validator/numeric"
)

func init() {
	// Make this method available to the RPC server.
	rpc.Export(SayHello)
}

var (
	errOK            = errors.NewErrorCode(0, "ok")
	errInvalidParams = errors.NewErrorCode(1000, "invalid params")

	serviceConfig = config.New[HelloConfig]("service")
)

type HelloConfig struct {
	Welcome string
}

func (c *HelloConfig) Init(ctx context.Context) {
	if c.Welcome == "" {
		c.Welcome = "Hello"
	}
}

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func (req *HelloRequest) Validate(ctx context.Context) {
	// Rethrow any error as errInvalidParams.
	defer errors.Rethrow(errInvalidParams)

	numeric.InRange(len(req.Name), 1, 10)
}

func SayHello(ctx context.Context, req *HelloRequest) (resp *HelloResponse, err error) {
	resp = &HelloResponse{
		Message: serviceConfig.Welcome + ", " + req.Name,
	}

	// Optionally set err to errOK if client expects a zero code in response.
	err = errOK
	return
}
