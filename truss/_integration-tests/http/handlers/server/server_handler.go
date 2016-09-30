package handler

// This file contains the Service definition, and a basic service
// implementation. It also includes service middlewares.

import (
	_ "errors"
	_ "time"

	"golang.org/x/net/context"

	_ "github.com/go-kit/kit/log"
	_ "github.com/go-kit/kit/metrics"

	pb "github.com/TuneLab/go-truss/truss/_integration-tests/http/httptest-service"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() Service {
	return httptestService{}
}

type httptestService struct{}

// GetWithQuery implements Service.
func (s httptestService) GetWithQuery(ctx context.Context, in *pb.GetWithQueryRequest) (*pb.GetWithQueryResponse, error) {
	_ = ctx
	_ = in
	response := pb.GetWithQueryResponse{
		V: in.A + in.B,
	}
	return &response, nil
}

// GetWithRepeatedQuery implements Service.

// V:

type Service interface {
	GetWithQuery(ctx context.Context, in *pb.GetWithQueryRequest) (*pb.GetWithQueryResponse, error)
}
