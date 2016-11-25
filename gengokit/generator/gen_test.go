package generator

import (
	"bytes"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"

	"github.com/TuneLab/go-truss/gengokit"
	templateFileAssets "github.com/TuneLab/go-truss/gengokit/template"
	"github.com/TuneLab/go-truss/svcdef"

	"github.com/TuneLab/go-truss/gengokit/gentesthelper"
)

var gopath []string

func init() {
	gopath = filepath.SplitList(os.Getenv("GOPATH"))
}

func init() {
	log.SetLevel(log.DebugLevel)
	_ = errors.Wrap
}

func TestTemplatePathToActual(t *testing.T) {
	pathToWants := map[string]string{
		"NAME-service/":                "package-service/",
		"NAME-service/test.gotemplate": "package-service/test.go",
		"NAME-service/NAME-server":     "package-service/package-server",
	}

	for path, want := range pathToWants {
		if got := templatePathToActual(path, "package"); got != want {
			t.Fatalf("\n`%v` got\n`%v` wanted", got, want)
		}
	}
}

func TestApplyTemplateFromPath(t *testing.T) {
	const def = `
		syntax = "proto3";

		// General package
		package general;

		import "google.golang.org/genproto/googleapis/api/serviceconfig/annotations.proto";

		// RequestMessage is so foo
		message RequestMessage {
			string input = 1;
		}

		// ResponseMessage is so bar
		message ResponseMessage {
			string output = 1;
		}

		// ProtoService is a service
		service ProtoService {
			// ProtoMethod is simple. Like a gopher.
			rpc ProtoMethod (RequestMessage) returns (ResponseMessage) {
				// No {} in path and no body, everything is in the query
				option (google.api.http) = {
					get: "/route"
				};
			}
		}
	`
	sd, err := svcdef.NewFromString(def, gopath)
	if err != nil {
		t.Fatal(err)
	}

	conf := gengokit.Config{
		GoPackage: "github.com/TuneLab/go-truss",
		PBPackage: "github.com/TuneLab/go-truss/gengokit/general-service",
	}

	te, err := gengokit.NewExecutor(sd, conf)
	if err != nil {
		t.Fatal(err)
	}

	end, err := applyTemplateFromPath("NAME-service/generated/endpoints.gotemplate", te)
	if err != nil {
		t.Fatal(err)
	}

	endCode, err := ioutil.ReadAll(end)
	if err != nil {
		t.Fatal(err)
	}

	_, err = testFormat(endCode)
	if err != nil {
		t.Fatal(err)
	}
}

func svcMethodsNames(methods []*svcdef.ServiceMethod) []string {
	var mNames []string
	for _, m := range methods {
		mNames = append(mNames, m.Name)
	}

	return mNames
}

func stringToTemplateExector(def, importPath string) (*gengokit.Executor, error) {
	sd, err := svcdef.NewFromString(def, gopath)
	if err != nil {
		return nil, err
	}

	conf := gengokit.Config{
		GoPackage: importPath,
		PBPackage: importPath,
	}

	te, err := gengokit.NewExecutor(sd, conf)
	if err != nil {
		return nil, err
	}

	return te, nil
}

