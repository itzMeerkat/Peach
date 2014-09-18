package main

import (
	"../Logger"
	"../Structs"
	"encoding/json"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

var managerClient *net.Conn
var serverList Structs.ServerList
var clientHandler *net.Listener
var SERVER_TYPE string

func setLogger() {
	Logger.SetConsole(true)
	Logger.SetRollingDaily("../logs", "Gate-logs.txt")
}

func setupManagerClient() {
	managerClient, err := net.Dial("tcp", serverList.Manager[0].Ip+serverList.Manager[0].Port)
	checkError(err)
	defer managerClient.Close()

	//var cmd []string

	for {
		buffer := make([]byte, 1024)
		length, err := managerClient.Read(buffer)
		checkError(err)

		cmd := strings.Split(string(buffer[:length]), "|")

		if cmd[0] == "STOP" {
			Logger.Info("Gate server closed")
			os.Exit(0)
		}
		if cmd[0] == "LISTEN" {
			Logger.Info("Listening port " + cmd[1])
			go setupClientHandler(cmd[1])
		}
	}
}

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
}

func setupClientHandler(p string) {
	rand.Seed(time.Now().Unix())
	addr, err := net.ResolveTCPAddr("tcp", p)
	checkError(err)

	clientHandler, err := net.ListenTCP("tcp", addr)
	checkError(err)

	for {
		conn, err := clientHandler.Accept()
		checkError(err)
		Logger.Info("A client from " + conn.RemoteAddr().String() + " is online")

		r := rand.Intn(2)
		conn.Write([]byte(serverList.Connector[r].Ip + ":" + serverList.Connector[r].Port))
	}
}

func main() {
	SERVER_TYPE = "GATE"
	setLogger()
	Logger.Info("Starting Gate Server...")
	getServerConfig()
	setupManagerClient()
}

func checkError(err error) {
	if err != nil {
		Logger.Error(err.Error())
	}
}
