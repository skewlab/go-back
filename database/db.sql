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

-- --------------------------------------------------------

--
-- Table structure for table users
--

CREATE TABLE Users (
	ID uuid NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
	EMAIL text NOT NULL,
	PASSWORD text NOT NULL,
	ALIAS text,
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
  TITLE varchar(255),
  HTML text,
  DATE_CREATED timestamp,
  DATE_UPDATED timestamp
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
	USERID uuid,
	CONTENT text,
	DATE_CREATED timestamp,
	DATE_UPDATED timestamp
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
	RequestingUser uuid,
	RespondingUser uuid,
	Accepted boolean
);

--
-- Comments:
-- A reaches out to B. If b Accepts the accepted field is set to true.
--  If relation is terminated the row is removed from the database.

-- --------------------------------------------------------

--
-- Table structure for table contacts
--

CREATE TABLE Ups (
	Userid uuid,
	Postid int
);

--
-- Comments:
--

-- --------------------------------------------------------
