/*
Author: Filip Johansson

Description:
Get all contacts of a user.
*/

package userConnections

import (
	"net/http"
	"database/sql"
	"../../database"
	"github.com/labstack/echo"
)

type Contact struct {
	Id 					string				 `json:"id"`
	Email 			string 				 `json:"email"`
	Alias 			sql.NullString `json:"alias"`
	Birthdate 	sql.NullString `json:"birthdate"`
	Avatar 			sql.NullString `json:"avatar"`
	Description sql.NullString `json:"description"`
	Website 		sql.NullString `json:"website"`
	Phonenumber sql.NullString `json:"phonenumber"`
}

func MyContacts() echo.HandlerFunc {

	const (
		query string = `
			SELECT id, email, alias, birthdate, avatar, description, website, phonenumber
			FROM Users
			JOIN ((
				SELECT RespondingUser contacts
				FROM UserConnections
				WHERE RequestingUser = $1
				AND Accepted = true)

			UNION (
				SELECT RequestingUser contact
				FROM UserConnections
				WHERE RespondingUser = $1
				AND Accepted = true
			)) as c ON c.contacts = Users.id ;
		`
	)

	var contact Contact
	var contacts []Contact

	return func ( c echo.Context ) error {

		loggedInUser := c.Param( "id" )

		rows, err := database.Connection().Query( query, loggedInUser )

		for rows.Next() {

			err = rows.Scan(
				&contact.Id,
				&contact.Email,
				&contact.Alias,
				&contact.Birthdate,
				&contact.Avatar,
				&contact.Description,
				&contact.Website,
				&contact.Phonenumber )

			contacts = append( contacts, contact )

			if err != nil {
				return err
			}

		}

		return c.JSON( http.StatusCreated, contacts )

	}

}
