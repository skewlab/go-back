/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: http://github.com/fippli

Description:
Up something.

*/

package ups

import (
	"github.com/labstack/echo"
	"net/http"
	"../../database"
)

func Post() echo.HandlerFunc {

	// TODO:
	// Change userid from posted object to logged in user.
	// The user should only be able to up once
	const (
		query string = `
			INSERT INTO
			Ups ( userid, postid )
			VALUES ( $1, $2 )`
	)

	var up Up

	return func( c echo.Context ) error {
		c.Bind( &up )

		_, err := database.Connection().Query( query, up.Userid, up.Postid )

		if err != nil { return err }

		return c.JSON( http.StatusCreated, H{ "message":"Up added" } )

	}

}
