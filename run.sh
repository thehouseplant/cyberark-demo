#!/bin/sh
x=1
while [ $x -le 30 ]
do
    echo "Running CyberArk demo application for 5s..."
    go run main.go
    sleep 5
    x=$(( $x + 1 ))
    echo "Killing application"
done