func TestAllTemplates(t *testing.T) {
	const goPackage = "github.com/TuneLab/go-truss/gengokit"
	const goPBPackage = "github.com/TuneLab/go-truss/gengokit/general-service"

	const def = `
		syntax = "proto3";

		// General package
		package general;

		import "google.golang.org/genproto/googleapis/api/serviceconfig/annotations.proto";

		// RequestMessage is so foo
		message RequestMessage {
			string input = 1;
		}

		// ResponseMessage is so bar
		message ResponseMessage {
			string output = 1;
		}

		// ProtoService is a service
		service ProtoService {
			// ProtoMethod is simple. Like a gopher.
			rpc ProtoMethod (RequestMessage) returns (ResponseMessage) {
				// No {} in path and no body, everything is in the query
				option (google.api.http) = {
					get: "/route"
				};
			}
		}
	`

	const def2 = `
		syntax = "proto3";

		// General package
		package general;

		import "google.golang.org/genproto/googleapis/api/serviceconfig/annotations.proto";

		// RequestMessage is so foo
		message RequestMessage {
			string input = 1;
		}

		// ResponseMessage is so bar
		message ResponseMessage {
			string output = 1;
		}

		// ProtoService is a service
		service ProtoService {
			// ProtoMethod is simple. Like a gopher.
			rpc ProtoMethod (RequestMessage) returns (ResponseMessage) {
				// No {} in path and no body, everything is in the query
				option (google.api.http) = {
					get: "/route"
				};
			}
			// ProtoMethodAgain is simple. Like a gopher again.
			rpc ProtoMethodAgain (RequestMessage) returns (ResponseMessage) {
				// No {} in path and no body, everything is in the query
				option (google.api.http) = {
					get: "/route2"
				};
			}
		}
	`

	sd1, err := svcdef.NewFromString(def, gopath)
	if err != nil {
		t.Fatal(err)
	}

	sd2, err := svcdef.NewFromString(def, gopath)
	if err != nil {
		t.Fatal(err)
	}

	conf := gengokit.Config{
		GoPackage: "github.com/TuneLab/go-truss/gengokit",
		PBPackage: "github.com/TuneLab/go-truss/gengokit/general-service",
	}

	executor1, err := gengokit.NewExecutor(sd1, conf)
	if err != nil {
		t.Fatal(err)
	}

	executor2, err := gengokit.NewExecutor(sd2, conf)
	if err != nil {
		t.Fatal(err)
	}

	for _, templFP := range templateFileAssets.AssetNames() {
		var prev io.Reader

		firstCode, err := testGenerateResponseFile(executor1, prev, templFP)
		if err != nil {
			t.Fatalf("%v failed to format on first generation\n\nERROR:\n\n%v\n\nCODE:\n\n%v", templFP, err.Error(), string(firstCode))
		}

		// store the file to pass back to testGenerateResponseFile for second generation
		prev = bytes.NewReader(firstCode)

		secondCode, err := testGenerateResponseFile(executor1, prev, templFP)
		if err != nil {
			t.Fatalf("%v failed to format on second identical generation\n\nERROR: %v\nCODE:\n\n%v",
				templFP, err.Error(), string(secondCode))
		}

		if bytes.Compare(firstCode, secondCode) != 0 {
			t.Fatal("Generated code differs after regeneration with same definition\n" + diff(string(firstCode), string(secondCode)))
		}

		// store the file to pass back to testGenerateResponseFile for third generation
		prev = bytes.NewReader(secondCode)

		// pass in Executor created from def2
		addRPCCode, err := testGenerateResponseFile(executor2, prev, templFP)
		if err != nil {
			t.Fatalf("%v failed to format on third generation with 1 rpc added\n\nERROR: %v\nCODE:\n\n%v",
				templFP, err.Error(), string(addRPCCode))
		}

		// store the file to pass back to testGenerateResponseFile for forth generation
		prev = bytes.NewReader(addRPCCode)

		// pass in Executor create from def1
		_, err = testGenerateResponseFile(executor1, prev, templFP)
		if err != nil {
			t.Fatalf("%v failed to format on forth generation with 1 rpc removed\n\nERROR: %v\nCODE:\n\n%v",
				templFP, err.Error(), string(addRPCCode))
		}
	}
}

func diff(a, b string) string {
	return gentesthelper.DiffStrings(
		a,
		b,
	)
}

func testGenerateResponseFile(executor *gengokit.Executor, prev io.Reader, templFP string) ([]byte, error) {
	// apply server_handler.go template
	code, err := generateResponseFile(executor, prev, templFP)
	if err != nil {
		return nil, err
	}

	// read the code off the io.Reader
	codeBytes, err := ioutil.ReadAll(code)
	if err != nil {
		return nil, err
	}

	// format the code
	formatted, err := testFormat(codeBytes)
	if err != nil {
		return codeBytes, err
	}

	return formatted, nil
}

// testFormat takes a string representing golang code and attempts to return a
// formated copy of that code.
func testFormat(code []byte) ([]byte, error) {
	formatted, err := format.Source(code)

	if err != nil {
		return code, err
	}

	return formatted, nil
}
