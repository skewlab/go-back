/*
Author: Jonas Johansson
Email:  jan.jonas.johansson@gmail.com
Github: jjojo
Description:
Session handeling and retrieving
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

/*
	Get session, if no session is found it creates one
*/
func GetSession(c echo.Context) *sessions.Session {
	// "session-key" is just name of the cookie value
	session, err := store.Get(c.Request(), "session-key")
	if err != nil {
		fmt.Println(err.Error())
	}
	return session
}

/*
	Generate unique session id for user
*/
func CreateSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
