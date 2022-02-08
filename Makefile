bench:
	go test -bench .

format:
	golangci-lint run --fix ./...

lint:
	golangci-lint run ./...

test:
	go test -v ./...

test-ci:
	go test -coverprofile=coverage.out -race -v ./...
