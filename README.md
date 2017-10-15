# go-back
A general purpose backend in GO

## Setup

### Database
** Make sure you have PostgreSQL installed on your computer.**
1. Create database in postrgesql and choose a **<dbname>**, **<username>**, **<password>**.
2. Duplicate example.db.config.json in **database** folder
3. Rename the copy to db.config.json
4. Change the fields to match your setup:
  ```JSON
    {
      "host"     : "localhost",
      "port"     : 5432,         // PostgreSQL usually runs on port 5432
      "user"     : "<username>",
      "password" : "<passeord>",
      "dbname"   : "<dbname>"
    }
  ```
 
### Server
1. Copy 
Edit the fields in the
