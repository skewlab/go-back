/*
Author: Jonas Johansson
Description: Endpoint for seraching the database
*/

package search

import (
	"database/sql"
	"fmt"
	"net/http"

	"../../database"
	"../../globalSessions"
	"github.com/labstack/echo"
)

type H map[string]interface{}

type SearchQuery struct {
	Searchstring string `json: "searchstring"`
}

type User struct {
	Id    string         `json:"id"`
	Email string         `json:"email"`
	Alias sql.NullString `json:"alias"`
	//Birthdate 	sql.NullString `json:"birthdate"`
	Avatar      sql.NullString `json:"avatar"`
	Description sql.NullString `json:"description"`
	Website     sql.NullString `json:"website"`
	Phonenumber sql.NullString `json:"phonenumber"`
}

func Post() echo.HandlerFunc {

	var searchQuery SearchQuery
	var user User

	const (
		query string = `
		SELECT id, email, alias, avatar, description, website, phonenumber
		FROM users 
		WHERE alias 
		ILIKE $1`
	)

	return func(c echo.Context) error {
		c.Bind(&searchQuery)
		fmt.Print("i search endpoint!")
		fmt.Print(searchQuery.Searchstring)
		var users []User
		session := globalSessions.GetSession(c)
		if _, ok := session.Values["userId"].(string); ok {
			rows, err := database.Connection().Query(query, "%"+searchQuery.Searchstring+"%")
			for rows.Next() {
				err = rows.Scan(
					&user.Id,
					&user.Email,
					&user.Alias,
					//&user.Birthdate,
					&user.Avatar,
					&user.Description,
					&user.Website,
					&user.Phonenumber)
				if err != nil {
					fmt.Print("i forLOOP ERRORORORORO")
					return err
				}
				users = append(users, user)
				fmt.Print(users)
			}
			// Return all articles
			return c.JSON(http.StatusCreated, users)
		}
		return c.JSON(http.StatusInternalServerError, "no user id found in session")
	}
}
