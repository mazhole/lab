#!/usr/bin/env bash

go build -o jsonapi.exe main.go

WORKDIR="$PWD"

gnome-terminal --working-directory="$WORKDIR" --command 'bash ./server.sh'

x-www-browser http://localhost:8081
