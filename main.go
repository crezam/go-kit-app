package main

import (
	"golang.org/x/net/context"
	"github.com/crezam/go-kit-app/businesslogic"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/crezam/go-kit-app/endpoint"
	"github.com/crezam/go-kit-app/encoding"
	"net/http"
	"github.com/labstack/gommon/log"
)

func main() {
	ctx := context.Background()
	svc := businesslogic.StringServiceImpl{}

	uppercaseHandler := httptransport.NewServer(
		ctx,
		endpoint.MakeUpperCaseEndpoint(svc),
		encoding.DecodeUpperCaseRequest,
		encoding.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)

	log.Fatal(http.ListenAndServe(":8080",nil))
}
