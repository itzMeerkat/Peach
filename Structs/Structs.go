package Structs

import (
	"net"
)

type ServerCommand struct {
	Args []string
}

type Server struct {
	Name string
	Ip   string
	Port string
}

type Connection struct {
	uid    int
	socket *net.Conn
}

type ServerList struct {
	Connector []Server
	Chat      []Server
	Gate      []Server
	Logic     []Server
	Manager   []Server
}
