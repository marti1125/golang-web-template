SHELL:=/bin/bash -O extglob

web:
	@clear
	@go run cmd/web/!(*_test).go -E dev