#!/usr/bin/env bash

go build -o deadlock.exe main.go
./deadlock.exe
# https://blog.golang.org/race-detector
#go run -race main.go