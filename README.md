# go-back
A general purpose backend in GO

## Setup

### Database
*Make sure you have PostgreSQL installed on your computer.*
1. Create database in postrgesql and choose a *\<dbname\>*, *\<username\>*, *\<password\>*.
2. Duplicate example.db.config.json in **database** folder
3. Rename the copy to db.config.json
4. Change the fields to match your setup:
```JSON
{
  "host"     : "localhost",
  "port"     : 5432,
  "user"     : "<username>",
  "password" : "<password>",
  "dbname"   : "<dbname>"
}
```
 
### Server
1. Duplicate example.config.json
2. Rename the copy to config.json
3. Select a port to run the server on. Note that the server port should be a string ":nnnn".
4. Change the fields to match your setup:
```JSON
{
  "port" : ":0000"
}
```
### Contact
If you are using the contact module, you will need to configure contact.config.json.
1. Duplicate example.contact.config.json
2. Rename the copy to contact.config.json
3. Change <gmail> to your gmail address (This would preferrably be another than your regular)
4. Change <password> to the password for the given gmail account.
5. Change <forward email> to the email you wish to send the contents of the form to.
```
{
  "SMTP": {
    "host": "smtp.gmail.com",
    "port": ":587"
  },
  "gmail": "<gmail>",
  "password": "<password>",
  "forward": "<forward email>"
}
```
