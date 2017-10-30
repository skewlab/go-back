/*
Author: Jonas Johansson
Description:
User sessions endpoint
*/


package session

import (
	"github.com/labstack/echo"	
	"net/http"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"fmt"
)

// ExampleHandler is an example that displays the usage of PGStore.
func SetSession( c echo.Context ) error {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
	  Path:     "/session",
	  MaxAge:   86400 * 7,
	  HttpOnly: true,
	}
	sess.Values["foo"] = "bar"
	sess.Save(c.Request(), c.Response())
	fmt.Println( sess )
	return c.NoContent(http.StatusOK)
}


/*
func GetSession( ) {
	// Get a session.
	session, err := store.Get(r, "session-key")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func AddValueToSession( ) {
	session.Values["foo"] = "bar"
}
*/
