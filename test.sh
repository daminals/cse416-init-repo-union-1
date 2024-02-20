#!/bin/bash

./compile.sh # build proto files

# Function to run client and capture output
run_client() {
    client_name=$1
    client_path=$2

    # Run client and capture output
    echo "$client_name Client:"
    go run "$client_path" 2>&1 | sed "s/^/$client_name: /" &
}

# Run market client
run_client "Market" "market/mock.go"

sleep 5

# Run consumer client
run_client "Consumer" "consumer/consumer.go"

sleep 5

# Run producer client
run_client "Producer" "producer/producer.go"

# Wait for all clients to finish
wait
