package main

import (
	"fmt"
	// "net/http"
	"encoding/json"
	"github.com/labstack/echo"
	"./api"
	"./util"
)

/*
	Config struct
	Map ./config.json to this struct.
*/
type Config struct {
	Port string
}

func main() {

	var config Config
	configFile := util.ReadFile( "config.json" )
	json.Unmarshal( []byte( configFile ), &config )

	fmt.Printf( "\n > Starting server on localhost%v\n", config.Port )

	// fs := http.FileServer( http.Dir( "static" ) )
	// http.Handle( "/", fs )
	e := echo.New()
	api.Module( e )

	// http.ListenAndServe( config.Port, nil )
	e.Start(":8000")

}
