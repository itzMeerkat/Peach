package Structs

type ServerCommand struct {
	ArgsAmount int
	Args       string
}

type ConnectorServer struct {
	Name string
	Ip   string
	Port string
}

type LogicServer struct {
	Name string
	Ip   string
	Port string
}

type ChatServer struct {
	Name string
	Ip   string
	Port string
}

type GateServer struct {
	Name string
	Ip   string
	Port string
}
