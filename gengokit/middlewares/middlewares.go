package middlewares

import (
	"io"

	"github.com/pkg/errors"

	"github.com/TuneLab/go-truss/gengokit"
	//"github.com/TuneLab/go-truss/svcdef"
	//"github.com/TuneLab/go-truss/gengokit/middlewares/templates"
)

const middlewaresPath = "NAME-server/middlewares/middlewares.gotemplate"

type middlewares struct {
}

func (m middlewares) Render(f string, te *gengokit.TemplateExecutor) (io.Reader, error) {
	if f != middlewaresPath {
		return nil, errors.Errorf("cannot render unknown file: %q", f)
	}
	return nil, nil

}
