// Package generator generates a gokit service based on a service definition.
package generator

import (
	"bytes"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"

	"github.com/TuneLab/go-truss/gengokit"
	"github.com/TuneLab/go-truss/gengokit/handler"
	templFiles "github.com/TuneLab/go-truss/gengokit/template"

	"github.com/TuneLab/go-truss/svcdef"
	"github.com/TuneLab/go-truss/truss"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetOutput(os.Stderr)
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})
}

type templateRegistry map[string]gengokit.Renderable

func (tr templateRegistry) Register(alias string, r gengokit.Renderable) error {
	if tr[alias] != nil {
		return errors.Errorf("cannot register alias twice %q:", alias)
	}
	tr[alias] = r
	return nil
}

func (tr templateRegistry) isRegistered(alias string) bool {
	return tr[alias] != nil
}

func (tr templateRegistry) RenderAll(te *gengokit.TemplateExecutor) ([]truss.NamedReadWriter, error) {
	var renderedFiles []truss.NamedReadWriter

	for k, _ := range tr {
		f, err := tr.Render(k, te)
		if err != nil {
			return nil, errors.Wrap(err, "cannot render template")
		}
		if f == nil {
			continue
		}

		renderedFiles = append(renderedFiles, f)
	}

	return renderedFiles, nil
}

func (tr templateRegistry) Render(f string, te *gengokit.TemplateExecutor) (truss.NamedReadWriter, error) {
	fileContent, err := tr[f].Render(f, te)

	fileBytes, err := ioutil.ReadAll(fileContent)
	if err != nil {
		return nil, err
	}

	// ignore error as we want to write the code either way to inspect after
	// writing to disk
	formattedCode := formatCode(fileBytes)

	var r truss.SimpleFile

	// Set the path to the file and write the code to the file
	r.Path = templatePathToActual(f, te.PackageName)
	if _, err = r.Write(formattedCode); err != nil {
		return nil, err
	}

	return &r, nil
}

func createRegistry(te *gengokit.TemplateExecutor, p prevGen) (templateRegistry, error) {
	tr := make(templateRegistry)

	// handler
	const hPath = "NAME-service/handlers/server/server_handler.gotemplate"
	actualFP := templatePathToActual(hPath, te.PackageName)
	file := p[actualFP]
	h, err := handler.New(te.Service, file, te.PackageName)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot parse previous handler: %q", actualFP)
	}
	tr.Register(hPath, h)

	// for anything not registered yet
	for _, templFP := range templFiles.AssetNames() {
		if !tr.isRegistered(templFP) {
			tr.Register(templFP, defaultRender{})
		}
	}

	return tr, nil
}

type prevGen map[string]io.Reader

// GenerateGokit returns a gokit service generated from a service definition (svcdef),
// the package to the root of the generated service goPackage, the package
// to the .pb.go service struct files (goPBPackage) and any prevously generated files.
func GenerateGokit(sd *svcdef.Svcdef, conf gengokit.Config) ([]truss.NamedReadWriter, error) {
	te, err := gengokit.NewTemplateExecutor(sd, conf)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create template executor")
	}

	prev := make(prevGen, len(conf.PreviousFiles))
	for _, f := range conf.PreviousFiles {
		prev[f.Name()] = f
	}

	tr, err := createRegistry(te, prev)

	return tr.RenderAll(te)
}

// templatePathToActual accepts a templateFilePath and the packageName of the
// service and returns what the relative file path of what should be written to
// disk
func templatePathToActual(templFilePath, packageName string) string {
	// Switch "NAME" in path with packageName.
	// i.e. for packageName = addsvc; /NAME-service/NAME-server -> /addsvc-service/addsvc-server
	actual := strings.Replace(templFilePath, "NAME", packageName, -1)

	actual = strings.TrimSuffix(actual, "template")

	return actual
}

type defaultRender struct{}

func (_ defaultRender) Render(f string, te *gengokit.TemplateExecutor) (io.Reader, error) {
	c, err := applyTemplateFromPath(f, te)
	if err != nil {
		return nil, errors.Wrap(err, "cannot render template")
	}

	return c, nil
}

// applyTemplateFromPath calls applyTemplate with the template at templFilePath
func applyTemplateFromPath(templFilePath string, executor *gengokit.TemplateExecutor) (io.Reader, error) {
	templBytes, err := templFiles.Asset(templFilePath)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to find template file: %v", templFilePath)
	}

	return applyTemplate(templBytes, templFilePath, executor)
}

func applyTemplate(templBytes []byte, templName string, executor *gengokit.TemplateExecutor) (io.Reader, error) {
	templateString := string(templBytes)

	codeTemplate, err := template.New(templName).Funcs(executor.FuncMap).Parse(templateString)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create template")
	}

	outputBuffer := bytes.NewBuffer(nil)
	err = codeTemplate.Execute(outputBuffer, executor)
	if err != nil {
		return nil, errors.Wrap(err, "template error")
	}

	return outputBuffer, nil
}

// formatCode takes a string representing golang code and attempts to return a
// formated copy of that code.  If formatting fails, a warning is logged and
// the original code is returned.
func formatCode(code []byte) []byte {
	formatted, err := format.Source(code)

	if err != nil {
		log.WithError(err).Warn("Code formatting error, generated service will not build, outputting unformatted code")
		// return code so at least we get something to examine
		return code
	}

	return formatted

}
