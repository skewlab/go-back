/*
	Author: Jonas Johansson
	Email:  jan.jonas.johansson@gmail.com
	Github: jjojo
	Description:
	User sign out.
*/

package signout

import (
	"net/http"

	"../../globalSessions"
	"github.com/labstack/echo"
)

type H map[string]interface{}

/*
	Sets sessions authentication value to false.
*/
func Post() echo.HandlerFunc {

	return func(c echo.Context) error {

		session := globalSessions.GetSession(c)

		session.Values["authenticated"] = false
		session.Save(c.Request(), c.Response())

		var responseString = "User destroyed his/hers token while signing out \n"
		return c.JSON(http.StatusOK, H{"message": responseString})

	}

}
