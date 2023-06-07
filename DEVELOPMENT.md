# Contributing Guide


## Contents

- [First-time Setup](#first-time-setup)
   - [Go Language Installation (if absent)](#1-go-language-installation-if-absent)
   - [Verification of Go Installation](#2-verification-of-go-installation)
   - [Project Execution](#3-project-execution)
   - [Build the Project](#4-build-the-project)
- [Troubleshooting](#troubleshooting)
   - [Running Tests](#running-tests)
   - [Code Coverage](#code-coverage)
   - [Writing Test Functions](#writing-test-functions)
- [Pull Request Review Process: How-to-fetch and run a PR?](#pull-request-review-process)


## First-time Setup

### 1. Go Language Installation (if absent)

Download Go by visiting [its official site](https://golang.org/dl/). On Unix-based systems, employ package managers:

```bash
brew install go # MacOS
sudo apt-get install golang # Linux
```

### 2. Verification of Go Installation

Confirm your Go installation via this command:

```bash
go version
```

### 3. Project Execution

Reach the project directory and perform these commands:

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

Alternatively, open `http://localhost:8080` on a web browser. Cease the server with `Ctrl+C`.

### 4. Build the Project

To generate a binary file, employ:

```bash
go build .
```

Execute the binary with:

```bash
./bechamel-api
```

## Troubleshooting

Experiencing problems? Recheck your Go installation:

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

To conduct all project tests, employ:

```bash
go test -v -cover -race -bench . ./...
```

or simply,

```bash
go test ./...
```

## Code Coverage

To identify untested code segments, follow these steps:

1. Run tests with `-coverprofile` flag

```bash
go test -coverprofile=coverage.out ./...
```

2. Generate coverage report

```bash
go tool cover -html=coverage.out
```

## Writing Test Functions

To introduce a new testing function:

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

This function verifies each test case, ascertaining expected and actual error states coincide.

## Pull Request Review Process: How-to-fetch and run a PR?

1. Copy the repository:
   - Execute:
     ```bash
     git clone https://github.com/Lasagna-Love-Portal/bechamel-api.git
     ```

2. Reach the repository:
   - Execute:
     ```bash
     cd bechamel-api
     ```

3. Fetch and switch to the pull request (e.g., PR#18):
   - Execute:
    ```bash
    git fetch upstream pull/18/head:PR18
    ```
    ```bash
    git checkout PR18
    ```

4. Install dependencies:


   - For Go projects, execute:
     ```bash
     go mod tidy
     ```

5. Compile/Run the project:
   - For Go projects, execute:
     ```bash
     go build
     ```
     or
     ```bash
     go run main.go
     ```

6. Assess the modifications:
   - Conduct unit tests and manual testing.

7. Perform unit tests:
   - Execute:
     ```bash
     go test ./...
     ```

8. Examine the alterations:
   - Inspect the code modifications and verify they adhere to coding and documentation standards.

9. Carry out manual testing:
   - Operate the program and test its functionalities.

10. Offer feedback:
    - Post comments on the GitHub page of the pull request.

11. Revert to the main branch:
    - Execute:
      ```bash
      git checkout main
      ```

12. Erase the PR branch (if redundant):
    - Execute:
      ```bash
      git branch -d PR18
      ```

Note: Modify the URL, repository name, and pull request number as needed. Ensure git and golang are installed. Adjust these steps in line with your project's prerequisites.