.DEFAULT_GOAL := run-docker-container

.PHONY: run
run:
	@go run cmd/server/main.go

.PHONY: run-bin
run-bin:
	@./log-parser

.PHONY: clean
clean:
	@rm -f log-parser

.PHONY: build-linux
build-linux: clean
	@GOOS=linux GOARCH=amd64 go build -o log-parser cmd/server/main.go

.PHONY: build-macos
build-macos: clean
	@GOOS=darwin GOARCH=amd64 go build -o log-parser cmd/server/main.go

.PHONY: build-docker-image
build-docker-image:
	@docker build -t log-parser .

.PHONY: run-docker-container
run-docker-container: build-docker-image
	@docker run -it log-parser

.PHONY: burn
burn:
	@echo "--- Kill all Docker Containers and Images locally ---"
	@docker stop $$(docker ps -a -q)
	@docker rm $$(docker ps -a -q)
	@docker rmi $$(docker images -a -q)
