# integration-tests

These tests run truss against definition files and tests if the generated code
behaves. 

# ./http

The http directory contains http transport tests. They include tests against
the generated http transport client libraries against the server and testing
hand made http requests against the server.

- Runs truss against `http/httptest.proto`
- Copy `http/handlers` into `http/handlers/httptest-service`
- Run go test

`http/handlers` has implemented handlers for the server. They add things
together for the purposes of testing.

`http_test.go` imports the generated code,
starts up a `httptest` server with the service http handler, and then runs
requests against this server, checking for errors and that the inputted values
add to the outputted values

# ./cli 

The truss cli integration runner does the following tasks:

- Runs truss on each service definition in `cli/test_service_definitions`
- Builds the server and cliclient for each service
- Runs the server
- Runs the cliclient against the server
- Passes if the server and cliclient were able to communicate. Fails if there
  were errors of any kind.

*Test service definition requirements*

Each service definition must have the package name `TEST`, all letters
uppercase.
