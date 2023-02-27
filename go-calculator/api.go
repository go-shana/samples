package calculator

import (
	"context"
	"math/big"

	"github.com/go-shana/core/errors"
	"github.com/go-shana/core/rpc"
	"github.com/go-shana/samples/go-calculator/internal/model"
)

func init() {
	// Export this method so it can be called by RPC clients.
	rpc.Export(Add)
	rpc.Export(Sub)
	rpc.Export(Mul)
	rpc.Export(Div)
	rpc.Export(Value)
}

var (
	errOK         = errors.NewErrorCode(ErrorCodeOK, "ok")
	errDivideZero = errors.NewErrorCode(ErrorCodeDivideZero, "divide by zero")
)

type Request struct {
	value *big.Int
	Value string `json:"value"`
}

// Validate makes sure the value is a valid integer.
func (req *Request) Validate(ctx context.Context) {
	value := new(big.Int)

	if _, ok := value.SetString(req.Value, 10); !ok {
		errors.Throwf("invalid value: %s", req.Value)
	}

	req.value = value
}

type Response struct {
	Value string `json:"value"`
}

func Add(ctx context.Context, req *Request) (resp *Response, err error) {
	result := model.Add(req.value)
	resp = &Response{
		Value: result.String(),
	}
	err = errOK
	return
}

func Sub(ctx context.Context, req *Request) (resp *Response, err error) {
	result := model.Sub(req.value)
	resp = &Response{
		Value: result.String(),
	}
	err = errOK
	return
}

func Mul(ctx context.Context, req *Request) (resp *Response, err error) {
	result := model.Mul(req.value)
	resp = &Response{
		Value: result.String(),
	}
	err = errOK
	return
}

func Div(ctx context.Context, req *Request) (resp *Response, err error) {
	// Check req.value is not zero.
	if req.value.Sign() == 0 {
		err = errDivideZero
		return
	}

	result := model.Div(req.value)
	resp = &Response{
		Value: result.String(),
	}
	err = errOK
	return
}

func Value(ctx context.Context, req *struct{}) (resp *Response, err error) {
	resp = &Response{
		Value: model.String(),
	}
	err = errOK
	return
}
