.PHONY: help
help:
	@echo 'Usage'
	@echo '    make run'
	@echo '    make build'
	@echo '    make tidy'
	@echo '    make clean'

.PHONY: run
run:
	@go run main.go

.PHONY: build
build:
	@CGO_ENABLED=0 go build -o main main.go

.PHONY: tidy
tidy:
	GOSUMDB=off go mod tidy

.PHONY: clean
clean:
	@rm -rf main
