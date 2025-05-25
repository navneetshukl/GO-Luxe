package luxe

import (
	"fmt"
	"net"
	"time"
)

type Luxe struct {
	Server Server
	logger *Logger
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
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(l.Server.ReadTimeout))

	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		l.logger.Error("Error reading request: %v", err)
		return
	}

	requestData := string(buffer[:n])
	l.logger.Info("Received request:\n%s", requestData)

	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: 13\r\n" +
		"Connection: close\r\n" +
		"\r\n" +
		"Hello, World!"

	_, err = conn.Write([]byte(response))
	if err != nil {
		l.logger.Error("Error writing response: %v", err)
	}
}
