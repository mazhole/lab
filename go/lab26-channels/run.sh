#!/usr/bin/env bash

go build -o channels.exe main.go
./channels.exe
# https://blog.golang.org/race-detector
#go run -race main.go