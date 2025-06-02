# Luxe

A Simple HTTP server framework written in Go.

## Features

- Lightweight HTTP server implementation
- Built-in logging system
- Configurable server settings (timeouts, address, port)
- Connection handling with goroutines

## Quick Start

```go
package main

import "your-module/luxe"

func main() {
    server := luxe.New()
    server.Run()
}
```

## Configuration

The server comes with sensible defaults:

- **Address**: `0.0.0.0` (all interfaces)
- **Port**: `8080`
- **Read Timeout**: `10 seconds`
- **Write Timeout**: `10 seconds`

## Usage

1. Import the package
2. Create a new Luxe instance with `luxe.New()`
3. Start the server with `Run()`

The server will listen for HTTP connections and respond with "Hello, World!" to all requests.

## Example

```go
package main

import "your-module/luxe"

func main() {
    l := luxe.New()
    l.Run() // Server starts on 0.0.0.0:8080
}
```

## License

MIT License
