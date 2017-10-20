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
	http.HandleFunc( "/api/article/create", article.Create )	// Create article
	http.HandleFunc( "/api/article/get", article.GetAll )			// Get all articles
	http.HandleFunc( "/api/article/update", article.Update )	// Update article
	http.HandleFunc( "/api/article/remove", article.Remove )	// Remove article

	/* User */
	http.HandleFunc( "/api/user/add", user.Add ) 							// Add user
	// TODO: Update user password
	// NOTE: Trying to update any
	http.HandleFunc( "/api/user/update", user.Update )				// Update user profile
	// TODO: Remove user

	/* Authentication */
	http.HandleFunc( "/api/signin", user.Login )

	/* Page */
	// TODO: Get page content, (articles)
	// NOTE: This should be the content for the specified page.

	/* Files */

}

func test( w http.ResponseWriter, r *http.Request ) {
	fmt.Println( "Test works" )
}
