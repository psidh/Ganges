#!/bin/bash

echo "Formatting Go Code"

go fmt ./...

echo "Running Lexer Test":

go test ./src/lexer

echo "Running Parser Test":

go test ./src/parser

echo "Running Evaluator Test":

go test ./src/eval

echo "Running AST Test":

go test ./src/ast