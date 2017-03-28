package test

import (
	"testing"

	"golang.org/x/net/context"

	pb "github.com/TuneLab/go-truss/cmd/_integration-tests/middlewares/middlewarestest-service"
)

func TestAlwaysWrapped(t *testing.T) {
	ctx := context.Background()

	_, err := middlewareEndpoints.AlwaysWrappedEndpoint(ctx, &pb.Empty{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSometimesWrapped(t *testing.T) {
	ctx := context.Background()

	_, err := middlewareEndpoints.SometimesWrappedEndpoint(ctx, &pb.Empty{})
	if err != nil {
		t.Fatal(err)
	}
}
