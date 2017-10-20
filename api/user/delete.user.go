/*
Author: Filip Johansson

Description:
Delete user with delete request via /api/user
*/

package user

import (
	"net/http"
	"../../database"
	"github.com/labstack/echo"
)

func Delete() echo.HandlerFunc {

	return func( c echo.Context ) error {

		userId := c.Param( "id" )

		var query string = `
			DELETE FROM Users WHERE id = $1
		`

		_, err := database.Connection().Query( query, userId )

		if err != nil { return err }

		return c.JSON( http.StatusCreated, H{ "message": "User deleted" } )
	}

}
