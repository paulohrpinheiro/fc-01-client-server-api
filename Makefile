createdb:
	sqlite3 exchange.sqlite3 'CREATE TABLE IF NOT EXISTS exchange (bid DECIMAL(10, 5), t TIMESTAMP DEFAULT CURRENT_TIMESTAMP);'

server: createdb
	go run server/cmd/server.go

client:
	go run client/cmd/client.go


.PHONY: createdb server client
