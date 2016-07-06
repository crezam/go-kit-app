package encoding

import (
	"golang.org/x/net/context"
	"net/http"
	"github.com/crezam/go-kit-app/rpc"
	"encoding/json"
)

func DecodeUpperCaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request rpc.UpperCaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}


