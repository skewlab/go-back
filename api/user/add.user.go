/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: http://github.com/fippli

Description:
Add a new user.

*/

package user

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
	"../../database"
)


type NewUser struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

func Add( w http.ResponseWriter, r *http.Request ) {

	var newUser NewUser
	err := json.NewDecoder( r.Body ).Decode( &newUser )

	if err != nil {
		panic( err )
	}

	/* INSERT INTO Users ( name, password ) VALUES ( 'admin', crypt( 'password', gen_salt( 'bf', 8 ) ) ); */
	var query string = "INSERT INTO Users ( username, password ) VALUES ( $1, crypt( $2, gen_salt( 'bf', 8 ) ) ) RETURNING ID"
	id, err := database.Connection().Query( query, newUser.Username, newUser.Password )

	if err != nil {
		panic( err )
	}

	fmt.Printf( "\n > New user id: %v", id )

	io.WriteString( w, "Added user" )


}
