/*
Author: Filip Johansson, Jonas Johansson
Description:
User signin endpoint
*/


package signin

import (
	"github.com/labstack/echo"
	"net/http"
	"database/sql"
	"../../database"
)

type H map[string]interface{}

type UserCredentials struct {
	Email string `json: "email"`
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

	return func ( c echo.Context ) error {
		c.Bind( &userCredentials )

		err := database.Connection().QueryRow( query, userCredentials.Email, userCredentials.Password).Scan( &id, &email )

		switch {
		case err == sql.ErrNoRows:
					return c.JSON( http.StatusCreated, H{ "message":"No such user" } )

			case err != nil:
					return err

			default:
					var responseString string = "User " + email + " successfully signed in \n"
					return c.JSON( http.StatusCreated, responseString )
		}
	}

}
