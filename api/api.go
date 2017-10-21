/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: fippli
Description:
Include all the api endpoints.
*/

package api

import (
	"github.com/labstack/echo"
	"./article"
	"./user"
	"./signin"
)

func Module( e *echo.Echo ) {

	// Static routes for main page and manage page
	e.File("/", "static/index.html")

	// Users
	e.GET("/api/user/:id", user.Get() )
	e.POST("/api/user", user.Post() )
	e.PUT("/api/user", user.Put() )
	e.DELETE("/api/user/:id", user.Delete() )

	// Articles
	e.GET("/api/article/:id", article.Get() )
	e.POST("/api/article", article.Post() )
	e.PUT("/api/article", article.Put() )
	e.DELETE("/api/article/:id", article.Delete() )

	// Signin
	e.POST("/api/signin", signin.Post() )

}
