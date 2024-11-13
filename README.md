# Go Tutorial
![Go Logo](./go.png)

## Table of Contents

- [Project Overview](#project-overview)
- [Directory Structure](#directory-structure)
- [How to Run](#how-to-run)
  - [Running Individual Projects](#running-individual-projects)
  - [Running the Bank API Project](#running-the-bank-api-project)
  - [Running the Memory Test](#running-the-memory-test)
- [Dependencies](#dependencies)
- [Contributing](#contributing)
- [License](#license)

## Project Overview

This repository contains various Go projects and exercises, each focusing on different aspects of the Go programming language. The projects are organized into directories, each with its own `main.go` file or other relevant files.


## Project Descriptions

### 1. Hello-World
This directory contains a simple "Hello, World!" program to get started with Go.

### 2. Variables
This directory demonstrates variable declarations, initializations, and usage in Go.

### 3. IF-else-Switch-Case
This directory contains examples of conditional statements, including `if-else` and `switch-case`.

### 4. Arrays
This directory demonstrates the usage of arrays in Go.

### 5. Strings-Runes-Bytes
This directory contains examples of working with strings, runes, and bytes in Go. It includes string concatenation using a `strings.Builder`.

### 6. Structs-Interfaces
This directory demonstrates the usage of structs and interfaces in Go.

### 7. Pointers
This directory contains examples of working with pointers in Go.

### 8. Go-Routines
This directory demonstrates the usage of goroutines for concurrent programming in Go.

### 9. Channels
This directory contains examples of using channels for communication between goroutines.

### 10. Generics
This directory demonstrates the usage of generics in Go.

### Bank-api-project
This is a more complex project that implements a simple API for managing bank accounts. It includes the following components:
- **api/api.go**: Defines the API request and response structures and error handling functions.
- **cmd/api/main.go**: The entry point for the API server.
- **internal/handlers**: Contains the API handlers, including `api.go` and `get_coin_balance.go`.
- **internal/middleware**: Contains middleware functions, including `authorization.go`.
- **internal/tools**: Contains utility functions and database interactions, including `database.go` and `mockdb.go`.
- **go.mod** and **go.sum**: Go module files for dependency management.

### Memory-test
This directory contains a performance test for measuring the time taken to append elements to a slice with and without pre-allocation.

## How to Run

### Running Individual Projects
To run any of the individual projects, navigate to the respective directory and execute the `main.go` file using the `go run` command. For example:

```sh
cd 1.Hello-World
go run main.go
```

### Running the Bank API Project
To run the Bank API project, navigate to the `cmd/api` directory and execute the `main.go` file:

```sh
cd Bank-api-project/cmd/api
go run main.go
```

### Running the Memory Test
To run the memory test, navigate to the `Memory-test` directory and execute the `speedtest.go` file:

```sh
cd Memory-test
go run speedtest.go
```

## Dependencies
The Bank API project uses external dependencies, which are managed using Go modules. Ensure you have Go installed and run the following command in the `Bank-api-project` directory to download the dependencies:

```sh
go mod tidy
```

## Contributing
Feel free to contribute to this repository by submitting pull requests. Ensure that your code follows Go conventions and includes appropriate tests.


## License
This project is licensed under the MIT License. See the LICENSE file for details.

