#!/bin/bash
while true; do
    echo "Running CyberArk demo application for 5s..."
    ./main
    sleep 5
    echo "Killing application"
done
