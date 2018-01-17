/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: http://github.com/fippli

Description:
Update article in the database.

*/

package article

import (
	"net/http"
	"time"

	"../../database"
	"github.com/labstack/echo"
)

type UpdateArticle struct {
	Id    int    `json: "id"`
	Title string `json: "title"`
	Html  string `json: "html"`
}

func Put() echo.HandlerFunc {

	now := time.Now()
	var updateArticle UpdateArticle

	const (
		updateQuery string = `
			UPDATE Article
			SET title = $1,
					html = $2,
					date_updated = $3
			WHERE id = $4;`

		insertQuery string = `
			INSERT INTO Article ( title, html, date_updated  )
			SELECT $1, $2, $3
			WHERE NOT EXISTS ( SELECT 1 FROM Article WHERE id = $4 )`
	)

	return func(c echo.Context) error {
		c.Bind(&updateArticle)

		// Run update query
		_, updateErr := database.DB.Query(updateQuery, updateArticle.Title, updateArticle.Html, now, updateArticle.Id)

		// Add new row if it doesnt exist
		_, insertErr := database.DB.Query(insertQuery, updateArticle.Title, updateArticle.Html, now, updateArticle.Id)

		if updateErr != nil {
			return updateErr
		}

		if insertErr != nil {
			return insertErr
		}

		return c.JSON(http.StatusCreated, H{"message": "Article updated"})
	}

}
