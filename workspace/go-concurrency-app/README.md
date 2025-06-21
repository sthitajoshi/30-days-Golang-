# Go Concurrency App

This project demonstrates the use of goroutines and channels in Go to perform concurrent execution of multiple functions. It organizes various tasks into a single main function and utilizes a select statement to manage the execution flow.

## Project Structure

```
go-concurrency-app
├── src
│   ├── main.go          # Entry point of the application
│   └── utils
│       └── helpers.go   # Utility functions for concurrent tasks
├── go.mod               # Module definition and dependencies
└── README.md            # Project documentation
```

## Getting Started

To run the application, follow these steps:

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd go-concurrency-app
   ```

2. **Navigate to the source directory**:
   ```bash
   cd src
   ```

3. **Run the application**:
   ```bash
   go run main.go
   ```

## Functionality

The application concurrently executes multiple tasks defined in `helpers.go` using goroutines. The main function orchestrates these tasks and uses a select statement to handle the results as they become available.

## Dependencies

This project uses Go modules for dependency management. Ensure you have Go installed and set up on your machine. The `go.mod` file contains the necessary module information.

## License

This project is open-source and available under the MIT License.