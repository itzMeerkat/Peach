package main

import (
	"../Logger"
	"net"
	"time"
)

var managerClient *net.Conn

func setLogger() {
	Logger.SetConsole(true)
	Logger.SetRollingDaily("../logs", "Gate-logs.txt")
}
func main() {
	managerClient, _ := net.Dial("tcp", "127.0.0.1:2000")
	defer managerClient.Close()
	for {
		managerClient.Write([]byte(`{"ArgsAmount":1,"Args":"GATE|Hello!Manager!"}`))
		time.Sleep(time.Microsecond * 10)
	}

}
