#!/usr/bin/env bash

go build -o select.exe main.go
./select.exe
# https://blog.golang.org/race-detector
#go run -race main.go