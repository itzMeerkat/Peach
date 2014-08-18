package main

import (
	"../Logger"
	"../Structs"
	"encoding/json"
	"net"
	//"os"
)

var managerClient *net.Conn
var clientListener *net.TCPListener
var managerCmd *Structs.ServerCommand

func initLogger() {
	Logger.SetConsole(true)
	Logger.SetRollingDaily("../logs", "Connector-logs.txt")
}

func initManager() {
	managerClient, err := net.Dial("tcp", "127.0.0.1:2000")
	checkError(err)
	for {
		buffer := make([]byte, 1024)
		length, err := managerClient.Read(buffer)
		checkError(err)
		err = json.Unmarshal(buffer[:length], managerCmd)
		checkError(err)
		Logger.Info(managerCmd)
	}
}

func setupClientListener() {
	laddr, err := net.ResolveTCPAddr("tcp", ":5001")
	checkError(err)
	clientListener, err = net.ListenTCP("tcp", laddr)
	checkError(err)
	for {
		conn, err := clientListener.Accept()
		checkError(err)

	}
}

func main() {
	initLogger()
	initManager()
	setupClientListener()
}

func checkError(err error) {
	if err != nil {
		Logger.Error(err.Error())
	}
}
