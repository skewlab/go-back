/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: fippli
Description:
Include all the api endpoints.
*/

package api

import (
	"fmt"
	"net/http"
	"./article"
	"./user"
)

func Module() {

	/* Test */
	http.HandleFunc( "/api/test", test )

	/* Article */
	http.HandleFunc( "/api/article/create", article.Create )	// Create
	http.HandleFunc( "/api/article/get", article.GetAll )			// Get all articles
	// TODO: Update article
	http.HandleFunc( "/api/article/update", article.Update )			// Get all articles
	// TODO: Remove article

	/* User */
	http.HandleFunc( "/api/user/add", user.Add ) 							// Add user
	// TODO: Update user password
	// TODO: Remove user

	/* Page */
	// TODO: Get page content, (articles)

	/* Files */

}

func test( w http.ResponseWriter, r *http.Request ) {
	fmt.Println( "Test works" )
}
