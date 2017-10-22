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

/*
TODO:
Change this later on so that B is changed to the logged in users id.
Only the userid of the connection requester should be posted.
Authentication should also be required for accepting connections.
*/
func Disconnect() echo.HandlerFunc {

	const (
		query string = `
			DELETE FROM connectionRequests
			WHERE A = $1 AND B = $2
			OR A = $2 AND B = $1
		`
	)

	var connectionRequest ConnectionRequest

	return func( c echo.Context ) error {

		c.Bind( &connectionRequest )

		// contactId := c.Param( "id" )
		// NOTE: Uncomment line above and use contactId instead of
		// connectionRequest.NewConnection when sessions are workning.

		_, err := database.Connection().Query(
			query,
			connectionRequest.LoggedInUser,
			connectionRequest.NewConnection )
			//contactId )

		if err != nil { return err }

		return c.JSON( http.StatusCreated, H{ "message": "User connection removed" } )

	}

}
