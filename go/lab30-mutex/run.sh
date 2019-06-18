#!/usr/bin/env bash

go build -o mutex.exe main.go
./mutex.exe
# https://blog.golang.org/race-detector
#go run -race main.go