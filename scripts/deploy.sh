#!/bin/bash
set -e  # Exit immediately if a command exits with a non-zero status

FUNCTION_NAME="user-management-lambda"

MAIN_FILE="main.go"

BINARY_NAME="bootstrap"

echo "Compiling Go Lambda..."
GOOS=linux GOARCH=amd64 go build -o $BINARY_NAME $MAIN_FILE

echo "Zipping binary..."
zip -j function.zip $BINARY_NAME

echo "Uploading to AWS Lambda..."
aws lambda update-function-code \
    --function-name $FUNCTION_NAME \
    --zip-file fileb://function.zip

echo "Deployment completed!"
