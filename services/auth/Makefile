.PHONY: test
.SILENT:

test:
	go test ./... --cover
test-cover:
	go test --coverprofile=coverage.out ./... > /dev/null; \
    go tool cover -func=coverage.out | grep total | grep -oE '[0-9]+(\.[0-9]+)?%'