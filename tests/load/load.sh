#!/bin/bash

echo "Starting load test... will take 1m"

echo 'GET http://localhost:8081/bins' | \
    vegeta attack -rate 200 -duration 1m | vegeta encode > ./tests/load/results.json

vegeta report ./tests/load/results.*
