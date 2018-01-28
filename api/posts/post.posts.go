/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: http://github.com/fippli

Description:
Add article entry to the database.

*/

package posts

import (
	"net/http"
	"time"

	"../../database"
	"github.com/labstack/echo"
)

type NewPost struct {
	Userid       string    `json:"userid"`
	Content      string    `json:"content"`
	Date_created time.Time `json:"date_created"`
	Date_updated time.Time `json:"date_updated"`
}

func Post() echo.HandlerFunc {

	var newPost NewPost
	now := time.Now()

	const (
		query string = `
			INSERT INTO
			Posts ( userid, content, date_created, date_updated )
			VALUES ( $1, $2, $3, $4 )`
	)

	return func(c echo.Context) error {
		c.Bind(&newPost)

		_, err := database.DB.Query(query, newPost.Userid, newPost.Content, now, now)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, H{"message": "Post added"})

	}

}
