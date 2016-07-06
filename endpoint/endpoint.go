package endpoint

import (
	"github.com/crezam/go-kit-app/businesslogic"
	"github.com/crezam/go-kit-app/rpc"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func makeUpperCaseEndpoint(service businesslogic.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(rpc.UpperCaseRequest)
		v, err := service.UpperCase(req.OriginalWord)
		if err != nil {
			return rpc.UpperCaseResponse{v, err.Error()}, nil
		}
		return rpc.UpperCaseResponse{v, ""}, nil
	}
}
