/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: http://github.com/fippli

Description:
Update article in the database.

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

type UpdateArticle struct {
	Id int `json: "id"`
	Title string `json: "title"`
	Html string `json: "html"`
}

func Update( w http.ResponseWriter, r *http.Request ) {

	var updateArticle UpdateArticle
	err := json.NewDecoder( r.Body ).Decode( &updateArticle )

	if err != nil {
		panic( err )
	}

	now := time.Now()

	var query string = "UPDATE Article SET title = $1, html = $2, date_updated = $3 WHERE id = $4 RETURNING ID"
	id, err := database.Connection().Query( query, updateArticle.Title, updateArticle.Html, now, updateArticle.Id )

	if err != nil {
		panic( err )
	}

	fmt.Printf( "\n > Updated article with id: %v", id )

	io.WriteString( w, "Updated article" )


}
