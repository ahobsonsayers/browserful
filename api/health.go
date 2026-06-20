package api

import "context"

func (Server) HealthCheck(context.Context, HealthCheckRequestObject) (HealthCheckResponseObject, error) {
	return HealthCheck200Response{}, nil
}
