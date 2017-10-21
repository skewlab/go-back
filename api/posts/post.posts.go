/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: http://github.com/fippli

Description:
Add article entry to the database.

*/

package posts

import (
	"github.com/labstack/echo"
	"time"
	"net/http"
	"../../database"
)

type NewPost struct {
	Userid 				string 		`json: "userid"`
	Content 			string 		`json: "content"`
	Date_created 	time.Time `json: "date_created"`
	Date_updated 	time.Time `json: "date_updated"`
	Ups 					int				`json:"ups"`
}

func Post() echo.HandlerFunc {

	var newPost NewPost
	now := time.Now()

	const (
		query string = `
			INSERT INTO
			Posts ( userid, content, date_created, date_updated, ups )
			VALUES ( $1, $2, $3, $4, $5 )`
	)

	return func( c echo.Context ) error {
		c.Bind( &newPost )

		_, err := database.Connection().Query( query, newPost.Userid, newPost.Content, now, now, 0 )

		if err != nil {
			return err
		}

		return c.JSON( http.StatusCreated, H{ "message":"Post added" } )

	}

}
