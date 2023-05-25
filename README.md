# bechamel-api

Bechamel: a REST API for Lasagna Love request, requester and volunteer management. This is part of [Lasagna Love's Project Ricotta](https://lasagnalove.org/).

For more details on Project Ricotta and our contributor guidelines, please see the [project-ricotta repository](https://github.com/Lasagna-Love-Portal/project-ricotta).

## Getting Started

### Step 1: Install Go (if not already installed)

Download Go from the official website: [https://golang.org/dl/](https://golang.org/dl/)

For Unix-based systems, you can use package managers:

```bash
brew install go # MacOS
sudo apt-get install golang # Ubuntu, Debian, etc.
```

### Step 2: Verify Go Installation

Check your Go installation:

```bash
go version
```

### Step 3: Run

```bash
cd path/to/your/project
go mod download # downloads all dependencies
go run .  # run all
```

It should start a server on `port 8080`. You'll see a message in the terminal:

```bash
[GIN-debug] Listening and serving HTTP on :8080
```

Interact with the server by sending HTTP requests to the endpoints, for example:

```bash
curl localhost:8080
```

Also, you can try opening your web browser and typing `http://localhost:8080` into the address bar.

To stop the server, use `Ctrl+C` in the terminal.

### Step 4: Build

```bash
go build .
```

This will create a binary file in the current directory. You can run it with:

```bash
./bechamel-api
```

## Troubleshooting

If you encounter any issues, make sure to check:

Your Go installation with

```bash

go version
```

Your dependencies with

```bash
go list -m all
```

Your dependencies with

```bash
go mod tidy
```

### Contribution

Please see [project-ricotta](https://github.com/Lasagna-Love-Portal/project-ricotta) for more details on contributor guidelines. Contributions are welcomed. Please submit a PR with your changes and they will be reviewed as soon as possible.

### License

This project is licensed under the terms of the license found in the file [`LICENSE`](LICENSE) in the root directory of this project.

### Contact

For any queries or concerns, please open an issue in the repository or send us a [pecorino-email @lasagnalove](info@lasagnalove.org)

## Acknowledgements

[Lasagna Love ](https://lasagnalove.org/)
