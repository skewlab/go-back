/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: fippli
Description:
Include all the api endpoints.
*/

package api

import (
	// "fmt"
	// "net/http"
	//"github.com/gorilla/mux"
	"github.com/labstack/echo"
	// "./article"
	"./user"
	// "./signin"
)

func Module( e *echo.Echo ) {

	// Create a new instance of Echo
	// e := echo.New()
	//fmt.Printf( e )

	// Create all routes

	// Static routes for main page and manage page
	e.File("/", "static/index.html")

	// Users
	e.GET("/api/user/:id", user.Get() )
	e.POST("/api/user", user.Post() )
	e.PUT("/api/user", user.Put() )
	e.DELETE("/api/user/:id", user.Delete() )

	// Start as a web server


}
