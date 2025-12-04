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

Output:

```text
2025-12-04:18:00:49,165 INFO     [/home/bugsfounder/workspace/logger/examples/demo/main.go:11] main() This is an info message from the default logger
2025-12-04:18:00:49,165 WARNING  [/home/bugsfounder/workspace/logger/examples/demo/main.go:12] main() This is a warning message
2025-12-04:18:00:49,165 ERROR    [/home/bugsfounder/workspace/logger/examples/demo/main.go:13] main() This is an error message
2025-12-04:18:00:49,165 DEBUG    [/home/bugsfounder/workspace/logger/examples/demo/main.go:17] main() This is a debug message from a custom logger
```

## Features

- Supports Debug, Info, Warning, and Error log levels.
- Automatically captures filename, line number, and function name.
- Thread-safe (wraps standard `log.Logger`).

## License

MIT
