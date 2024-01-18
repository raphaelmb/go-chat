run:
	@go build -o tmp/app cmd/web/*.go && ./tmp/app

.PHONY: run