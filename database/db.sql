-- Author: Skewlab
-- Github: https://github.com/skewlab/go-back
--
-- Description:
--

-- Remove old data: (This should only be used in development mode)
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO public;

CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_trgm;

-- --------------------------------------------------------

--
-- Table structure for table users
--

CREATE TABLE Users (
	ID uuid NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
	EMAIL text NOT NULL,
	PASSWORD text NOT NULL,
	ALIAS text NOT NULL,
	BIRTHDATE timestamp,
	AVATAR text,
	DESCRIPTION text,
	WEBSITE text,
	PHONENUMBER varchar(255)
);

--
-- Comments:
-- Insert query:
-- 	INSERT INTO
-- 	Users ( name, password )
-- 	VALUES ( 'admin', crypt( 'password', gen_salt( 'bf', 8 ) ) );
--

-- --------------------------------------------------------

--
-- Table structure for table article
--

CREATE TABLE Article (
  ID SERIAL,
  TITLE varchar(255) NOT NULL,
  HTML text NOT NULL,
  DATE_CREATED timestamp NOT NULL,
  DATE_UPDATED timestamp NOT NULL
);

--
-- Comments:
-- Used for page content
-- Different from post wich has a user
--

-- --------------------------------------------------------

--
-- Table structure for table posts
--

CREATE TABLE Posts (
	ID SERIAL,
	USERID uuid NOT NULL,
	CONTENT text NOT NULL,
	DATE_CREATED timestamp NOT NULL,
	DATE_UPDATED timestamp NOT NULL
);

--
-- Comments:
--
--

-- --------------------------------------------------------

--
-- Table structure for table contacts
--

CREATE TABLE UserConnections (
	RequestingUser uuid NOT NULL,
	RespondingUser uuid NOT NULL,
	Accepted boolean NOT NULL
);

CREATE TABLE connectionrequests(
	requestinguser uuid,
	respondinguser uuid,
	accepted boolean
);

--
-- Comments:
-- A reaches out to B. If b Accepts the accepted field is set to true.
--  If relation is terminated the row is removed from the database.

-- --------------------------------------------------------

--
-- Table structure for table ups
--

CREATE TABLE Ups (
	Userid uuid NOT NULL,
	Postid int NOT NULL,
	UNIQUE ( Userid, Postid )
);

--
-- Comments:
-- Add textsearch for user Alias, email and phone number

-- --------------------------------------------------------
create index alias_gin on users using gin(alias gin_trgm_ops);

--
-- Comments:
-- add a trigger function that notifies if any posts change

-- --------------------------------------------------------

CREATE OR REPLACE FUNCTION notify_event() RETURNS TRIGGER AS $$

    DECLARE 
        data json;
        notification json;
    
    BEGIN
    
        -- Convert the old or new row to JSON, based on the kind of action.
        -- Action = DELETE?             -> OLD row
        -- Action = INSERT or UPDATE?   -> NEW row
        IF (TG_OP = 'DELETE') THEN
            data = row_to_json(OLD);
        ELSE
            data = row_to_json(NEW);
        END IF;
        
        -- Contruct the notification as a JSON string.
        notification = json_build_object(
                          'table',TG_TABLE_NAME,
                          'action', TG_OP,
                          'data', data);
        
                        
        -- Execute pg_notify(channel, notification)
        PERFORM pg_notify('events',notification::text);
        
        -- Result is ignored since this is an AFTER trigger
        RETURN NULL; 
    END;
    
$$ LANGUAGE plpgsql;

CREATE TRIGGER posts_notify_event
AFTER INSERT OR UPDATE OR DELETE ON posts
    FOR EACH ROW EXECUTE PROCEDURE notify_event();