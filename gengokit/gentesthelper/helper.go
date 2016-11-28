package gentesthelper

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"reflect"
	"runtime"
	"strings"

	"github.com/pmezard/go-difflib/difflib"
)

// FuncSourceCode returns a string representing the source code of the function
// provided to it.
func FuncSourceCode(val interface{}) (string, error) {
	ptr := reflect.ValueOf(val).Pointer()
	fpath, _ := runtime.FuncForPC(ptr).FileLine(ptr)

	funcName := runtime.FuncForPC(ptr).Name()
	parts := strings.Split(funcName, ".")
	funcName = parts[len(parts)-1]

	// Parse the go file into the ast
	fset := token.NewFileSet()
	fileAst, err := parser.ParseFile(fset, fpath, nil, parser.ParseComments)
	if err != nil {
		return "", fmt.Errorf("ERROR: go parser couldn't parse file '%v'\n", fpath)
	}

	// Search ast for function declaration with name of function passed
	var fAst *ast.FuncDecl
	for _, decs := range fileAst.Decls {
		if f, ok := decs.(*ast.FuncDecl); ok && f.Name.String() == funcName {
			fAst = f
			break
		}
	}
	code := bytes.NewBuffer(nil)
	err = printer.Fprint(code, fset, fAst)

	if err != nil {
		return "", fmt.Errorf("couldn't print code for func %q: %v\n", funcName, err)
	}

	return code.String(), nil
}

// DiffStrings returns the line differences of two strings. Useful for
// examining how generated code differs from expected code.
// Callers should call TestFormat on code first compare formatted code
func DiffStrings(want, got string) string {
	t := difflib.UnifiedDiff{
		A:        difflib.SplitLines(want),
		B:        difflib.SplitLines(got),
		FromFile: "Want",
		ToFile:   "Got",
		Context:  5,
	}
	text, _ := difflib.GetUnifiedDiffString(t)
	return text
}

// TestFormat takes a string representing golang code and attempts to return a
// formated copy of that code.
func TestFormat(code string) (string, error) {
	code = strings.TrimSpace(code)
	formatted, err := format.Source([]byte(code))

	if err != nil {
		return code, err
	}

	return string(formatted), nil
}
