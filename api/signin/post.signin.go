/*
Author: Filip Johansson, Jonas Johansson
Description:
User signin endpoint
*/

package signin

import (
	"database/sql"
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

		session := globalSessions.GetSession(c)

		err := database.Connection().QueryRow(query, userCredentials.Email, userCredentials.Password).Scan(&id, &email)

		switch {
		case err == sql.ErrNoRows:
			return c.JSON(http.StatusForbidden, H{"message": "No such user"})

		case err != nil:
			return err

		default:
			session.Values["authenticated"] = true
			session.Save(c.Request(), c.Response())

			var responseString string = "User " + email + " successfully signed in \n"
			return c.JSON(http.StatusOK, H{"message": responseString})
		}
	}

}
