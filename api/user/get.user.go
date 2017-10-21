/*
Author: Filip Johansson

Description:
Get user
RETURN: <json>
*/

package user

import (
	"fmt"
	"net/http"
	"database/sql"
	"../../database"
	"github.com/labstack/echo"
)

type H map[string]interface{}

type User struct {
	Id 					string				 `json:"id"`
	Email 			string 				 `json:"id"`
	Alias 			sql.NullString `json:"id"`
	Birthdate 	sql.NullString `json:"id"`
	Avatar 			sql.NullString `json:"id"`
	Description sql.NullString `json:"id"`
	Website 		sql.NullString `json:"id"`
	Phonenumber sql.NullString `json:"id"`
}

func Get() echo.HandlerFunc {

	return func ( c echo.Context ) error {

		var user User

		userId := c.Param( "id" )

		var query string = `
			SELECT id, email, alias, birthdate, avatar, description, website, phonenumber
			FROM Users
			WHERE id = $1
		`
		rows, err := database.Connection().Query( query, userId )

		for rows.Next() {
			err = rows.Scan( &user.Id, &user.Email, &user.Alias, &user.Birthdate, &user.Avatar, &user.Description, &user.Website, &user.Phonenumber )
			if err != nil { return err }
			fmt.Printf( "\n > %v", user.Id )
		}

		// data := H{ "id":&user.Id, "email":&user.Email, "alias":&user.Alias, "birthdate":&user.Birthdate, "avatar":&user.Avatar, "description":&user.Description, "website":&user.Website, "phonenumber":&user.Phonenumber }

		return c.JSON( http.StatusCreated, user )
	}

}
