# bechamel-api

Bechamel: a REST API for Lasagna Love request, requester and volunteer management. This is part of [Lasagna Love's Project Ricotta](https://lasagnalove.org/).

For more details on Project Ricotta and our contributor guidelines, please see the [project-ricotta repository](https://github.com/Lasagna-Love-Portal/project-ricotta).

## Getting Started

### Step 1: Install Go (if not already installed)

Download Go from the official site: [https://golang.org/dl/](https://golang.org/dl/)

Unix-based systems can use package managers:

```bash
brew install go # MacOS
sudo apt-get install golang # Ubuntu, Debian, etc.
```

### Step 2: Verify Go Installation

Confirm your Go installation:

```bash
go version
```

### Step 3: Run

```bash
cd path/to/project
go mod download # downloads dependencies
go run .  # runs all
```

This starts a server on `port 8080`. A message appears:

```bash
[GIN-debug] Listening and serving HTTP on :8080
```

Interact with the server by sending HTTP requests, like:

```bash
curl localhost:8080
```

Try `http://localhost:8080` in your web browser. To stop the server, press `Ctrl+C`.

### Step 4: Build

```bash
go build .
```

This creates a binary file. Run it:

```bash
./bechamel-api
```

## Troubleshooting

If issues arise, check:

Your Go installation:

```bash
go version
```

Your dependencies:

```bash
go list -m all
```

And tidy up:

```bash
go mod tidy
```

### Run all project tests

```bash
go test -v -cover -race -bench . ./...
```

or,

```bash
go test ./...
```

## Inspect which code parts aren't covered by tests

Step 1 - run tests with `-coverprofile` flag

```bash
go test -coverprofile=coverage.out ./...
```

Coverage Report

```bash
go tool cover -html=coverage.out
```

### Contribution

Please see [project-ricotta](https://github.com/Lasagna-Love-Portal/project-ricotta) for more details on contributor guidelines. Contributions are welcomed. Please submit a PR with your changes and they will be reviewed as soon as possible.

For an examples of how to add new tests, see the [a quick testing guide](./documentation/TESTING.md).

### License

This project is licensed under the terms of the license found in the file [`LICENSE`](LICENSE) in the root directory of this project.

### Contact

For any queries or concerns, please open an [issue](https://github.com/Lasagna-Love-Portal/bechamel-api/issues) in the repository or send us an email: [info@lasagnalove.org](mailto:info@lasagnalove.org)

## Acknowledgements

[Lasagna Love ](https://lasagnalove.org/)
