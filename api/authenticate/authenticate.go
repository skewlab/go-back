/*
	Author: Jonas Johansson
	Email:  jan.jonas.johansson@gmail.com
	Github: jjojo
	Description:
	Authenticate based on existing cookie store
*/

package auth

import (
	"net/http"

	"../../globalSessions"
	"github.com/labstack/echo"
)

type H map[string]interface{}

/*
	Checks for existing session with authentication
	value set to true, if non found returns Http
	status 402 (forbidden)
*/
func Auth() echo.HandlerFunc {

	return func(c echo.Context) error {

		session := globalSessions.GetSession(c)

		if session.Values["authenticated"] == true {
			return c.JSON(http.StatusOK, H{"message": "Authenticated"})
		}
		return c.JSON(http.StatusForbidden, H{"message": "not authenticated"})
	}

}
