/*

Description:
User login endpoint
*/


package user

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
	"database/sql"
	"../../database"
)

type UserCredentials struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

func Login( w http.ResponseWriter, r *http.Request ) {

	var userCredentials UserCredentials
	var username string
	var id string

	err := json.NewDecoder( r.Body ).Decode( &userCredentials )

	if err != nil {
		panic( err )
	}

	var query string = "SELECT id, username FROM users WHERE username = lower( $1 ) AND password = crypt( $2, password )"
	queryErr := database.Connection().QueryRow( query, userCredentials.Username, userCredentials.Password).Scan( &id, &username )

	switch {
		case queryErr == sql.ErrNoRows:
				fmt.Printf("Username or password failed.")
				io.WriteString( w, "Unsuccessfull! Username or password failed \n" )
		case queryErr != nil:
				panic(queryErr)
		default:
				var responseString string = "User " + username + " successfully signed in \n"
				fmt.Printf(responseString)
				io.WriteString( w, responseString )
	}
}
