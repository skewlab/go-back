/*
Author: Filip Johansson

Description:
Get article with specified id
*/

package article

import (
	"net/http"
	"time"

	"../../database"
	"github.com/labstack/echo"
)

type H map[string]interface{}

type Article struct {
	Id           int       `json: "id"`
	Title        string    `json: "title"`
	Html         string    `json: "html"`
	Date_created time.Time `json: "date_created"`
	Date_updated time.Time `json: "date_updated"`
}

func Get() echo.HandlerFunc {

	var article Article
	var query string

	return func(c echo.Context) error {

		id := c.Param("id")

		if id == "all" {
			var articles []Article
			query = `SELECT * FROM Article`
			rows, err := database.DB.Query(query)
			for rows.Next() {
				err = rows.Scan(&article.Id, &article.Title, &article.Html, &article.Date_created, &article.Date_updated)
				if err != nil {
					return err
				}
				articles = append(articles, article)
			}
			// Return all articles
			return c.JSON(http.StatusCreated, articles)
		}

		query = `SELECT * FROM Article WHERE id = $1`
		rows, err := database.DB.Query(query, id)
		for rows.Next() {
			err = rows.Scan(&article.Id, &article.Title, &article.Html, &article.Date_created, &article.Date_updated)
			if err != nil {
				return err
			}
		}

		return c.JSON(http.StatusCreated, article)
	}

}
