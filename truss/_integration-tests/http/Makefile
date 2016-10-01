all: test

test:
	truss httptest.proto
	cp -r handlers httptest-service
	go test
	$(MAKE) clean

clean:
	rm -rf httptest-service
	rm -rf third_party
