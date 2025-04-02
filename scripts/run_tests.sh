#!/bin/bash

echo "Running tests..."
go test ./test/... -v
echo "Tests completed."