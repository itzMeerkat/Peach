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
var SERVER_NAME string

func setLogger() {
	Logger.SetConsole(true)
	Logger.SetRollingDaily("../logs", "Gate-logs.txt")
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

		r := rand.Intn(len(serverList.Connector[:]))
		conn.Write([]byte(serverList.Connector[r].Ip + ":" + serverList.Connector[r].Port))
	}
}

func main() {
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
