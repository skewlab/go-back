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
	"../../globalSessions"
	"github.com/labstack/echo"
)

type H map[string]interface{}

type UserPost struct {
	Id 						int 			`json:"id"`
	Userid 				string 		`json:"userid"`
	Content 			string 		`json:"content"`
	Date_created 	time.Time `json:"date_created"`
	Date_updated 	time.Time `json:"date_updated"`
	Ups 					int 			`json:"ups""`
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
		) up_table ON up_table.postid = posts.id
			WHERE userid=$1 OR userid IN (SELECT id
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
			)) as c ON c.contacts = Users.id)`
			// NOTE: Above query can probably be written more efficient

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
			session := globalSessions.GetSession(c)
			if value, ok := session.Values["userId"].(string); ok {
				loggedInUser := value
	
				rows, err := database.Connection().Query( allQuery, loggedInUser )
				for rows.Next() {
					err = rows.Scan(
						&userPost.Id,
						&userPost.Userid,
						&userPost.Content,
						&userPost.Date_created,
						&userPost.Date_updated,
						&userPost.Ups )
					if err != nil { return err }
					userPosts = append( userPosts, userPost )
				}
				// Return all articles
				return c.JSON( http.StatusCreated, userPosts )
			}
			return c.JSON(http.StatusInternalServerError, "no user id found in session")
		}


		rows, err := database.Connection().Query( oneQuery, id )
		for rows.Next() {
			err = rows.Scan(
				&userPost.Id,
				&userPost.Userid,
				&userPost.Content,
				&userPost.Date_created,
				&userPost.Date_updated,
			 	&userPost.Ups )
			if err != nil { return err }
		}

		return c.JSON( http.StatusCreated, userPost )
	}

}
