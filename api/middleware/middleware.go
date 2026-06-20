package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/httplog/v2"

	oapimiddleware "github.com/oapi-codegen/nethttp-middleware"
)

func OpenAPIValidation(urlPrefix string, openapiSpec *openapi3.T) func(http.Handler) http.Handler {
	// Set server url to the prefix. This ensures request validation doesn't fail.
	// See: https://github.com/oapi-codegen/oapi-codegen/issues/1123
	openapiSpec.Servers = openapi3.Servers{&openapi3.Server{URL: urlPrefix}}

	return oapimiddleware.OapiRequestValidatorWithOptions(
		openapiSpec,
		&oapimiddleware.Options{
			SilenceServersWarning: true,
			Options: openapi3filter.Options{
				AuthenticationFunc: func(_ context.Context, _ *openapi3filter.AuthenticationInput) error {
					return nil // Do nothing
				},
			},
		},
	)
}

func Logger(name string) func(http.Handler) http.Handler {
	// This middleware includes the recover middleware
	return httplog.RequestLogger(
		httplog.NewLogger(
			name,
			httplog.Options{
				LogLevel:       slog.LevelDebug,
				RequestHeaders: true,
				Concise:        true,
			},
		),
	)
}
