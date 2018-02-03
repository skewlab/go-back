/*
Author: Filip Johansson

Description:
Connect on ws://localhost:<port>/websockets

Comment:
Most of the code is taken from:
https://echo.labstack.com/cookbook/websocket
*/

package websock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"../../util"

	"github.com/labstack/echo"
	"github.com/lib/pq"
	"golang.org/x/net/websocket"
)

func waitForNotification(l *pq.Listener) []byte {
	for {
		select {
		case n := <-l.Notify:
			fmt.Println("Received data from channel [", n.Channel, "] :")
			// Prepare notification payload for pretty print
			var prettyJSON bytes.Buffer
			err := json.Indent(&prettyJSON, []byte(n.Extra), "", "\t")
			if err != nil {
				fmt.Println("Error processing JSON: ", err)
				return prettyJSON.Bytes()
			}
			fmt.Println(string(prettyJSON.Bytes()))
			return prettyJSON.Bytes()
		case <-time.After(90 * time.Second):
			fmt.Println("Received no events for 90 seconds, checking connection")
			go func() {
				l.Ping()
			}()

		}
	}
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func Connect(c echo.Context) error {
	var DBconfig DBConfig
	DBconfigFile := util.ReadFile("database/db.config.json")
	err := json.Unmarshal([]byte(DBconfigFile), &DBconfig)

	websocket.Handler(func(ws *websocket.Conn) {

		defer ws.Close()
		reportProblem := func(ev pq.ListenerEventType, err error) {
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		listener := pq.NewListener("dbname="+DBconfig.Dbname+" user="+DBconfig.User+" password="+DBconfig.Password+" sslmode=disable", 10*time.Second, time.Minute, reportProblem)
		err = listener.Listen("events")
		if err != nil {
			panic(err)
		}

		fmt.Println("Start monitoring PostgreSQL...")

		for {
			data := waitForNotification(listener)
			// Write
			fmt.Println(data)
			err := websocket.Message.Send(ws, data)
			if err != nil {
				c.Logger().Error(err)
			}

			// Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Printf("%s\n", msg)

		}

	}).ServeHTTP(c.Response(), c.Request())

	return nil

}
