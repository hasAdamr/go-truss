package test

import (
	"testing"

	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	// 3d Party
	"golang.org/x/net/context"

	// Go Kit
	"github.com/go-kit/kit/log"

	// This Service
	pb "github.com/TuneLab/go-truss/truss/_integration-tests/http/httptest-service"
	svc "github.com/TuneLab/go-truss/truss/_integration-tests/http/httptest-service/generated"
	httpclient "github.com/TuneLab/go-truss/truss/_integration-tests/http/httptest-service/generated/client/http"
	handler "github.com/TuneLab/go-truss/truss/_integration-tests/http/httptest-service/handlers/server"
	//"github.com/davecgh/go-spew/spew"
)

var httpserver *httptest.Server

func init() {
	var logger log.Logger
	logger = log.NewNopLogger()

	var service handler.Service
	{
		service = handler.NewService()
	}

	// Endpoint domain.
	getWithQueryE := svc.MakeGetWithQueryEndpoint(service)
	getWithRepeatedQueryE := svc.MakeGetWithRepeatedQueryEndpoint(service)

	endpoints := svc.Endpoints{
		GetWithQueryEndpoint:         getWithQueryE,
		GetWithRepeatedQueryEndpoint: getWithRepeatedQueryE,
	}

	ctx := context.Background()

	h := svc.MakeHTTPHandler(ctx, endpoints, logger)
	httpserver = httptest.NewServer(h)
}

func TestGetWithQueryClient(t *testing.T) {
	var req pb.GetWithQueryRequest
	req.A = 12
	req.B = 45360
	want := req.A + req.B

	svchttp, err := httpclient.New(httpserver.URL)
	if err != nil {
		t.Fatalf("failed to create httpclient: %q", err)
	}

	resp, err := svchttp.GetWithQuery(context.Background(), &req)
	if err != nil {
		t.Fatalf("httpclient returned error: %q", err)
	}

	if resp.V != want {
		t.Fatalf("Expect: %d, got %d", want, resp.V)
	}
}

func TestGetWithQueryRequest(t *testing.T) {
	var A, B int64
	A = 12
	B = 45360
	want := A + B

	route := fmt.Sprintf("%s/%s?%s=%d&%s=%d", httpserver.URL, "getwithquery", "A", A, "B", B)
	httpReq, err := http.NewRequest("GET", route, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		t.Fatal(err)
	}
	defer httpResp.Body.Close()

	var resp pb.GetWithQueryResponse

	respBytes, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.V != want {
		t.Fatalf("Expect: %d, got %d", want, resp.V)
	}
}

func TestGetWithRepeatedQueryClient(t *testing.T) {
	var req pb.GetWithRepeatedQueryRequest
	req.A = []int64{12, 45360}
	want := req.A[0] + req.A[1]

	svchttp, err := httpclient.New(httpserver.URL)
	if err != nil {
		t.Fatalf("failed to create httpclient: %q", err)
	}

	resp, err := svchttp.GetWithRepeatedQuery(context.Background(), &req)
	if err != nil {
		t.Fatalf("httpclient returned error: %q", err)
	}

	if resp.V != want {
		t.Fatalf("Expect: %d, got %d", want, resp.V)
	}
}

func TestGetWithRepeatedQueryRequest(t *testing.T) {
	var A []int64
	A = []int64{12, 45360}
	want := A[0] + A[1]

	route := fmt.Sprintf("%s/%s?%s=%d&%s=%d", httpserver.URL, "getwithrepeatedquery", "A", A[0], "A", A[1])
	httpReq, err := http.NewRequest("GET", route, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		t.Fatal(err)
	}
	defer httpResp.Body.Close()

	var resp pb.GetWithRepeatedQueryResponse

	respBytes, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.V != want {
		t.Fatalf("Expect: %d, got %d", want, resp.V)
	}
}
