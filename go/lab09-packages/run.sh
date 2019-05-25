#!/usr/bin/env bash

go build -o packages.exe main.go
./packages.exe

printf "======"
printf "\nworld package documentation (look at run.sh)\n"
printf "======\n"

go doc ./world
go doc ./world.PrintStartRoom

printf "Open => http://localhost:6060"
#godoc -http=:6060
#ss -lptn 'sport = :6060'