test: # Run unit tests with coverage
	mkdir -p coverage
	go test -v -tags=unit -covermode=count -coverprofile coverage/coverage.out ./...
	go tool cover -func coverage/coverage.out

vendor:
	rm -rf vendor
	go mod tidy -v
	go mod download
	go mod vendor
