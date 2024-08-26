# Taxi Fare Calculator

This project is a taxi fare calculator implemented in Go. It processes records of distances traveled and calculates the total fare based on specified rates. The fare is calculated with different rates for different distance ranges.

## Table of Contents
- [Fare Calculation](#fare-calculation)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [Testing](#testing)
- [Docker](#docker)
- [Project Structure](#project-structure)

## Fare Calculation
1. Given the input data
```text
00:00:00.000 0.0
00:01:00.123 480.9
00:02:00.125 1141.2
00:03:00.100 1800.8
```
2. Workflow
- Processing line: `00:00:00.000 0.0`
    ```
    Step 1: Initial fare: 400 yen for up to 1 km.
    ```        
- Processing line: `00:01:00.123 480.9`
    ```
    Step 2: Current Distance: 480.9 meters
    Still within the first 1 km, no additional fare. Fare remains: 400 yen
    ```        
- Processing line: `00:02:00.125 1141.2`
    ```
    Step 3: Current Distance: 1141.2 meters
    Additional distance beyond 1 km: 141.2 meters
    Number of 400m units: 0.35
    Additional fare: 14.12 yen (0.35 * 40.00)
    Total fare after this step: 414 yen
    ```
- Processing line: `00:03:00.100 1800.8`
    ```
    Step 4: Current Distance: 1800.8 meters
    Additional distance beyond 1 km: 659.6 meters
    Number of 400m units: 1.65
    Additional fare: 65.96 yen (1.65 * 40.00)
    Total fare after this step: 480 yen
    ```
- Total Fare: 480 yen




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

Output 
```
PS E:\Workspace\taxi-fare> docker run --rm taxi-fare-app
2024/08/26 12:58:46 =========================================================
2024/08/26 12:58:46 Processing line: {0000-01-01 00:00:00 +0000 UTC 0 0}
2024/08/26 12:58:46 Step 1: Initial fare: 400 yen for up to 1 km.
2024/08/26 12:58:46 =========================================================
2024/08/26 12:58:46 Processing line: {0000-01-01 00:01:00.123 +0000 UTC 480.9 480.9}
2024/08/26 12:58:46 Step 2: Current Distance: 480.9 meters
2024/08/26 12:58:46 Still within the first 1 km, no additional fare. Fare remains: 400 yen
2024/08/26 12:58:46 =========================================================
2024/08/26 12:58:46 Processing line: {0000-01-01 00:02:00.125 +0000 UTC 1141.2 660.3000000000001}
2024/08/26 12:58:46 Step 3: Current Distance: 1141.2 meters
2024/08/26 12:58:46 Additional distance beyond 1 km: 141.2 meters
2024/08/26 12:58:46 Number of 400m units: 0.35
2024/08/26 12:58:46 Additional fare: 14.12 yen (0.35 * 40.00)
2024/08/26 12:58:46 Total fare after this step: 414 yen
2024/08/26 12:58:46 =========================================================
2024/08/26 12:58:46 Processing line: {0000-01-01 00:03:00.1 +0000 UTC 1800.8 659.5999999999999}
2024/08/26 12:58:46 Step 4: Current Distance: 1800.8 meters
2024/08/26 12:58:46 Additional distance beyond 1 km: 659.6 meters
2024/08/26 12:58:46 Number of 400m units: 1.65
2024/08/26 12:58:46 Additional fare: 65.96 yen (1.65 * 40.00)
2024/08/26 12:58:46 Total fare after this step: 480 yen
2024/08/26 12:58:46 Total Fare: 480 yen
```

## Testing

To run the unit tests for the project, use the following command:

```bash
go test ./... -coverprofile=coverage
```
Result Test Coverage 
```bash
PS E:\Workspace\taxi-fare> go test ./... -coverprofile=coverage
ok      taxi-fare       0.351s  coverage: 84.6% of statements
ok      taxi-fare/meter 0.310s  coverage: 79.5% of statements
ok      taxi-fare/record        0.298s  coverage: 100.0% of statements
ok      taxi-fare/utils 0.307s  coverage: 100.0% of statements
```
Preview Code Coverage with this file : https://github.com/kemul/taxi-fare/blob/main/coverage.html
![image](https://github.com/user-attachments/assets/459357f1-bab0-43ec-9fdb-fe4ac84addce)


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

### Key Points:

- **Running the Application**: Instructions are provided for running the application both directly with Go and within a Docker container.
- **Testing**: Instructions are given for running the unit tests to ensure the code behaves as expected.
- **Project Structure**: The project structure is outlined, showing the organization of the codebase.
- **Docker**: The Docker section provides an easy way to build and run the application in an isolated environment.

This `README.md` provides clear instructions on how to set up, run, and test the application, making it easy for anyone to get started with your project.
