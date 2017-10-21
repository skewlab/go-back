/*
Author: Filip Johansson

Description:
Delete user with delete request via /api/user
*/

package posts

import (
	"net/http"
	"../../database"
	"github.com/labstack/echo"
)

func Delete() echo.HandlerFunc {

	const (
		query string = `DELETE FROM Posts WHERE id = $1`
	)

	return func( c echo.Context ) error {

		id := c.Param( "id" )
		_, err := database.Connection().Query( query, id )

		if err != nil { return err }

		return c.JSON( http.StatusCreated, H{ "message": "Post deleted" } )
	}

}
