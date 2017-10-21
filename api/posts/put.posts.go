/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: http://github.com/fippli

Description:
Update article in the database.

*/

package posts

import (
	"time"
	"net/http"
	"../../database"
	"github.com/labstack/echo"
)

type UpdatePost struct {
	Id						int				`json: "id"`
	Userid 				string 		`json: "userid"`
	Content 			string 		`json: "content"`
	Date_updated 	time.Time `json: "date_updated"`
	Ups 					int				`json: "ups"`
}

func Put() echo.HandlerFunc {

	now := time.Now()
	var updatePost UpdatePost

	const (
		updateQuery string = `
			UPDATE Posts
			SET userid = $1,
					content = $2,
					date_updated = $3,
					ups = $4
			WHERE id = $5;`
	)

	return func( c echo.Context ) error {
		c.Bind( &updatePost )

		// Run update query
		_, updateErr := database.Connection().Query( updateQuery, updatePost.Userid, updatePost.Content,  now, updatePost.Ups, updatePost.Id )

		if updateErr != nil {
			return updateErr
		}

		return c.JSON( http.StatusCreated, H{ "message": "Post updated" } )
	}



}
