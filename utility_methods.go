package luxe

import "net"

// IP returns the client's IP address
func (l *LTX) IP() string {
	addr := l.conn.RemoteAddr().String()
	if host, _, err := net.SplitHostPort(addr); err == nil {
		return host
	}
	return addr
}
