package main

import (
	"../Logger"
	"../Structs"
	"encoding/json"
	"net"
	"os"
)

var serverList Structs.ServerList
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

	Logger.Info("Get server config complete")
	return
}

func setLogger() {
	Logger.SetConsole(true)
	Logger.SetRollingDaily("../logs", "logs.txt")
}

func manageServer(conn net.Conn) {
	defer conn.Close()

	_, _ = conn.Write([]byte("LISTEN|:824"))

	//Logger.Debug("LALALA")
	for {
		buffer := make([]byte, 512)
		length, err := conn.Read(buffer)
		checkError(err)
		Logger.Debug(buffer[:length])

	}
}

func setupManager() {
	listenPort, err := net.ResolveTCPAddr("tcp", serverList.Manager[0].Port)
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
	Logger.Info("Starting Manager Server...")
	getServerConfig()
	setupManager()
}

func checkError(err error) {
	if err != nil {
		Logger.Error(err.Error())
	}
}
