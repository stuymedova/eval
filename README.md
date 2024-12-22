# Eval - Expression Calculator

A web service that evaluates mathematical expressions using the Shunting Yard algorithm for parsing and Reverse Polish Notation (RPN) for evaluation.

## Features

- Supports basic arithmetic operations (+, -, *, /)
- Handles parentheses for expression grouping
- Respects mathematical operator precedence
- Supports decimal numbers

## Installation

Requires Go 1.23.1 or higher.

## Running

```bash
go run ./cmd
```

The server will start on port 8080.

## API

**Endpoint:** POST /api/v1/calculate

**Request Body:**

```json
{
	"expression": "2+2*2"
}
```

**Successful Response (200 OK):**

```json
{
	"result": 6
}
```

**Method not allowed (405 Method Not Allowed):**

```json
{
	"error": "Method not allowed"
}
```

**Invalid request body (422 Unprocessable Entity):**

```json
{
	"error": "Invalid request body"
}
```

**Invalid Expression (422 Unprocessable Entity):**

```json
{
	"error": "Expression is not valid"
}
```

### Examples

Successful calculation:

```bash
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
"expression": "2+2*2"
}'
```

Invalid expression:

```bash
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
"expression": "2+(2*2"
}'
```

## Running tests

```bash
# Test all packages with verbose output
go test -v ./...

# Test handler with verbose output
go test -v ./internal/handler

# Test eval with verbose output
go test -v ./pkg/eval
```
