all: test

test: httptest 

httptest:
	$(MAKE) -C http

clitest:
	go test -tags=integration -v ./cli

