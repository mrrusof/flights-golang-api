all build:
	go build

e2e:
	E2E=true go test ./...

test:
	go test ./...

clean:
	go clean

run: build
	./flights-golang-api

sample_request:
	curl -vX POST --data-raw '{"path": [{"origin": "ATL", "destination": "EWR"}, {"origin": "SFO", "destination": "ATL"}]}' localhost:8080/flights/summarize; echo

.PHONY: all build test clean run sample_request
