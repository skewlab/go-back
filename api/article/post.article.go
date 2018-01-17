/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: http://github.com/fippli

Description:
Add article entry to the database.

*/

package article

import (
	"net/http"
	"time"

	"../../database"
	"github.com/labstack/echo"
)

type NewArticle struct {
	Title string `json: "title"`
	Html  string `json: "html"`
}

func Post() echo.HandlerFunc {

	var newArticle NewArticle
	now := time.Now()

	const (
		query string = `
			INSERT INTO
			Article ( title, html, date_created, date_updated )
			VALUES ( $1, $2, $3, $4 )`
	)

	return func(c echo.Context) error {
		c.Bind(&newArticle)

		_, err := database.DB.Query(query, newArticle.Title, newArticle.Html, now, now)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, H{"message": "Article added"})

	}

}
