package main

import (
	"encoding/json"

	"./api"
	"./util"
	"github.com/labstack/echo"
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
		AllowOrigins: []string{"localhost", "http://localhost:3001", "127.0.0.1"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	///////////////

	// Static routes for main page and manage page
	e.Use(middleware.Static(config.Static))

	api.Module(e)

	e.Start(config.Port)

}
