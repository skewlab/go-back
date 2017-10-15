/*

Description:
Get all articles.
*/

package article

import (
	"fmt"
	//"io"
	"../../database"
	"net/http"
	"time"
	"encoding/json"
)

type Article struct {
	Id int `json: "id"`
	Title string `json: "title"`
	Html string `json: "html"`
	Date_created time.Time `json: "date_created"`
	Date_updated time.Time `json: "date_updated"`
}

func GetAll( w http.ResponseWriter, r *http.Request ) {

	var query string = "SELECT * FROM Article"

	rows, err := database.Connection().Query( query )

	if err != nil {
		panic( err )
	}

	var articles []Article

	if err != nil {
		panic( err )
	}

	for rows.Next() {
		var article Article
		err = rows.Scan( &article.Id, &article.Title, &article.Html, &article.Date_created, &article.Date_updated )
		if err != nil { panic ( err ) } // Check for error
		articles = append( articles, article )
		fmt.Printf( "Article: %v", article )
	}

	fmt.Printf( "articles: %v", articles )
	json.NewEncoder( w ).Encode( articles )
	//io.WriteString( w, articles )

}
