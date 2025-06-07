package luxe

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"time"
)

type Luxe struct {
	Server Server
	logger *Logger
	router *Router
}

type Server struct {
	Address            string
	Port               int
	MaxHeaderBytes     int
	MaxRequestBodySize int64
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
}

var defaultServer = Server{
	Address:      "0.0.0.0",
	Port:         8080,
	ReadTimeout:  10 * time.Second,
	WriteTimeout: 10 * time.Second,
}

func New() Luxe {
	return Luxe{
		Server: defaultServer,
		logger: NewLogger(),
		router: NewRouter(),
	}
}

// Run starts the HTTP server
func (l *Luxe) Run() {
	l.logger.Info("Starting server on %s:%d", l.Server.Address, l.Server.Port)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", l.Server.Address, l.Server.Port))
	if err != nil {
		l.logger.Error("Failed to start server: %v", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			l.logger.Warn("Failed to accept connection: %v", err)
			continue
		}
		l.logger.Info("Connection is: %v", conn)

		go l.handleConnection(conn)
	}
}

func (l *Luxe) handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			l.logger.Error("Error in closing the connection: %v", err)
		}
	}(conn)

	conn.SetReadDeadline(time.Now().Add(l.Server.ReadTimeout))
	req, err := readRequest(conn, l.Server.MaxRequestBodySize)
	log.Println("Request is ", req)
	log.Println("*******************************************************")
	if err != nil {
		if errors.Is(err, errReadMaxSize) {
			// Send 413 Payload Too Large response
			response := NewHTTPResponse()
			response.SetStatus(413, "Payload Too Large")
			response.SetHeader("Content-Type", "text/plain")
			response.SetHeader("Connection", "close")
			response.SetTextBody("Request too large")
			conn.Write(response.ToBytes())
			return
		}
		l.logger.Error("Error in reading the request: %v", err)
		return
	}

	// Create context
	ctx := NewLTX(conn, l)

	//Parse the request to proper structure

	err = ctx.ParseRequest(req)
	if err != nil {
		l.logger.Error("Error parsing request: %v", err)
		return
	}

	l.logger.Info("Received %s request to %s", ctx.GetMethod(), ctx.GetPath())

	fmt.Println("CTXXXXXXXXXX is ", ctx.Request.Path)

	router := NewRouter()

	// handle the request
	router.HandleRequest(ctx)

	// Send proper HTTP response
	// response := NewHTTPResponse()
	// response.SetStatus(200, "OK")
	// response.SetHeader("Content-Type", "text/plain")
	// response.SetHeader("Connection", "close")
	// response.SetTextBody("Hello, World!")

	// _, err = conn.Write(response.ToBytes())
	// if err != nil {
	// 	l.logger.Error("Error writing response: %v", err)
	// }
}

// readRequest will read the data from connection
func readRequest(conn net.Conn, maxRequestBodySize int64) (string, error) {
	data := make([]byte, 0)
	buffer := make([]byte, 4096)

	if maxRequestBodySize <= 0 {
		maxRequestBodySize = 1024 * 1024 // 1MB default
	}

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" && len(data) > 0 {
				break
			}
			return "", err
		}
		data = append(data, buffer[:n]...)
		if bytes.Contains(data, []byte("\r\n\r\n")) {
			break
		}
		if int64(len(data)) > maxRequestBodySize {
			return "", errReadMaxSize
		}
	}
	return string(data), nil
}
