/*
Author: Filip Johansson

Description:
Get article with specified id
*/

package posts

import (
	"net/http"
	"../../database"
	"time"
	"github.com/labstack/echo"
)

type H map[string]interface{}

type UserPost struct {
	Id 						int 			`json:"id"`
	Userid 				string 		`json:"userid"`
	Content 			string 		`json:"content"`
	Date_created 	time.Time `json:"date_created"`
	Date_updated 	time.Time `json:"date_updated"`
}

func Get() echo.HandlerFunc {

	var userPost UserPost
	var query string

	const(
		allQuery string = `SELECT * FROM Posts` // Also count all ups
		countQuery string = ``
		oneQuery string = ``
	)

	return func( c echo.Context ) error {

		id := c.Param( "id" )

		if id  == "all" {
			var userPosts []UserPost
			// TODO: Only posts from the user and its friends should be available.
			rows, err = database.Connection().Query( allQuery )
			for rows.Next() {
				err = rows.Scan( &userPost.Id, &userPost.Userid, &userPost.Content, &userPost.Date_created, &userPost.Date_updated )
				if err != nil { return err }
				userPosts = append( userPosts, userPost )
			}
			// Return all articles
			return c.JSON( http.StatusCreated, userPosts )
		}

		/*
		TODO:
		Update posts to count its ups.
		SELECT COUNT(*) FROM Ups WHERE postid = $1
		Store count in &userPost.Ups
		*/
		query = `SELECT * FROM Posts WHERE id = $1`
		rows, err := database.Connection().Query( query, id )
		for rows.Next() {
			err = rows.Scan( &userPost.Id, &userPost.Userid, &userPost.Content, &userPost.Date_created, &userPost.Date_updated )
			if err != nil { return err }
		}

		return c.JSON( http.StatusCreated, userPost )
	}

}
