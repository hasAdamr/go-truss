package templates

// file_section as var name
// middlewares.go base
const middlewares_base = `
package middlewares

import (
	"github.com/go-kit/kit/endpoint"
	pb "{{.PBImportPath -}}"
)

func InjectEndpointMiddlewares(in endpoint.Endpoint) endpoint.Endpoint {
	return in
}

func InjectServiceMiddlewares(in pb.{{.Service.Name}}Server) pb.{{.Service.Name}}Server {
	return in
}
`
