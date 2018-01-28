/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: fippli
Description:
Include all the api endpoints.
*/

package api

import (
	"./article"
	"./authenticate"
	"./feed"
	"./posts"
	"./search"
	"./signin"
	"./signout"
	"./ups"
	"./user-connections"
	"./users"
	"./websockets"
	"github.com/labstack/echo"
)

func Module(e *echo.Echo) {

	//
	// Users
	//
	e.GET("/api/users/:id", users.Get())
	e.POST("/api/users", users.Create())
	e.PUT("/api/users", users.Update())
	e.DELETE("/api/users/:id", users.Delete())

	//
	// User connection
	//

	// NOTE: /:id will later be removed
	//			 All connections should be returned
	//			 /:id is now a replacement for logged in user id until sessions works.
	e.GET("/api/user-connections/:id", userConnections.Get())
	e.POST("/api/user-connections", userConnections.Connect())
	e.PUT("/api/user-connections", userConnections.Accept())

	// NOTE: Change this to DELETE after sessions are working,
	//			 change /remove to /:id
	e.POST("/api/user-connections/remove", userConnections.Disconnect())
	e.GET("/api/my-contacts", userConnections.MyContacts())

	//
	// Articles
	//
	e.GET("/api/article/:id", article.Get())
	e.POST("/api/article", article.Post())
	e.PUT("/api/article", article.Put())
	e.DELETE("/api/article/:id", article.Delete())

	//
	// Posts
	//
	e.GET("/api/posts/:id", posts.Get())
	e.POST("/api/posts", posts.Post())
	e.PUT("/api/posts", posts.Put())
	e.DELETE("/api/posts/:id", posts.Delete())

	//
	// Posts
	//
	e.GET("/api/feed/:id", feed.Get())

	//
	// Signin
	//
	e.POST("/api/signin", signin.Post())

	//
	// Sign out
	//
	e.POST("/api/signout", signout.Post())

	//
	// Authenticate
	//
	e.GET("/api/auth", auth.Auth())

	//
	// Websockets
	//
	e.GET("/websocket", websock.Connect)

	//
	// Ups
	//
	e.GET("/api/ups/user/:id", ups.Get())  // Get a users ups
	e.POST("/api/ups", ups.Post())         // Get a users ups
	e.DELETE("/api/ups/:id", ups.Delete()) // Get a posts ups

	//
	// Search
	//
	e.POST("/api/search", search.Post()) // Search for a alias

}
