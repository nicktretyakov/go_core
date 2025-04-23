# Go Core

A foundational Go library providing common utilities and abstractions used across multiple Go services, including configuration management, logging, HTTP client wrappers, and error handling patterns.

---

## Table of Contents
1. [Overview](#overview)
2. [Features](#features)
3. [Prerequisites](#prerequisites)
4. [Installation](#installation)
5. [Configuration](#configuration)
6. [Package Structure](#package-structure)
7. [Usage Examples](#usage-examples)
   - [Configuration](#configuration-usage)
   - [Logging](#logging-usage)
   - [HTTP Client](#http-client-usage)
   - [Error Handling](#error-handling-usage)
8. [Testing](#testing)
9. [Contributing](#contributing)
10. [License](#license)

---

## Overview

`go_core` is a shared Go module that consolidates cross-cutting concerns into reusable packages. It aims to reduce boilerplate in Go services by offering:

- Environment-based configuration loading
- Structured, leveled logging
- Robust HTTP client with retry and timeout
- Standardized error types and wrapping
- Utilities for context propagation and metrics

## Features

- 🔧 **Config**: Load structured configuration from JSON, YAML, or environment variables using `viper`.
- 📝 **Logging**: Unified logging interface built on `uber/zap` with JSON output and optional file rotation.
- 🌐 **HTTP Client**: Wrapper around `net/http` supporting retries, backoff, and tracing.
- 🚨 **Error Handling**: Custom error types with stack traces and context using `pkg/errors`.
- 📊 **Metrics**: Helpers to instrument applications with Prometheus metrics.

## Prerequisites

- Go 1.18+ installed
- Modules enabled (`GO111MODULE=on`)

## Installation

```bash
# Initialize go.mod in your project if needed
go mod init github.com/yourorg/yourapp

# Add go_core dependency
go get github.com/nicktretyakov/go_core@latest
```

## Configuration

Configuration structures are defined using Go structs and tagged for `viper`. Example config file (`config.yaml`):

```yaml
env: production
log:
  level: info
  file: /var/log/app.log
httpClient:
  timeout: 5s
  retries: 3
  backoff:
    initial: 100ms
    factor: 2.0
```

Load config in your application:

```go
import (
  "github.com/nicktretyakov/go_core/config"
)

var cfg config.AppConfig
if err := config.Load(&cfg, "config.yaml"); err != nil {
  panic(err)
}
```

## Package Structure

```text
go_core/
├── config/         # Viper-based loader and validation
├── log/            # zap logger initialization and helpers
├── http/           # HTTP client with retries, timeouts, and tracing
├── errors/         # Custom error types and wrappers
├── metrics/        # Prometheus instrumentation helpers
└── ctx/            # Context key definitions and propagation utilities
```

## Usage Examples

### Configuration Usage

```go
import "github.com/nicktretyakov/go_core/config"

type AppConfig struct {
  Env        string `mapstructure:"env"`
  LogLevel   string `mapstructure:"log.level"`
  HTTPClient struct {
    Timeout  time.Duration `mapstructure:"timeout"`
    Retries  int           `mapstructure:"retries"`
  }
}

var cfg AppConfig
config.Load(&cfg, "config.yaml")
```

### Logging Usage

```go
import (
  "github.com/nicktretyakov/go_core/log"
)

logger := log.NewLogger(cfg.LogLevel, cfg.LogFile)
logger.Info("application started")
```

### HTTP Client Usage

```go
import (
  "github.com/nicktretyakov/go_core/http"
  "context"
)

client := http.NewClient(cfg.HTTPClient.Timeout, cfg.HTTPClient.Retries)
resp, err := client.Get(context.Background(), "https://api.example.com/data")
if err != nil {
  logger.Error("failed to fetch data", zap.Error(err))
}
```

### Error Handling Usage

```go
import (
  "github.com/nicktretyakov/go_core/errors"
)

if err := doSomething(); err != nil {
  return errors.Wrap(err, "doSomething failed")
}
```

## Testing

Run unit tests for all packages:

```bash
go test ./config/... ./log/... ./http/... ./errors/... ./metrics/... ./ctx/... -cover
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a feature branch: `git checkout -b feature/foo`.
3. Make your changes with proper tests.
4. Commit and push to your fork.
5. Open a pull request describing your changes.

## License

This module is licensed under the MIT License. See [LICENSE](LICENSE) for details.

