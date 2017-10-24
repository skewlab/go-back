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
	Id						int				`json:"id"`
	Content 			string 		`json:"content"`
	Date_updated 	time.Time `json:"date_updated"`
}

func Put() echo.HandlerFunc {

	now := time.Now()
	var updatePost UpdatePost

	// TODO:
	// Check if the logged in user is the owner of the
	// post that is updated.

	const (
		updateQuery string = `
			UPDATE Posts
			SET content = $1,
					date_updated = $2,
			WHERE id = $3` // TODO: add `AND userid = $4`
	)

	return func( c echo.Context ) error {
		c.Bind( &updatePost )

		// Run update query
		_, updateErr := database.Connection().Query(
			updateQuery,
			updatePost.Content,
			now,
			updatePost.Id ) //, TODO: Add logged in user here)

		if updateErr != nil {
			return updateErr
		}

		return c.JSON( http.StatusCreated, H{ "message": "Post updated" } )
	}



}
