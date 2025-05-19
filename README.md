# Logging Package (Go)

## Overview

**Logger** is a Go package for structured, extensible, and technology-agnostic logging and transaction telemetry.  
The goal is to provide a unified interface for logging (including grouping logs into transactions), while supporting multiple pluggable output drivers (CLI, JSON file, text file, etc). This way, your application code never needs to change if you switch the logging backend.

The implementation leverages **SOLID principles**, **Onion Architecture**, and **Domain Driven Design (DDD)** for maintainability and extensibility.

---

## Features

- **Timestamped logs**: Every log entry is timestamped.
- **Log levels**: Debug, Info, Warning, Error.
- **Attributes/tags**: Add arbitrary metadata to any log.
- **Transactions**: Group multiple logs under a transaction ID and attributes.
- **Multiple drivers**: Output to CLI, formatted JSON file, plain text file, or extend with your own driver.
- **Configuration-based**: Change backend drivers via configuration, no code changes required.
- **Extensible**: Add new drivers easily by implementing the Logger interface.
- **Unit tested**: Core logic is covered by unit tests.

---

## Folder structure

```
logger-app/
├── cmd/                  # Entry point
│   └── main.go
├── config/               # Config loader
│   └── config.go
├── internal/
│   ├── domain/           # Core entities 
│   ├── app/              # Use cases, services
│   ├── infra/            # Drivers, driver factory
│   └── tests/            # Unit tests
├── go.mod
├── go.sum
└── README.md
```

---

## Installation

Clone this repo and install dependencies:

```sh
git clone git@github.com:your_user/logger-app.git
cd logger-app
go mod tidy
```

---

## Usage

### 1. Configure the logger

Edit `config.json` to choose the driver and file path (when needed):

```json
{
  "logger": {
    "type": "cli",         // or "json" or "txt"
    "path": ""             // e.g., "logs.json" or "logs.txt" for file drivers
  }
}
```

Available types:
- `"cli"`: Outputs logs to standard output.
- `"json"`: Writes logs in formatted JSON file.
- `"txt"`: Writes logs as plain text, one per line.

### 2. Run the example

```sh
go run ./cmd/main.go
```

Example output (CLI driver):

```
=== Start Transaction: 6f7f2c6e-4c3e-4b8b-8e5c-5beb2e72ebc2 map[customerId:123 origin:http] ===
2025-05-19T19:39:48Z [INFO][TX:6f7f2c6e-4c3e-4b8b-8e5c-5beb2e72ebc2]Processing started | attrs=map[step:1]
2025-05-19T19:39:48Z [WARNING][TX:6f7f2c6e-4c3e-4b8b-8e5c-5beb2e72ebc2]Slow DB response | attrs=map[ms:234 step:2]
2025-05-19T19:39:48Z [ERROR][TX:6f7f2c6e-4c3e-4b8b-8e5c-5beb2e72ebc2]Failed to process payment | attrs=map[error:insufficient funds]
=== End Transaction: 6f7f2c6e-4c3e-4b8b-8e5c-5beb2e72ebc2 ===
```

## How requirements are implemented

| Requirement                  | How it is implemented                                    |
|------------------------------|---------------------------------------------------------|
| Timestamp for every log      | Field `Timestamp` in `LogEntry` set automatically       |
| Several log levels           | Enum `LogLevel` with Debug, Info, Warning, Error        |
| Meta info (tags/attributes)  | `Attributes` map in `LogEntry` and in `Transaction`     |
| Multiple drivers             | Drivers implement `Logger` interface. Factory pattern.  |
| Transaction-styled logs      | Transaction ID/attributes, all logs link to transaction |
| Usability                    | Runs out of the box, CLI driver as default              |
| Configurable                 | Driver chosen via config file                           |
| Extensible                   | Add new driver: implement the `Logger` interface        |
| Language: Go                 | All code in Go, modular and testable                    |
| Unit tested                  | See `internal/tests/service_test.go`          |

---

## Running the tests

```sh
go test ./internal/tests/...
```

---

## Adding a new driver

1. Implement the `Logger` interface (see `internal/domain/logger.go`).
2. Add your driver to `driverfactory.go`.
3. Update your `config.json` to select the new type.

---
