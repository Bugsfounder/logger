# Logger

A simple, customizable logger for Go applications.

## Installation

```bash
go get github.com/bugsfounder/logger
```

## Usage

```go
package main

import (
	"os"
	"github.com/bugsfounder/logger"
)

func main() {
	// Use the singleton logger
	log := logger.Logging()
	log.Info("This is an info message")

	// Or create a new logger instance
	customLogger := logger.AppLogger(os.Stdout)
	customLogger.Debug("This is a debug message")
}
```

## Local Tests

Ran go test -v ./... and all tests passed:

```bash
go test -v ./...
```

## Example Run

Ran `go run examples/demo/main.go`:

```bash
go run examples/demo/main.go
```

## Features

- Supports Debug, Info, Warning, and Error log levels.
- Automatically captures filename, line number, and function name.
- Thread-safe (wraps standard `log.Logger`).

## License

MIT
