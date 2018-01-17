/*
Author: Filip Johansson

Description:
Delete up from database
*/

package ups

import (
	"net/http"

	"../../database"
	"github.com/labstack/echo"
)

func Delete() echo.HandlerFunc {

	// TODO:
	// Update querys to check if the logged in user is the owner of
	// the post
	const (
		removeQuery string = `
			DELETE FROM Ups
			WHERE Postid = $1`

		updateQuery string = `
			UPDATE Ups
			SET ups = ups - 1
			WHERE Postid = $1`
	)

	return func(c echo.Context) error {

		id := c.Param("id")

		_, err := database.DB.Query(removeQuery, id)
		if err != nil {
			return err
		}

		_, err = database.DB.Query(updateQuery, id)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, H{"message": "Up removed"})
	}

}
