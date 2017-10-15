/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: http://github.com/fippli

Description:
Remove article from the database.

*/

package article

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
	"../../database"
)

type RemoveArticle struct {
	Id int `json: "id"`
}

func Remove( w http.ResponseWriter, r *http.Request ) {

	var removeArticle RemoveArticle
	err := json.NewDecoder( r.Body ).Decode( &removeArticle )
	if err != nil { panic( err ) }

	var query string = "DELETE FROM Article WHERE id = $1 RETURNING id"
	id, err := database.Connection().Query( query, removeArticle.Id )
	if err != nil { panic( err ) }

	fmt.Printf( "\n > Removed article with id: %v", id )
	io.WriteString( w, "Removed article" )

}
