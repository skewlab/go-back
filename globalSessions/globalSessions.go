/*
Author: Jonas Johansson
Description:
Session handeling
*/

package globalSessions

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

type H map[string]interface{}

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte(CreateSessionId())
	store = sessions.NewCookieStore(key)
)

func GetSession(c echo.Context) *sessions.Session {
	session, err := store.Get(c.Request(), "cookie-name")
	if err != nil {
		fmt.Println(err.Error())
	}
	return session
}

// func SaveSession(c echo.Context, sess *sessions.Session) {
// 	if err := sess.Save(c.Request(), c.Response()); err != nil {
// 		fmt.Printf("Error saving session: %v", err)
// 	}
// }

func CreateSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// func StartSession(c echo.Context) *http.Cookie {
// 	cookie, err := c.Request().Cookie("session-key")
// 	if err != nil || cookie.Value == "" {
// 		sid := CreateSessionId()
// 		Store(sid)
// 		sess := GetSession(c)
// 		sess.Values["authenticated"] = "false"
// 		SaveSession(c, sess)
// 		//defer store.Close()

// 	} else {
// 		sess := GetSession(c)
// 		fmt.Println(sess.Values["authenticated"])
// 		fmt.Println(cookie)
// 		fmt.Println(cookie.Value)
// 		//sid := cookie.Value
// 		//session := GetSession(c)

// 	}
// 	return cookie
// }
