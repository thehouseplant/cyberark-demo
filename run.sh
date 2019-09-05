#!/bin/sh
while true; do
    echo "Running CyberArk demo application for 5s..."
    go run main.go
    sleep 5
    echo "Killing application"
done
