#!/usr/bin/env bash

go build -o atomic.exe main.go
./atomic.exe
# https://blog.golang.org/race-detector
#go run -race main.go