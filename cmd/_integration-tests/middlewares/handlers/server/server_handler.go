package handler

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"

	pb "github.com/TuneLab/go-truss/cmd/_integration-tests/middlewares/middlewarestest-service"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.MiddlewaresTestServer {
	return middlewarestestService{}
}

type middlewarestestService struct{}

// AlwaysWrapped implements Service.
func (s middlewarestestService) AlwaysWrapped(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	everyone := ctx.Value("Everyone")
	if e, ok := everyone.(bool); !ok || !e {
		return nil, errors.New("ctx does not contain value for key 'Everyone'")
	}

	notSometimes := ctx.Value("notsometimes")
	if ns, ok := notSometimes.(bool); !ok || !ns {
		return nil, errors.New("ctx does not contain value for key 'notsometimes'")
	}
	return &pb.Empty{}, nil
}

// SometimesWrapped implements Service.
func (s middlewarestestService) SometimesWrapped(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	everyone := ctx.Value("Everyone")
	if e, ok := everyone.(bool); !ok || !e {
		return nil, errors.New("ctx does not contain value for key 'Everyone'")
	}

	notSometimes := ctx.Value("notsometimes")
	// SometimesWrapped should not have the "NotSometimes" value in the context
	if ns, ok := notSometimes.(bool); ok && ns {
		return nil, errors.New(
			"ctx contains value for key 'NotSometimes', implying that this endpoint has been unexpectedly wrapped")
	}
	return &pb.Empty{}, nil
}
