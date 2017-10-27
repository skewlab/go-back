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
	"fmt"
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

func Connect( c echo.Context ) error {

	websocket.Handler( func( ws *websocket.Conn ) {

		defer ws.Close()

		for {
			// Read
			msg := ""
			err := websocket.Message.Receive( ws, &msg )

			if err != nil {
				c.Logger().Error( err )
			}

			// Write
			err = websocket.Message.Send( ws, "Hello, Client!" )

			if err != nil {
				c.Logger().Error( err )
			}

			fmt.Printf( "%s\n", msg )

		}

	}).ServeHTTP( c.Response(), c.Request() )

	return nil

}
