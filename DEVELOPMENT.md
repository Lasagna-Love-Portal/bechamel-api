# Contributing Guide

## Initial Setup

### 1. Install Go (if not installed)

To download Go, visit [the official site](https://golang.org/dl/). For Unix-based systems, use package managers:

```bash
brew install go # MacOS
sudo apt-get install golang # Linux
```

### 2. Confirm Go Installation

Verify your Go installation using the following command:

```bash
go version
```

### 3. Run the Project

Navigate to the project directory and execute these commands:

```bash
cd path/to/project
go mod download # download dependencies
go run .  # execute the program
```

This initiates a server on `port 8080`. You should see:

```bash
[GIN-debug] Listening and serving HTTP on :8080
```

Interact with the server by sending HTTP requests:

```bash
curl localhost:8080
```

Or, open `http://localhost:8080` in a web browser. Stop the server with `Ctrl+C`.

### 4. Build the Project

To create a binary file, use:

```bash
go build .
```

Execute the binary with:

```bash
./bechamel-api
```

## Troubleshooting

Encounter issues? Check your Go installation:

```bash
go version
```

Inspect your dependencies:

```bash
go list -m all
```

And tidy up:

```bash
go mod tidy
```

### Running Tests

To execute all project tests, use:

```bash
go test -v -cover -race -bench . ./...
```

or simply,

```bash
go test ./...
```

## Code Coverage

To inspect which code segments aren't covered by tests, follow these steps:

1. Run tests with `-coverprofile` flag

```bash
go test -coverprofile=coverage.out ./...
```

2. Generate coverage report

```bash
go tool cover -html=coverage.out
```

## Writing Test Functions

To add a new function for testing:

1. **Identify the Function**: Determine which function you want to test. 

2. **Create Test Cases**: Define an array of `testType` structs for the test cases. Each struct should include:

    - `name`: A descriptive name of the test case.
    - `call`: A closure that calls the function to test with case-specific parameters.
    - `wantErr`: A boolean indicating if an error is expected.

   Example:
   ```go
   tests := []testType{
		{
			name: "Your Test Name",
			call: func() (model.LasagnaLoveUser, error) {
				return YourFunctionToTest(param1, param2)
			},
			wantErr: false,
		},
   }
   ```

3. **Execute `runTests`**: Call `runTests()` with the `testing.T` instance and your test cases array.

   Example:
   ```go
   runTests(t, tests)
   ```

This function validates each test case, ensuring expected and actual error states align.

## References
For additional insight, refer to [internal/user_access_test.go](../internal/user_access_test.go).