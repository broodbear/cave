.PHONY: cave lint

cave:
	go build -o ./bin/cave ./cmd/cave

lint:
	@golangci-lint -v run
