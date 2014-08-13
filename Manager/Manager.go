package main

import (
	"../Logger"
	"../Structs"
	"encoding/json"
	"net"
	"os"
)

type serverInfoList struct {
	Connector []Structs.ConnectorServer
	Chat      []Structs.ChatServer
	Gate      []Structs.GateServer
	Logic     []Structs.LogicServer
}

var serverList serverInfoList
var serverManager *net.TCPListener

func getServerConfig() {
	serverConfig, err := os.Open("../Config/servers.conf")
	defer serverConfig.Close()
	checkError(err)

	buf := make([]byte, 1024)
	length, err := serverConfig.Read(buf)
	checkError(err)

	err = json.Unmarshal(buf[:length], &serverList)
	checkError(err)

	Logger.Info(serverList)
	Logger.Info("Get server config complete")
	return
}

func setLogger() {
	Logger.SetConsole(true)
	Logger.SetRollingDaily("../logs", "logs.txt")
}

func manageServer(conn net.Conn) {
	defer conn.Close()

	var cmd Structs.ServerCommand

	for {
		buffer := make([]byte, 512)
		len, err := conn.Read(buffer)
		if len > 0 {
			checkError(err)
			json.Unmarshal(buffer[:len], &cmd)
			Logger.Debug(cmd.Args)
		}

	}
}

func setupManager() {
	listenPort, err := net.ResolveTCPAddr("tcp", ":2000")
	checkError(err)

	serverManager, err = net.ListenTCP("tcp", listenPort)
	checkError(err)

	Logger.Info("Server manager setup success")
	for {
		conn, err := serverManager.Accept()
		checkError(err)
		go manageServer(conn)
	}
}

func main() {
	setLogger()
	Logger.Info("Starting Server...")
	getServerConfig()
	setupManager()
}

func checkError(err error) {
	if err != nil {
		Logger.Error(err.Error())
	}
}
