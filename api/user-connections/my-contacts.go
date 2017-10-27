/*
Author: Filip Johansson

Description:
Get all contacts of a user.
*/

package userConnections

import (
	"net/http"
	"../../database"
	"github.com/labstack/echo"
)

type UserContact struct {
	RequestingUser	string `json:"requestingUser"`
	RespondingUser	string `json:"respondingUser"`
	Accepted  			bool	 `json:"accepted"`
}

func Get() echo.HandlerFunc {

	const (
		query string = `
			SELECT * FROM (
				(SELECT RespondingUser
				FROM UserConnections
				WHERE RequestingUser = $1
				AND Accepted = true)

				UNION

				(SELECT RequestingUser
				FROM UserConnections
				WHERE RespondingUser = $1
				AND Accepted = true)
			)
		`
	)

	var userContact UserContact
	var userContacts []UserContacts

	return func ( c echo.Context ) error {

		loggedInUser := c.Param( "id" )

		rows, err := database.Connection().Query( query, loggedInUser )

		for rows.Next() {

			err = rows.Scan(
				&userContact.RequestingUser
			)

			userContacts = append( userContacts, userContact )

			if err != nil { return err }

		}

		return c.JSON( http.StatusCreated, userConnections )

	}

}
