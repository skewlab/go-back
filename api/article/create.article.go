/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: http://github.com/fippli

Description:
Add article entry to the database.

*/

package article

import (
	"fmt"
	"encoding/json"
	"time"
	"io"
	"net/http"
	"../../database"
)

type NewArticle struct {
	Title string `json: "title"`
	Html string `json: "html"`
}

func Create( w http.ResponseWriter, r *http.Request ) {

	var newArticle NewArticle
	err := json.NewDecoder( r.Body ).Decode( &newArticle )

	if err != nil {
		panic( err )
	}

	now := time.Now()

	var query string = "INSERT INTO Article ( title, html, date_created, date_updated ) VALUES ( $1, $2, $3, $4 ) RETURNING ID"
	id, err := database.Connection().Query( query, newArticle.Title, newArticle.Html, now, now )

	if err != nil {
		panic( err )
	}

	fmt.Printf( "\n > New article id: %v", id )

	io.WriteString( w, "Added article" )


}
