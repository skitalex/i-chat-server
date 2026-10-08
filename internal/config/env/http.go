package env

import (
	"chatsrv/internal/config"
	"net"
)

type httpCfg struct {
	host string
	port string
}

func NewHttpConfig() *httpCfg {
	host := config.GetEnvStringOrDefault("HTTP_HOST_CHAT", "0.0.0.0")
	port := config.GetEnvStringOrDefault("HTTP_PORT_CHAT", "8181")

	return &httpCfg{
		host: host,
		port: port,
	}
}

func (c *httpCfg) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
