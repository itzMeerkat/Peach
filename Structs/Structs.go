package Structs

import (
	"net"
)

const (
	GATE_SERVER = iota
	CONNECTOR_SERVER
	CHANNEL_SERVER
	LOGIC_SERVER
)

type ServerCommand struct {
	Args []string
}

type Server struct {
	Name        string
	Ip          string
	Port        string
	IsAvailable bool
	Conn        net.Conn
}

type Connection struct {
	uid    int
	socket *net.Conn
}

type ServerList struct {
	Connector []Server
	Channel   []Server
	Gate      []Server
	Logic     []Server
	Manager   []Server
}
