package main

import (
	"../Logger"
	"../Structs"
	"encoding/json"
	"net"
)

var managerClient *net.Conn
var serverList Structs.ServerList
var SERVER_NAME string

func setLogger() {
	Logger.SetConsole(true)
	Logger.SetRollingDaily("../logs", "Channel-logs.txt")
}

func setupManagerClient() {
	managerClient, err := net.Dial("tcp", serverList.Manager[0].Ip+serverList.Manager[0].Port)
	checkError(err)
	defer managerClient.Close()

	managerClient.Write([]byte("ONLINE|GATE_SERVER"))

	for {
		buffer := make([]byte, 512)
		length, err := managerClient.Read(buffer)
		checkError(err)

		cmd := strings.Split(string(buffer[:length]), "|")

		if cmd[0] == "STOP" {
			Logger.Info("Gate server closed")
			os.Exit(0)
		}
		if cmd[0] == "SETUP" {
			Logger.Info("Now my name is " + cmd[1])
			SERVER_NAME = cmd[1]
			Logger.Info("Listening port " + cmd[2])
			go setupClientHandler(cmd[2])
		}
	}
}

func main() {
	setLogger()
}
