package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"./api"
	"./util"
	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
)

/*
	Config struct
	Map ./config.json to this struct.
*/
type Config struct {
	Port   string
	Static string
}

func main() {

	var config Config
	configFile := util.ReadFile("config.json")
	json.Unmarshal([]byte(configFile), &config)
	e := echo.New()

	// NOTE: Allow CORS for development
	// This should be carefully set in production mode
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"localhost", "http://localhost:3000", "127.0.0.1"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	///////////////

	// Static routes for main page and manage page
	e.Use(middleware.Static(config.Static))

	// Sessions
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.GET("/sess", func(c echo.Context) error {

		//dbConfig := database.GetDbConfig()
		// Fetch new store.
		fmt.Println("hej")
		store, err := pgstore.NewPGStore("postgres://postgres:jonas@localhost:5432/go-back?sslmode=disable", []byte("secret-key"))
		if err != nil {
			fmt.Println(err)
			fmt.Println(err.Error())
		}
		fmt.Println("efter store")

		defer store.Close()
		fmt.Println(store)

		// Run a background goroutine to clean up expired sessions from the database.
		defer store.StopCleanup(store.Cleanup(time.Minute * 5))

		// Get a session.
		session, err := store.Get(c.Request(), "session-key")
		if err != nil {
			fmt.Println(err.Error())
		}

		// Add a value.
		session.Values["foo"] = "bar"

		// Save.
		if err = session.Save(c.Request(), c.Response()); err != nil {
			fmt.Printf("Error saving session: %v", err)
		}

		// Delete session.
		//session.Options.MaxAge = -1
		//if err = session.Save(c.Request(), c.Response()); err != nil {
		//	fmt.Printf("Error saving session: %v", err)
		//}

		// sess, _ := session.Get("session", c)
		// sess.Options = &sessions.Options{
		// 	Path:     "/",
		// 	MaxAge:   86400 * 7,
		// 	HttpOnly: true,
		// }
		// sess.Values["foo"] = "bar"
		// sess.Save(c.Request(), c.Response())
		fmt.Println(session)
		return c.JSON(http.StatusOK, "session set")
	})

	api.Module(e)

	e.Start(config.Port)

}
