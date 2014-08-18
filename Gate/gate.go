package main

import (
	"../Logger"
	"../Structs"
	"encoding/json"
	"math/rand"
	"net"
	"os"
	"time"
)

type connectorInfo struct {
	Connector []Structs.ConnectorServer
}

var managerClient *net.Conn
var serverList connectorInfo
var clientHandler *net.Listener

func setLogger() {
	Logger.SetConsole(true)
	Logger.SetRollingDaily("../logs", "Gate-logs.txt")
}

func setupManagerClient() {
	managerClient, _ := net.Dial("tcp", "127.0.0.1:2000")
	defer managerClient.Close()
	for {
		buffer := make([]byte, 1024)
		length, err := managerClient.Read(buffer)
		checkError(err)
		if string(buffer[:length]) == "STOP" {
			managerClient.Close()
			clientHandler.Close()
			Logger.Info("Gate server closed")
			os.Exit(0)
		}
	}
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

	//Logger.Info(serverList)
	Logger.Info("Get connector server config complete")
	return
}

func setupClientHandler() {
	rand.Seed(time.Now().Unix())
	addr, err := net.ResolveTCPAddr("tcp", ":5000")
	checkError(err)

	clientHandler, err := net.ListenTCP("tcp", addr)
	checkError(err)

	for {
		conn, err := clientHandler.Accept()
		checkError(err)

		r := rand.Intn(2)
		conn.Write([]byte(serverList.Connector[r].Ip + ":" + serverList.Connector[r].Port))

	}
}

func main() {
	setLogger()
	getConnectorServer()
	go setupClientHandler()
	setupManagerClient()
}

func checkError(err error) {
	if err != nil {
		Logger.Error(err.Error())
	}
}
