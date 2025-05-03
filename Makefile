.PHONY: run build

run:
	go run ./cmd/promptsentry

build:
	go build -o promptsentry ./cmd/promptsentry

%:
	go run ./cmd/promptsentry $@