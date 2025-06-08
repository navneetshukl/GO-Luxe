# Luxe

A Simple HTTP server framework written in Go.

## Features

- Lightweight HTTP server implementation
- Built-in logging system
- Configurable server settings (timeouts, address, port)
- Connection handling with goroutines

## Quick Start

Run `go get github.com/navneetshukl/GO-Luxe@v1.0.0`

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

import (
	luxe "github.com/navneetshukl/GO-Luxe"
)

func main() {
	mux := luxe.New()

	mux.GET("/a", func(l *luxe.LTX) {
		l.SendJSON(200, luxe.H{
			"message": "This is get method 1",
		})
	})

	mux.GET("/b",GetB)

	mux.Run()

}

func GetB(l *luxe.LTX) {
	l.SendJSON(200, luxe.H{
		"message": "This is get method 2",
		"error":   "No Error",
		"age":10,
	})
}

```

## License

MIT License
