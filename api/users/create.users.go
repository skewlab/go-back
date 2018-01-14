/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: http://github.com/fippli

Description:
Add a new user via post request to /api/users
*/

package users

import (
	"net/http"

	"../../database"
	"github.com/labstack/echo"
)

type NewUser struct {
	Email    string `json: "username"`
	Password string `json: "password"`
	Alias    string `json:"alias"`
	Avatar   string `json:"avatar"`
	//Birthdate   string `json:"birthdate"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Phonenumber string `json:"phonenumber"`
}

/*
TODO: fix Birthdate formats/functionality
*/
func Create() echo.HandlerFunc {

	const (
		query string = `
			INSERT INTO Users ( email, password, alias, avatar, description, website, phonenumber )
			VALUES ( $1, crypt( $2, gen_salt( 'bf', 8 ) ), $3, $4, $5, $6, $7 )
			RETURNING ID
		`
	)

	var newUser NewUser

	return func(c echo.Context) error {

		c.Bind(&newUser)

		_, err := database.Connection().Query(
			query,
			newUser.Email,
			newUser.Password,
			newUser.Alias,
			newUser.Avatar,
			//newUser.Birthdate,
			newUser.Description,
			newUser.Website,
			newUser.Phonenumber)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, H{"message": "User added"})
	}

}
