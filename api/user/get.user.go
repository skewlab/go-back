/*
Author: Filip Johansson

Description:
Get user profile
RETURN: <json>
*/

package user

import (
	"net/http"
	//"../../database"
)

func Get( w http.ResponseWriter, r *http.Request ) {

	// TODO:
	// If it is the logged in user
	// Get all data

	//var privateQuery string = "SELECT * FROM Users WHERE userid = $1"
	//rows, err := database.Connection().Query( privateQuery, userId )

	// TODO:
	// If it is another user
	// Only get the public data

	//var publicQuery string = "SELECT id FROM Users WHERE userid <> $1"
	//rows, err := database.Connection().Query( query )

}

// Get my own user profile
func getMe( userId string ) {}

// Get another user profile
func getFriend( userId string ) {}
