package main

import (
	"../Logger"
	"../Structs"
	"encoding/json"
	"net"
	"os"
)

type connectorInfo struct {
	Connector []Structs.ConnectorServer
}

var managerClient *net.Conn
var serverList connectorInfo

func setLogger() {
	Logger.SetConsole(true)
	Logger.SetRollingDaily("../logs", "Gate-logs.txt")
}

func setupManagerClient() {
	managerClient, _ := net.Dial("tcp", "127.0.0.1:2000")
	defer managerClient.Close()
}

func getConnectorServer() {
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

func main() {
	setLogger()
	setupManagerClient()
	getConnectorServer()
}

func checkError(err error) {
	if err != nil {
		Logger.Error(err.Error())
	}
}
