.PHONY: run test test-cov

run:
	go run app/main.go

test:
	go test -v ./...

test-cov:
	go test -cover -coverprofile=cover.out -v ./... && go tool cover -html=cover.out -o cover.html