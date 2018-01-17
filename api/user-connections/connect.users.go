/*
Author: Filip Johansson

Description:
Connect a user with another.
*/

package userConnections

import (
	"net/http"

	"../../database"
	"github.com/labstack/echo"
)

type ConnectionRequest struct {
	LoggedInUser  string `json: "loggedInUser"`
	NewConnection string `json: "newConnection"`
}

/*
TODO:
Change this later on so that only the userid of B is posted.
Userid of A should be the logged in users id stored in the session.
Authentication should also be required for connection between users.
*/
func Connect() echo.HandlerFunc {

	const (
		query string = `
			INSERT INTO connectionRequests ( RequestingUser, RespondingUser, Accepted )
			VALUES ( $1, $2, false )
		`
	)

	var connectionRequest ConnectionRequest

	return func(c echo.Context) error {

		c.Bind(&connectionRequest)

		_, err := database.DB.Query(
			query,
			connectionRequest.LoggedInUser,
			connectionRequest.NewConnection)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, H{"message": "User connection requested"})

	}

}
