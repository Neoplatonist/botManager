package = botmanager

all: proto buildServer buildClient

server:
	@go run ./cmd/server/*.go

client:
	@go run ./cmd/client/*.go

proto:
	@protoc --go_out=plugins=grpc:. ./services/cli/proto/*.proto

buildServer: 
	@go build -o botmanager-server ./cmd/server

buildClient: 
	@go build -o botmanager-client ./cmd/client

prod:
	@if [ ! $(docker ps -q -f name=$(package)) ]; then \
		make prodRM; \
	fi
	@docker build \
		--build-arg DISCORD_APP_USER=$(DISCORD_APP_USER) \
		-t $(package) .
	@docker run -d --name $(package) $(package)

prodRM:
	@docker stop $(package)
	@docker rm $(package)

	# Discord Bot Fix & Rewrite: Golang