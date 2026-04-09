# jwt-auth

[![Build Status](https://img.shields.io/github/actions/workflow/status/user/jwt-auth/ci.yml?branch=main)]()
[![Go Version](https://img.shields.io/badge/go-1.22+-blue.svg)]()
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

**jwt-auth** zero-dependency JWT authentication library with RS256 and EdDSA support. Built with simplicity and performance in mind.

## Features

- Structured Logging: Built-in structured logging with slog compatibility
- Graceful Shutdown: Clean shutdown handling with configurable drain timeout
- Minimal Allocations: Designed to minimize GC pressure in hot paths
- High Performance: Optimized for low-latency, high-throughput workloads

## Getting Started

### Installation

```bash
go get github.com/user/jwt-auth@latest
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/user/jwt-auth"
)

func main() {
	client := jwtauth.New(
		jwtauth.WithTimeout(30 * time.Second),
	)

	result, err := client.Run(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
```

## Configuration

Configuration can be provided via environment variables, a config file, or programmatically.

| Variable | Description | Default |
|----------|-------------|--------|
| `JWT_AUTH_TIMEOUT` | Request timeout in seconds | `30` |
| `JWT_AUTH_RETRIES` | Number of retry attempts | `3` |
| `JWT_AUTH_LOG_LEVEL` | Log verbosity (debug, info, warn, error) | `info` |

## Contributing

Contributions are welcome! Please read the [contributing guidelines](CONTRIBUTING.md) before submitting a pull request.

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.
