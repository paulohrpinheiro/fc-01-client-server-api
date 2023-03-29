server:
	go run server/cmd/server.go

client:
	go run client/cmd/client.go


.PHONY: createdb server client
