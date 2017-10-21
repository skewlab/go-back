package main

import (
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

	e := echo.New()
	api.Module( e )

	e.Start( config.Port )

}
