/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: https://github.com/fippli

Description:
Get all ups that with specified id.
*/

package ups

import (
	"net/http"
	"../../database"
	"github.com/labstack/echo"
)

type H map[string]interface{}

type Up struct {
	Userid string `json:"userid"`
	Postid int 		`json:"postid"`
}

func Get() echo.HandlerFunc {

	var up Up
	var upArray []Up
	const (
		query string = `
			SELECT *
			FROM Ups
			WHERE userid = $1
		`
	)


	return func( c echo.Context ) error {

		id := c.Param( "id" )

		rows, err := database.Connection().Query( query, id )
		for rows.Next() {
			err = rows.Scan( &up.Userid, &up.Postid )
			if err != nil { return err }
			upArray = append( upArray, up )
		}

		return c.JSON( http.StatusCreated, upArray )
	}

}
