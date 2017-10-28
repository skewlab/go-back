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
	"net/http"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

type H map[string]interface{}

func Store(secret string) *pgstore.PGStore {
	//fetch new store.... ( ? )
	store, StoreErr := pgstore.NewPGStore("postgres://postgres:jonas@localhost:5432/go-back?sslmode=disable", []byte(secret))
	if StoreErr != nil {
		fmt.Println(StoreErr.Error())
	}

	// MEMORY LEAK, NEED TO CLOSE AFTER SAVE
	// defer store.Close()
	// defer store.StopCleanup(store.Cleanup(time.Minute * 2))

	return store
}

func GetSession(c echo.Context) *sessions.Session {
	store := Store(CreateSessionId())
	session, sessionErr := store.Get(c.Request(), "session-key")
	if sessionErr != nil {
		fmt.Println(sessionErr.Error())
	}
	return session
}

func SaveSession(c echo.Context, sess *sessions.Session) {
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		fmt.Printf("Error saving session: %v", err)
	}
}

func CreateSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func StartSession(c echo.Context) *http.Cookie {
	cookie, err := c.Request().Cookie("session-key")
	if err != nil || cookie.Value == "" {
		sid := CreateSessionId()
		Store(sid)
		sess := GetSession(c)
		sess.Values["authenticated"] = "false"
		SaveSession(c, sess)
		//defer store.Close()

	} else {
		sess := GetSession(c)
		fmt.Println(sess.Values["authenticated"])
		fmt.Println(cookie)
		fmt.Println(cookie.Value)
		//sid := cookie.Value
		//session := GetSession(c)

	}
	return cookie
}
