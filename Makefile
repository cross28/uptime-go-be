GO=go

.PHONY: run
run:
	$(GO) run ./cmd/app -env dev -port 8000