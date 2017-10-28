/*
Author: Filip Johansson, Jonas Johansson
Description:
User signin endpoint
*/

package signin

import (
	"database/sql"
	"fmt"
	"net/http"

	"../../database"
	"../../globalSessions"
	"github.com/labstack/echo"
)

type H map[string]interface{}

type UserCredentials struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

// func CreateSessionId() string {
// 	b := make([]byte, 32)
// 	if _, err := io.ReadFull(rand.Reader, b); err != nil {
// 		return ""
// 	}
// 	return base64.URLEncoding.EncodeToString(b)
// }

func Post() echo.HandlerFunc {

	var userCredentials UserCredentials
	var email string
	var id string

	const (
		query string = `
			SELECT id, email
			FROM users
			WHERE email = lower( $1 )
			AND password = crypt( $2, password )`
	)

	return func(c echo.Context) error {
		c.Bind(&userCredentials)

		globalSessions.StartSession(c)
		// secretId := CreateSessionId()
		// //fetch new store.... ( ? )
		// store, StoreErr := pgstore.NewPGStore("postgres://postgres:jonas@localhost:5432/go-back?sslmode=disable", []byte(secretId))
		// if StoreErr != nil {
		// 	fmt.Println(StoreErr.Error())
		// }

		// defer store.Close()
		// // Maby change to like every hour in production
		// defer store.StopCleanup(store.Cleanup(time.Minute * 2))
		// //defer store.Cleanup(time.Minute * 2)

		// // Get a session.
		// session, sessionErr := store.Get(c.Request(), "session-key")
		// if sessionErr != nil {
		// 	fmt.Println(sessionErr.Error())
		// }

		err := database.Connection().QueryRow(query, userCredentials.Email, userCredentials.Password).Scan(&id, &email)
		fmt.Println(userCredentials)
		//fmt.Println(session)
		switch {
		case err == sql.ErrNoRows:
			// // Add a value.
			// session.Values["authenticated"] = "false"
			// // Save.
			// if err = session.Save(c.Request(), c.Response()); err != nil {
			// 	fmt.Printf("Error saving session: %v", err)
			// }
			// cookie := new(http.Cookie)
			// cookie.Name = "username"
			// cookie.Value = "jon"
			// cookie.Expires = time.Now().Add(24 * time.Hour)
			// c.SetCookie(cookie)
			return c.JSON(http.StatusCreated, H{"message": "No such user"})

		case err != nil:
			return err

		default:
			// // Add a value.
			// session.Values["authenticated"] = "true"
			// // Save.
			// if err = session.Save(c.Request(), c.Response()); err != nil {
			// 	fmt.Printf("Error saving session: %v", err)
			// }
			// cookie := new(http.Cookie)
			// cookie.Name = "username"
			// cookie.Value = "jon"
			// cookie.Expires = time.Now().Add(24 * time.Hour)
			// c.SetCookie(cookie)
			var responseString string = "User " + email + " successfully signed in \n"
			return c.JSON(http.StatusCreated, responseString)
		}
	}

}
