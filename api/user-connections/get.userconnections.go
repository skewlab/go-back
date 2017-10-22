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
	Requester string `json: "requester"`
	Accepter  string `json: "accepter"`
	Accepted  bool	 `json: "accepted"`
}

func Get() echo.HandlerFunc {

	const (
		query string = `
			SELECT *
			FROM UserConnections
			WHERE A = $1 OR B = $1
		`
	)

	var userConnection UserConnection
	var userConnections []UserConnection

	return func ( c echo.Context ) error {

		loggedInUser := c.Param( "id" )

		rows, err := database.Connection().Query( query, loggedInUser )

		for rows.Next() {

			err = rows.Scan(
				&userConnection.Requester,
				&userConnection.Accepter,
				&userConnection.Accepted )

			userConnections = append( userConnections, userConnection )

			if err != nil { return err }

		}

		return c.JSON( http.StatusCreated, userConnections )

	}

}
