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

	"../../database"
	"../posts"
	"github.com/labstack/echo"

	"../../util"
	"github.com/lib/pq"
	"golang.org/x/net/websocket"
)

func waitForNotification(l *pq.Listener) []byte {
	for {
		select {
		case n := <-l.Notify:
			//fmt.Println("Received data from channel [", n.Channel, "] :")
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

type JSONTime time.Time

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
			notification := waitForNotification(listener)

			var raw map[string]interface{}
			json.Unmarshal(notification, &raw)

			if raw["table"] == "posts" {
				var userPost posts.UserPost
				data, _ := json.Marshal(raw["data"])

				json.Unmarshal(data, &raw)
				id, _ := json.Marshal(raw["id"])

				const (
					oneQuery string = `
					SELECT posts.*, coalesce( up_table.up_count, 0 ) ups, users_table.alias, users_table.avatar
					FROM posts
					LEFT JOIN (
						SELECT postid, count(*) up_count
						FROM ups
						GROUP BY postid
					) up_table ON up_table.postid = posts.id
					LEFT JOIN ( 
						SELECT users.id, users.alias, users.avatar 
						FROM users 
					) users_table ON users_table.id = posts.userid 
					WHERE posts.id = $1;`
				)

				rows, err := database.DB.Query(oneQuery, string(id))
				fmt.Println(rows)
				for rows.Next() {
					err = rows.Scan(
						&userPost.Id,
						&userPost.Userid,
						&userPost.Content,
						&userPost.Date_created,
						&userPost.Date_updated,
						&userPost.Ups,
						&userPost.Alias,
						&userPost.Avatar)
					if err != nil {
						fmt.Println("query-error")
						fmt.Println(err)
					}
				}

				jsonUserPost, err := json.Marshal(userPost)
				// Write
				err = websocket.Message.Send(ws, string(jsonUserPost))
				if err != nil {
					fmt.Println("websocket send error")
					c.Logger().Error(err)
				}
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
