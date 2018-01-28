/*
Author: Jonas Johansson

Description:
Get article with specified id
*/

package feed

import (
	"net/http"

	"../../database"
	"../../globalSessions"
	"../posts"
	"github.com/labstack/echo"
)

type H map[string]interface{}

func Get() echo.HandlerFunc {

	var userPost posts.UserPost

	const (
		feedQuery string = `
		SELECT posts.*, coalesce( up_table.up_count, 0 ) ups, users_table.alias, users_table.avatar
		FROM posts
		LEFT JOIN (
			SELECT postid, count(*) up_count
			FROM ups
			GROUP BY postid
		) up_table ON up_table.postid = posts.id
		LEFT JOIN (
			SELECT users.id, users.alias, users.avatar
			FROM users
		) users_table ON users_table.id = posts.userid
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
	)

	return func(c echo.Context) error {

		id := c.Param("id")

		var userPosts []posts.UserPost
		session := globalSessions.GetSession(c)
		if _, ok := session.Values["userId"].(string); ok {
			rows, err := database.DB.Query(feedQuery, id)
			for rows.Next() {
				err = rows.Scan(
					&userPost.Id,
					&userPost.Userid,
					&userPost.Content,
					&userPost.Date_created,
					&userPost.Date_updated,
					&userPost.Ups,
					&userPost.Alias,
					&userPost.Avatar)
				if err != nil {
					return err
				}
				userPosts = append(userPosts, userPost)
			}
			// Return all articles
			return c.JSON(http.StatusCreated, userPosts)
		}
		return c.JSON(http.StatusInternalServerError, "no user id found in session")
	}
}
