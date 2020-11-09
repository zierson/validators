SHELL=/bin/bash

default:
	@echo 'Usage:'
	@echo '$$ make test'
	@echo '$$ make fmt'

test:
	@go test

fmt:
	@go fmt *.go
