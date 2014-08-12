package Structs

type ServerInfo struct {
	Type string
	Name string
	Ip   string
	Port string
}

type ServerInfoList struct {
	Servers []ServerInfo
}
