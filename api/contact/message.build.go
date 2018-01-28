/*

 */

package contact

import (
	"fmt"
)

func BuildMessage(email Email) string {
	message := ""
	message += fmt.Sprintf("From: %s <%s>\r\n", email.Name, email.From)
	message += fmt.Sprintf("Subject: %s\r\n", email.Subject)
	message += "\r\n" + email.Body
	message += fmt.Sprintf("\r\n\n// %s, %s\r\n", email.Name, email.From)
	return message
}
