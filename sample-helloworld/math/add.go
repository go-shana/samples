package math

import (
	"context"

	"github.com/go-shana/core/rpc"
)

func init() {
	rpc.Export(Add)
}

type AddRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type AddResponse struct {
	Sum int `json:"sum"`
}

func Add(ctx context.Context, req *AddRequest) (resp *AddResponse, err error) {
	resp = &AddResponse{
		Sum: req.A + req.B,
	}
	return
}
