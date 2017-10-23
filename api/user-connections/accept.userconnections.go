/*
Author: Filip Johansson

Description:
Accept a user with connection.
*/

package userConnections

import (
	"net/http"
	"github.com/labstack/echo"
	"../../database"
)

type H map[string]interface{}

/*
TODO:
Change this later on so that B is changed to the logged in users id.
Only the userid of the connection requester should be posted.
Authentication should also be required for accepting connections.
*/
func Accept() echo.HandlerFunc {

	const (
		query string = `
			UPDATE connectionRequests
			SET
				Accepted = true
			WHERE RequestingUser = $1
			AND   RespondingUser = $2
		`
	)

	var connectionRequest ConnectionRequest

	return func( c echo.Context ) error {

		c.Bind( &connectionRequest )

		_, err := database.Connection().Query(
			query,
			connectionRequest.NewConnection,
			connectionRequest.LoggedInUser )

		if err != nil { return err }

		return c.JSON( http.StatusCreated, H{ "message": "User connection requested" } )

	}

}
