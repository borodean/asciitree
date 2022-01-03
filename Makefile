bench:
	go test -bench .

test:
	go test -v ./...

test-ci:
	go test -coverprofile=coverage.out -race -v ./...
