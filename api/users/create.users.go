/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: http://github.com/fippli

Description:
Add a new user via post request to /api/users
*/

package users

import (
	"net/http"
	"github.com/labstack/echo"
	"../../database"
)

type NewUser struct {
	Email 		string `json: "username"`
	Password	string `json: "password"`
}

/*
TODO:
*/
func Create() echo.HandlerFunc {

	const (
		query string = `
			INSERT INTO Users ( email, password )
			VALUES ( $1, crypt( $2, gen_salt( 'bf', 8 ) ) )
			RETURNING ID
		`
	)

	var newUser NewUser

	return func( c echo.Context ) error {

		c.Bind( &newUser )

		_, err := database.Connection().Query(
			query,
			newUser.Email,
			newUser.Password )

		if err != nil { return err }

		return c.JSON( http.StatusCreated, H{ "message": "User added" } )
	}



}
