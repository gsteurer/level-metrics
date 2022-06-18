.PHONY: metrics-cli

default: metrics-cli

metrics-cli:
	go build ./cmd/metrics-cli/metrics-cli.go 

test:
	go test ./... -count=1 -cover

clean:
	rm -rf metrics-cli out.json