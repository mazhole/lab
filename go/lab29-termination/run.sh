#!/usr/bin/env bash

go build -o termination.exe main.go
./termination.exe
# https://blog.golang.org/race-detector
#go run -race main.go