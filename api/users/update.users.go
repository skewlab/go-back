/*
Author: Filip Johansson
Email: filip.carl.johansson@gmail.com

Description:
Update the fields of the user profile.
*/

package users

import (
	"net/http"
	"time"
	"github.com/labstack/echo"
	"../../database"
)

type UserProfile struct {
	Id 					string
	Email 			string
	Alias 			string
	Avatar 			string
	Birthdate 	time.Time
	Description string
	Website 		string
	Phonenumber string
}

/*
TODO:
Authenticate the user so that only
the user can update its information
*/
func Update() echo.HandlerFunc {

	const (
		insertQuery string = `
			INSERT INTO Users ( email, alias, birthdate, avatar, description, website, phonenumber )
			SELECT $1, $2, $3, $4, $5, $6, $7
			WHERE NOT EXISTS (
				SELECT 1 FROM Users WHERE id = $8
			)`

		updateQuery string = `
			UPDATE Users
			SET email = $1,
					alias = $2,
					birthdate = $3,
					avatar = $4,
					description = $5,
					website = $6,
					phonenumber = $7
			WHERE id = $8;`
	)

	var userProfile UserProfile

	return func( c echo.Context ) error {
		c.Bind( &userProfile )

		_, updateErr := database.Connection().Query(
			updateQuery,
			userProfile.Email,
			userProfile.Alias,
			userProfile.Birthdate,
			userProfile.Avatar,
			userProfile.Description,
			userProfile.Website,
			userProfile.Phonenumber,
			userProfile.Id )

		if updateErr != nil { return updateErr }

		_, insertErr := database.Connection().Query(
			insertQuery,
			userProfile.Email,
			userProfile.Alias,
			userProfile.Birthdate,
			userProfile.Avatar,
			userProfile.Description,
			userProfile.Website,
			userProfile.Phonenumber,
			userProfile.Id )

		if insertErr != nil { return insertErr }

		return c.JSON( http.StatusCreated, H{ "message": "User updated" } )
	}

}
