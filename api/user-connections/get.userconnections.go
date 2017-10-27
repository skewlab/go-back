/*
Author: Filip Johansson

Description:
Get user
RETURN: <json>
*/

package userConnections

import (
	"net/http"
	"../../database"
	"github.com/labstack/echo"
)

type UserConnection struct {
	RequestingUser	string `json:"requestingUser"`
	RespondingUser	string `json:"respondingUser"`
	Accepted  			bool	 `json:"accepted"`
}

func Get() echo.HandlerFunc {

	const (
		query string = `
			SELECT *
			FROM UserConnections
			WHERE RequestingUser = $1 OR RespondingUser = $1
		`
	)

	var userConnection UserConnection
	var userConnections []UserConnection

	return func ( c echo.Context ) error {

		loggedInUser := c.Param( "id" )

		rows, err := database.Connection().Query( query, loggedInUser )

		for rows.Next() {

			err = rows.Scan(
				&userConnection.RequestingUser,
				&userConnection.RespondingUser,
				&userConnection.Accepted )

			userConnections = append( userConnections, userConnection )

			if err != nil { return err }

		}

		return c.JSON( http.StatusCreated, userConnections )

	}

}
