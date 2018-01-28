/*
Author: Filip Johansson
Email: filip.carl.johansson@gmail.com

Description:
Update the fields of the user profile.
*/

package users

import (
	"fmt"
	"net/http"
	"time"

	"../../database"
	"github.com/labstack/echo"
)

type UserProfile struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	Alias       string `json:"alias"`
	Avatar      string `json:"avatar"`
	Birthdate   string `json:"birthdate"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Phonenumber string `json:"phonenumber"`
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

	return func(c echo.Context) error {

		c.Bind(&userProfile)

		birthdate, err := time.Parse(userProfile.Birthdate, userProfile.Birthdate)

		if err != nil {
			return err
		}

		fmt.Printf("%v", userProfile)

		_, updateErr := database.DB.Query(
			updateQuery,
			userProfile.Email,
			userProfile.Alias,
			birthdate,
			userProfile.Avatar,
			userProfile.Description,
			userProfile.Website,
			userProfile.Phonenumber,
			userProfile.Id)

		if updateErr != nil {
			return updateErr
		}

		_, insertErr := database.DB.Query(
			insertQuery,
			userProfile.Email,
			userProfile.Alias,
			birthdate,
			userProfile.Avatar,
			userProfile.Description,
			userProfile.Website,
			userProfile.Phonenumber,
			userProfile.Id)

		if insertErr != nil {
			return insertErr
		}

		return c.JSON(http.StatusCreated, H{"message": "User updated"})
	}

}
