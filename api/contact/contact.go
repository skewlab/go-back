/*
Author: Filip Johansson

Description:
Send a mail from a contact form with fields: from, name, subject, body
Edit contact.config.json to add your credentials.

*/

package contact

import (
	"log"
	"net/smtp"
	"encoding/json"
	"net/http"
	"../../util"
)

type ContactConfig struct {
	SMTP struct {
		Host string
		Port string
	}
	Gmail string
	Password string
	Forward string
}

type Email struct {
	From string
	Subject string
	Body string
	Name string
}

func Send( w http.ResponseWriter, r *http.Request ) {

	// Read config file
	var config ContactConfig
	configFile := util.ReadFile( "api/contact/contact.config.json" )
	err := json.Unmarshal( []byte( configFile ), &config )
	if err != nil { panic( err ) }

	// Parse request body
	var email Email
	err = json.NewDecoder( r.Body ).Decode( &email )
	if err != nil { panic( err ) }

	messageBody := BuildMessage( email )

	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		config.Gmail,
		config.Password,
		config.SMTP.Host,
	)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err = smtp.SendMail(
		config.SMTP.Host + config.SMTP.Port,
		auth,
		email.From,
		[]string{ config.Forward },
		[]byte( messageBody ),
	)
	if err != nil { log.Fatal( err ) }
}
