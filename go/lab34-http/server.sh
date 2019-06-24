#!/usr/bin/env bash

./http.exe -port 8081 &
./http.exe -port 8082 &
./http.exe -port 8083 &
./http.exe -port 8084 &
echo '=> Servers are running'
wait