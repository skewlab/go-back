/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: http://github.com/fippli

Description:
Add a new user via post request to /api/user

*/

package user

import (
	// "fmt"
	// "encoding/json"
	// "io"
	"net/http"
	"github.com/labstack/echo"
	"../../database"
)


type NewUser struct {
	Email string `json: "username"`
	Password string `json: "password"`
}

func Post() echo.HandlerFunc {

	return func( c echo.Context ) error {
		var newUser NewUser
		c.Bind( &newUser )

		var query string = `
			INSERT INTO Users ( email, password )
			VALUES ( $1, crypt( $2, gen_salt( 'bf', 8 ) ) )
			RETURNING ID
		`
		_, err := database.Connection().Query( query, newUser.Email, newUser.Password )
		if err != nil { return err }

		return c.JSON( http.StatusCreated, H{ "message": "User added" } )
	}



}
