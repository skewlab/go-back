# go-back
A general purpose backend in GO

## Setup

### Database
** Make sure you have PostgreSQL installed on your computer.**
1. Create database in postrgesql and choose a \<dbname\>, \<username\>, \<password\>.
2. Duplicate example.db.config.json in **database** folder
3. Rename the copy to db.config.json
4. Change the fields to match your setup:
```JSON
{
  "host"     : "localhost",
  "port"     : 5432,
  "user"     : "<username>",
  "password" : "<passeord>",
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
