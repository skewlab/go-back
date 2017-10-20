/*
Author: Filip Johansson
Email: filip.carl.johansson@gmail.com

Description:
Update the tables that the user profile consists of.
*/

package user

import (
	"net/http"
	"fmt"
	"time"
	"encoding/json"
	"reflect"
	"../../database"
)

type UserProfile struct {
	Id string
	Email string
	Alias string
	Avatar string
	Birthdate time.Time
	Description string
}

func Update( w http.ResponseWriter, r *http.Request ) {

	// TODO:
	// Parse json to user profile object
	var userProfile UserProfile
	err := json.NewDecoder( r.Body ).Decode( &userProfile )
	if err != nil { panic( err ) }
	//fmt.Printf( "UP: %v", userProfile )

	// TODO:
	// Authenticate the user so that only
	// the user can update its information

	// TODO:
	// Check if field exists if not add it
	// otherwhise update the field

	up := reflect.ValueOf( userProfile )

	//values := make([]interface{}, v.NumField())
	// NOTE;
	// This is cumbersome. See if there is a better solution.
	// Its now time for Fear the walking dead date: 2017.10.20 time: 01:23
	for i := 0; i < up.NumField(); i++ {
		if up.IsNil() {
			fmt.Println( up.Type().Field( i ).Name )
			fmt.Println( up.Field( i ) )
		}

	}

	updateQuery := "UPDATE Email SET email = $1, public = $2 WHERE userid = $3;"
	_, updateErr := database.Connection().Query( updateQuery, userProfile.Email, false, userProfile.Id )
	if updateErr != nil { panic( updateErr ) }

	insertQuery := "INSERT INTO Email ( email, public, userid ) SELECT $1, $2, $3 WHERE NOT EXISTS (SELECT 1 FROM Email WHERE userid = $3)"
	_, insertErr := database.Connection().Query( insertQuery, userProfile.Email, false, userProfile.Id )
	if insertErr != nil { panic( insertErr ) }

}
