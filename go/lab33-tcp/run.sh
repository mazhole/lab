#!/usr/bin/env bash

go build -o tcp.exe main.go

WORKDIR="$HOME/work/mazhole/gitlab.com/lab/go/lab33-tcp"

gnome-terminal --working-directory="$WORKDIR" --command 'bash ./server.sh'

gnome-terminal --working-directory="$WORKDIR" --command 'expect ./client.exp 1 0.5 10'
gnome-terminal --working-directory="$WORKDIR" --command 'expect ./client.exp 2 0.1 20'

## Debug mode
#$ ./test.exp -d params...
