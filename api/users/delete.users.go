/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com

Description:
Delete user with delete request via /api/users
*/

package users

import (
	"net/http"
	"../../database"
	"github.com/labstack/echo"
)

/*
TODO:
Authenticate user. Anyone should not be able to deleta a user.
*/
func Delete() echo.HandlerFunc {

	const (
		query string = `DELETE FROM Users WHERE id = $1`
	)

	return func( c echo.Context ) error {

		id := c.Param( "id" )

		_, err := database.Connection().Query( query, id )

		if err != nil { return err }

		return c.JSON( http.StatusCreated, H{ "message": "User deleted" } )
		
	}

}
