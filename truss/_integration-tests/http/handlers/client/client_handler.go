package clienthandler

import (
	pb "github.com/TuneLab/go-truss/truss/_integration-tests/http/httptest-service"
)

// GetWithQuery implements Service.
func GetWithQuery(AGetWithQuery int64, BGetWithQuery int64) (*pb.GetWithQueryRequest, error) {
	request := pb.GetWithQueryRequest{
		A: AGetWithQuery,
		B: BGetWithQuery}
	return &request, nil
}

// GetWithRepeatedQuery implements Service.
func GetWithRepeatedQuery(AGetWithRepeatedQuery string) (*pb.GetWithRepeatedQueryRequest, error) { // Add custom business logic for interpreting AGetWithRepeatedQuery,
	request := pb.GetWithRepeatedQueryRequest{}
	return &request, nil
}
