#!/usr/bin/env bash

go build -o http.exe main.go

WORKDIR="$PWD"

gnome-terminal --working-directory="$WORKDIR" --command 'bash ./server.sh'

x-www-browser http://localhost:8081/test-port-8081
x-www-browser http://localhost:8082/test-port-8082
x-www-browser http://localhost:8083/test-port-8083
x-www-browser http://localhost:8084/test-port-8084