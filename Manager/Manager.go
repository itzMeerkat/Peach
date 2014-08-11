package main

import (
	"../Logger"
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type serverInfo struct {
	Type string
	Name string
	Ip   string
	Port string
}

type serverInfoList struct {
	Servers []serverInfo
}

var serverList serverInfoList

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

func test1(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		//msg := "Received from " + ws.Request().Host + "  " + reply
		msg := "welcome to websocket do by pp"
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func Client(w http.ResponseWriter, r *http.Request) {
	html := `<!doctype html>
<html>
    
    <script type="text/javascript" src="http://img3.douban.com/js/packed_jquery.min6301986802.js" async="true"></script>
      <script type="text/javascript">
         var sock = null;
         var wsuri = "ws://127.0.0.1:8001";

         window.onload = function() {

            console.log("onload");

            
            try
            {
                sock = new WebSocket(wsuri);
            }catch (e) {
                alert(e.Message);
            }
            
            
            

            sock.onopen = function() {
               console.log("connected to " + wsuri);
            }
            
            sock.onerror = function(e) {
               console.log(" error from connect " + e);
            }
            
            

            sock.onclose = function(e) {
               console.log("connection closed (" + e.code + ")");
            }

            sock.onmessage = function(e) {
               console.log("message received: " + e.data);
               
               $('#log').append('<p> server say: '+e.data+'<p>');
               $('#log').get(0).scrollTop = $('#log').get(0).scrollHeight;
            }
            
         };

         function send() {
            var msg = document.getElementById('message').value;
            $('#log').append('<p style="color:red;">I say: '+msg+'<p>');
                $('#log').get(0).scrollTop = $('#log').get(0).scrollHeight;
                $('#msg').val('');
            sock.send(msg);
         };
      </script>
      <h1>WebSocket chat with server </h1>
          <div id="log" style="height: 300px;overflow-y: scroll;border: 1px solid #CCC;">
          </div>
          <div>
            <form>
                <p>
                    Message: <input id="message" type="text" value="Hello, world!"><button onclick="send();">Send Message</button>
                </p>
            </form>
            
          </div>

</html>`
	io.WriteString(w, html)
}

func setupManager() {
	http.Handle("/", websocket.Handler(test1))
	http.HandleFunc("/chat", Client)
	err := http.ListenAndServe(":824", nil)
	Logger.Info(err.Error())
}

func main() {
	setLogger()
	getServerConfig()
	setupManager()
}

func checkError(err error) {
	if err != nil {
		Logger.Error("Open servers config file failed!" + err.Error())
	}
}
