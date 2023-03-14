package sniproxy

import (
	"net"
)

// Config is the SNI proxy configuration.
type Config struct {
	// TLSListenAddr is the listen address the SNI proxy will be listening to
	// TLS connections.
	TLSListenAddr *net.TCPAddr

	// HTTPListenAddr is the listen address the SNI proxy will be listening to
	// plain HTTP connections.
	HTTPListenAddr *net.TCPAddr

	// ForwardProxy is the address of the SOCKS5 proxy that the connections will
	// be forwarded to according to ForwardRules.
	ForwardProxy string

	// ForwardRules is a list of wildcards that define what connections will be
	// forwarded to the proxy using ProxyDialer.  If the list is empty and
	// ProxyDialer is set, all connections will be forwarded.
	ForwardRules []string

	// BlockRules is a list of wildcards that define connections to which hosts
	// will be blocked.
	BlockRules []string
}
