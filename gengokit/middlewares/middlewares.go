package middlewares

import (
	"io"

	"github.com/pkg/errors"
	. "github.com/y0ssar1an/q"

	"github.com/TuneLab/go-truss/gengokit"
	//"github.com/TuneLab/go-truss/svcdef"
	"github.com/TuneLab/go-truss/gengokit/middlewares/templates"
)

const EndpointsFile = "NAME-service/middlewares/endpoints.gotemplate"
const ServiceFile = "NAME-service/middlewares/service.gotemplate"

func New() *Middlewares {
	var m Middlewares

	return &m
}

type Middlewares struct {
	prevEndpoints io.Reader
	prevService   io.Reader
}

func (m *Middlewares) LoadEndpoints(prev io.Reader) {
	Q(prev)
	m.prevEndpoints = prev
}

func (m *Middlewares) LoadService(prev io.Reader) {
	Q(prev)
	m.prevService = prev
}

func (m *Middlewares) Render(alias string, executor *gengokit.Executor) (io.Reader, error) {
	switch alias {
	case EndpointsFile:
		if m.prevEndpoints != nil {
			Q("Keeping old endpoint Middlewares")
			return m.prevEndpoints, nil
		}
		Q("creating new endpoints middlewares")
		return executor.ApplyTemplate(templates.EndpointsBase, "Endpoint")
	case ServiceFile:
		if m.prevService != nil {
			Q("Keeping old service Middlewares")
			return m.prevService, nil
		}
		Q("creating new service middlewares")
		return executor.ApplyTemplate(templates.ServiceBase, "Service")
	}
	return nil, errors.Errorf("cannot render unknown file: %q", alias)
}
