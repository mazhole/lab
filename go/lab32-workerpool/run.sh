#!/usr/bin/env bash

go build -o workerpool.exe main.go
./workerpool.exe
# https://blog.golang.org/race-detector
#go run -race main.go