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
		fmt.Println(" I SIGN IN innan QUERY!")
		fmt.Println(database.DB)
		fmt.Println("after")
		err := database.DB.QueryRow(query, userCredentials.Email, userCredentials.Password).Scan(&id, &email)
		fmt.Println(" I SIGN IN EFTER QUERY!")
		switch {
		case err == sql.ErrNoRows:
			return c.JSON(http.StatusForbidden, H{"message": "No such user"})

		case err != nil:
			return err

		default:
			fmt.Println(id)
			session.Values["authenticated"] = true
			session.Values["userId"] = id
			session.Save(c.Request(), c.Response())

			var responseString string = "User " + email + " successfully signed in \n"
			return c.JSON(http.StatusOK, H{"message": responseString})
		}
	}

}
