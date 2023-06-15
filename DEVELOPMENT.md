# Contributing Guide


## Contents

- [First-time Setup](#first-time-setup)
  - [Go language installation (if absent)](#go-language-installation-if-absent)
  - [Verification of Go installation](#verification-of-go-installation)
- [Developing the Bechamel API](#developing-the-bechamel-api)
  - [API documentation](#api-documentation)
  - [Building the project](#building-the-project)
  - [Running the server locally](#running-the-server-locally)
  - [Running unit tests](#running-unit-tests)
  - [Writing test functions](#writing-unit-tests)
  - [Running the GitHub super-linter](#running-the-linter)
- [Troubleshooting](#troubleshooting)
- [Pull requests](#pull-requests)
  - [Pull request checklist](#pull-request-checklist)
  - [Review process: How to fetch and review a PR?](#review-process-how-to-fetch-and-run-a-pr)


## First-time Setup

### Go language installation (if absent)

Download Go by visiting [its official site](https://golang.org/dl/). On Unix-based systems, employ package managers:

```bash
brew install go # MacOS
sudo apt-get install golang # Linux
```

After installing Go, download the `goimports` package:

```bash
go install golang.org/x/tools/cmd/goimports@latest
```

### Verification of Go installation

Confirm your Go installation via this command:

```bash
go version
```

## Developing the Bechamel API

### API documentation

The Bechamel API documentation can be found in OpenAPI format in the documentation folder in the bechamel-api.yaml file.
You can view this in an OpenAPI - aware file viewer such as Visual Studio Code to browse the Bechamel API.

### Building the project

To verify there are no syntax or compilation issues, you may build the project. From the bechamel-api directory:

```bash
go mod download # download dependencies
go build .
```

### Running the server locally

From a command window in the top-level bechamel-api project directory:

```bash
go run .
```

The Bechamel API server listens on `http://localhost:8080` by default. You should see:

```bash
[GIN-debug] Listening and serving HTTP on :8080
```

You can use HTTP development tools such as curl or Postman to interact with the Bechamel API server.

You may stop the server with `Ctrl+C`.

### Running unit tests

To run the unit tests, from the top-level bechamel-api project directory:

```bash
go test ./...
```

For more verbose output for passing and failing tests, add a '-v' flag to the command.

To identify untested code segments:

1. Run tests with `-coverprofile` flag

```bash
go test -coverprofile=coverage.out ./...
```

2. Generate coverage report

```bash
go tool cover -html=coverage.out
```

### Writing unit tests

For simple unit tests where only a pass/fail state is being tested, you can use the tests and framework
in internal\user_access_test.go as a template for creating your own.
Unit tests associated with a given source file are put in a corresponding \[filename\]_test.go file
To use this template:

1. **Identify the function**: Determine which functions you want to test.

2. **Create test cases**: Define an array of `testType` structs for the test cases. Each struct should include:

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

For more complicated unit testing, you can use the unit tests in the internal\jwt_utils_test.go as a guide.

### Running the linter

The Bechamel API project will be using the GitHub super-linter package for scanning code in pull requests.
It's recommended to run this locally and resolve any issues detected in files you've modified before
putting your pull requests out for review.

You can run this locally using Docker. First, obtain the Super Linter Docker container:

   ```bash
   docker pull github/super-linter:latest
   ```

Then you can use Docker to run the super-linter. To run against all the files in the project directory, from the top level bechamel-api directory:

   ```bash
   docker run -e RUN_LOCAL=true -v .:/tmp/lint github/super-linter
   ```

To narrow the files to lint, pass the file(s) to run as a regular expression
in the environment variable FILTER_REGEX_INCLUDE. Or to exclude file(s) pass an appropriate
regular expression in the FILTER_REGEX_EXCLUDE environment variable.
For example, to only lint the file DEVELOPMENT.md in the top level directory:

   ```bash
   docker run -e RUN_LOCAL=true -e FILTER_REGEX_INCLUDE=DEVELOPMENT.md -v .:/tmp/lint github/super-linter
   ```

You can also use the `super-linter.env` file in the root of the project to store environment variables
to pass to the `docker` invocation. Use the `--env-file` flag to pass the file:

   ```bash
   docker run --env-file super-linter.env -v ./tmp/lint github/super-linter
   ```

See the [super-linter GitHub repository](https://github.com/github/super-linter/blob/main/README.md#filter-linted-files)
for more information on how to specify files for super-linter runs.

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

### Linter specific troubleshooting

#### File not goimports-ed

If you get super-linter errors about files not being goimportes-ed, such as:

   ```bash
   api_profile.go:1: File is not `goimports`-ed (goimports)
   ```

Make sure you have the goimports package as detailed in [Go language installation (if absent)](#go-language-installation-if-absent). Run the goimports program on the file:

   ```bash
   goimports -local [filename] -w .
   ```

This should fix up the ordering of imports in the file.

## Pull requests

### Pull request checklist

Before posting a pull request and requesting reviewers, you may want to make sure you've
done the following:

1. Have you verified that the functionality you've implemented functions as expected?
2. If you added or changed any APIs, did you update the API documentation?
3. If you've added any functionality, did you add appropriate unit tests?
4. Have you verified that existing functionality is working properly?
5. Did you review and proofread your changes to ensure that all your changes are intentional?
6. Did you run 'go mod tidy' to get any changes to dependencies into your pull request?
7. Did you run the linter and resolve any newly found issues in the code your pull request changes?

You can make a pull request and mark the pull request as a draft in order to post the pull request
and allow it to run through the GitHub CI actions, while going through the checklist items above
and making any necessary changes to finalize the pull request.

### Review process: how to fetch and run a PR

If you're new to reviewing pull requests, here's a helpful workflow you can use and adapt to your needs.
In this example, the reviewer is reviewing pull request #18 that is stored in a branch named 'pr-branch':

1. Copy the repository:
     ```bash
     git clone https://github.com/Lasagna-Love-Portal/bechamel-api.git
     ```

2. Reach the repository:
     ```bash
     cd bechamel-api
     ```

3. Fetch and switch to the pull request (e.g., 'pr-branch'):
    ```bash
    git fetch upstream pull/18/head:pr-branch
    ```
    ```bash
    git checkout pr-branch
    ```

4. Install dependencies: see [Building the project](#building-the-project) above

5. Compile the project and ensure it executes: see [Building the project](#building-the-project) and [Running the server locally](#running-the-server-locally) above

6. Verify the unit tests are functioning: see [Running unit tests](#running-unit-tests) above

7. Examine the alterations:
   - Inspect the code modifications and verify they adhere to coding and documentation standards.

8. Carry out manual testing:
   - Operate the program and test its functionalities as appropriate.

9. Offer feedback:
   - Open a review on the GitHub page for the pull request and make comments inline.
     Provide an approval if the changes in the pull request are suitable.
     If changes are required, indicate in the GitHub "Review changes" tab that changes are required.
     Note specifically what changes need to be carried out to unblock the pull request,
     and what changes are suggestions that are not blocking approval.
   - Post general comments about the pull request, if any, on the main GitHub page of the pull request.

Once you're done reviewing, you can switch back to the 'dev' branch (or any other branch you have checked out),
and remove the branch you were reviewing:
   ```bash
   git checkout dev
   git branch -d pr-branch
   ```
