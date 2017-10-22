/*
Author: Filip Johansson

Description:
Get user
RETURN: <json>
*/

package users

import (
	"net/http"
	"database/sql"
	"../../database"
	"github.com/labstack/echo"
)

type H map[string]interface{}

type User struct {
	Id 					string				 `json:"id"`
	Email 			string 				 `json:"email"`
	Alias 			sql.NullString `json:"alias"`
	Birthdate 	sql.NullString `json:"birthdate"`
	Avatar 			sql.NullString `json:"avatar"`
	Description sql.NullString `json:"description"`
	Website 		sql.NullString `json:"website"`
	Phonenumber sql.NullString `json:"phonenumber"`
}

func Get() echo.HandlerFunc {

	const (
		query string = `
			SELECT id, email, alias, birthdate, avatar, description, website, phonenumber
			FROM Users
			WHERE id = $1
		`
	)

	var user User

	return func ( c echo.Context ) error {

		userId := c.Param( "id" )

		rows, err := database.Connection().Query( query, userId )

		for rows.Next() {

			err = rows.Scan(
				&user.Id,
				&user.Email,
				&user.Alias,
				&user.Birthdate,
				&user.Avatar,
				&user.Description,
				&user.Website,
				&user.Phonenumber )

			if err != nil { return err }

		}

		return c.JSON( http.StatusCreated, user )
	}

}
