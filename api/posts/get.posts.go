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

	const(
		allQuery string = `
			SELECT posts.*, coalesce( up_table.up_count, 0 ) ups
			FROM posts
			LEFT JOIN (
				SELECT postid, count(*) up_count
				FROM ups
				GROUP BY postid
			) up_table ON up_table.postid = posts.id`
			// NOTE: Later add where user ids are any of the logged in users contacts.

		oneQuery string = `
			SELECT posts.*, coalesce( up_table.up_count, 0 ) ups
			FROM posts
			LEFT JOIN (
				SELECT postid, count(*) up_count
				FROM ups
				GROUP BY postid
			) up_table ON up_table.postid = posts.id
			WHERE postid = $1`
	)

	return func( c echo.Context ) error {

		id := c.Param( "id" )

		if id  == "all" {
			var userPosts []UserPost
			// TODO: Only posts from the user and its friends should be available.
			rows, err := database.Connection().Query( allQuery )
			for rows.Next() {
				err = rows.Scan( &userPost.Id, &userPost.Userid, &userPost.Content, &userPost.Date_created, &userPost.Date_updated )
				if err != nil { return err }
				userPosts = append( userPosts, userPost )
			}
			// Return all articles
			return c.JSON( http.StatusCreated, userPosts )
		}


		rows, err := database.Connection().Query( oneQuery, id )
		for rows.Next() {
			err = rows.Scan( &userPost.Id, &userPost.Userid, &userPost.Content, &userPost.Date_created, &userPost.Date_updated )
			if err != nil { return err }
		}

		return c.JSON( http.StatusCreated, userPost )
	}

}
