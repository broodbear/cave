.PHONY: cave lint

cave:
	go build -o ./bin/cave ./cmd

lint:
	@golangci-lint -v run
