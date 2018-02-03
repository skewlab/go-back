package main

import (
	"encoding/json"
	"fmt"

	"./database"

	"./api"
	"./util"
	"github.com/coussej/pgbroadcast"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/*
  Config struct
  Map ./config.json to this struct.
*/
type Config struct {
	Port            string
	Static          string
	FrontEndDevPort string
}

// func waitForNotification(l *pq.Listener) {
// 	for {
// 		select {
// 		case n := <-l.Notify:
// 			fmt.Println("Received data from channel [", n.Channel, "] :")
// 			// Prepare notification payload for pretty print
// 			var prettyJSON bytes.Buffer
// 			err := json.Indent(&prettyJSON, []byte(n.Extra), "", "\t")
// 			if err != nil {
// 				fmt.Println("Error processing JSON: ", err)
// 				return
// 			}
// 			fmt.Println(string(prettyJSON.Bytes()))
// 			return
// 		case <-time.After(90 * time.Second):
// 			fmt.Println("Received no events for 90 seconds, checking connection")
// 			go func() {
// 				l.Ping()
// 			}()
// 			return
// 		}
// 	}
// }

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func main() {

	var DBconfig DBConfig
	DBconfigFile := util.ReadFile("database/db.config.json")
	err := json.Unmarshal([]byte(DBconfigFile), &DBconfig)

	var config Config
	configFile := util.ReadFile("config.json")
	json.Unmarshal([]byte(configFile), &config)
	e := echo.New()

	// Init Database connection
	database.DB = database.Connection()
	fmt.Print("db done")
	fmt.Print(database.DB)

	// Create a new broadcaster
	pb, err := pgbroadcaster.NewPgBroadcaster("dbname=" + DBconfig.Dbname + " user=" + DBconfig.User + " password=" + DBconfig.Password + " sslmode=disable")

	// listen to the events channel
	err = pb.Listen("events")
	if err != nil {
		fmt.Println(err)
	}
	// reportProblem := func(ev pq.ListenerEventType, err error) {
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// }

	// listener := pq.NewListener("dbname="+DBconfig.Dbname+" user="+DBconfig.User+" password="+DBconfig.Password+" sslmode=disable", 10*time.Second, time.Minute, reportProblem)
	// err = listener.Listen("feed")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Start monitoring PostgreSQL...")
	// for {
	// 	waitForNotification(listener)
	// }

	// NOTE: Allow CORS for development
	// This should be carefully set in production mode
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"localhost", "http://localhost" + config.FrontEndDevPort},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))
	///////////////

	// Static routes for main page and manage page
	e.Use(middleware.Static(config.Static))

	api.Module(e)
	e.Start(config.Port)
}
