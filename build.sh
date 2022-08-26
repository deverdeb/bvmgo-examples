#!/bin/sh

go mod download

go build -o bin/ ./net/socket/server/server.go
go build -o bin/ ./net/socket/client/client.go

go build -o bin/ ./file/embed/embed.go
