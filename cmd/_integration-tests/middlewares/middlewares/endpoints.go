package middlewares

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	//pb "github.com/TuneLab/go-truss/cmd/_integration-tests/middlewares/middlewarestest-service"
	svc "github.com/TuneLab/go-truss/cmd/_integration-tests/middlewares/middlewarestest-service/generated"
)

// WrapEndpoints accepts the service's entire collection of endpoints, so that a
// set of middlewares can be wrapped around every middleware (e.g., access
// logging and instrumentation), and others wrapped selectively around some
// endpoints and not others (e.g., endpoints requiring authenticated access).
func WrapEndpoints(in svc.Endpoints) svc.Endpoints {

	// Pass in the middlewares you want applied to every endpoint.
	in.WrapAll(addBoolToContext("Everyone"))
	in.WrapAll(addBoolToContext("notsometimes"), "SometimesWrapped")
	//in.WrapAll(setWrapAllTest(&pb.WrapAllTest{Always: true}))
	//in.WrapAll(setWrapAllTest(
	//&pb.WrapAllTest{
	//Always:       true,
	//NotSometimes: true,
	//}),
	//in.SometimesWrappedEndpoint,
	//)

	// How to apply a middleware selectively.
	// in.ExampleEndpoint = authMiddleware(in.ExampleEndpoint)

	return in
}

//func setWrapAllTest(resp *pb.WrapAllTest) endpoint.Middleware {
//return func(next endpoint.Endpoint) endpoint.Endpoint {
//return func(ctx context.Context, resp interface{}) (interface{}, error) {
//return next(ctx, resp)
//}
//}
//}

func addBoolToContext(key string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			ctx = context.WithValue(ctx, key, true)
			return next(ctx, request)
		}
	}
}
