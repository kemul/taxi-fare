# Taxi Fare Calculator

This project is a taxi fare calculator implemented in Go. It processes records of distances traveled and calculates the total fare based on specified rates. The fare is calculated with different rates for different distance ranges.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [Testing](#testing)
- [Docker](#docker)
- [Project Structure](#project-structure)

## Prerequisites

Before running this project, make sure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.20 or higher)
- [Docker](https://docs.docker.com/get-docker/) (optional, if you want to run the application in a Docker container)

## Installation

1. Clone the repository to your local machine:
    ```bash
    git clone https://github.com/kemul/taxi-fare.git
    cd taxi-fare
    ```

2. Install the Go modules:
    ```bash
    go mod tidy
    ```

## Running the Application

### With Go

1. Prepare an input file (e.g., `input.txt`) with the following format:
    ```
    00:00:00.000 0.0
    00:01:00.123 480.9
    00:02:00.125 1141.2
    00:03:00.100 1800.8
    ```

2. Run the application:
    ```bash
    go run main.go < input.txt
    ```

3. The application will output the calculated fare and the sorted records.

### With Docker

1. Build the Docker image:
    ```bash
    docker build -t taxi-fare-app .
    ```

2. Run the application inside the Docker container:
    ```bash
    docker run --rm -v $(pwd)/input.txt:/root/input.txt taxi-fare-app
    ```

3. The application will output the calculated fare and the sorted records inside the container.

## Testing

To run the unit tests for the project, use the following command:

```bash
go test -v ./...
```

This command will run all the tests in the project and provide a coverage report.

## Project Structure

```
├── Dockerfile           # Dockerfile to containerize the application
├── README.md            # This README file
├── go.mod               # Go module file
├── go.sum               # Go module dependencies
├── main.go              # Main entry point of the application
├── meter                # Package for handling taxi meter logic
│   ├── meter.go
│   └── meter_test.go    # Unit tests for meter package
├── record               # Package for handling record parsing and data structure
│   ├── record.go
│   └── record_test.go   # Unit tests for record package
└── utils                # Utility package for logging and other helper functions
    ├── utils.go
    └── utils_test.go    # Unit tests for utils package
```

```
### Key Points:

- **Running the Application**: Instructions are provided for running the application both directly with Go and within a Docker container.
- **Testing**: Instructions are given for running the unit tests to ensure the code behaves as expected.
- **Project Structure**: The project structure is outlined, showing the organization of the codebase.
- **Docker**: The Docker section provides an easy way to build and run the application in an isolated environment.

This `README.md` provides clear instructions on how to set up, run, and test the application, making it easy for anyone to get started with your project.
