-- Author: Skewlab
-- Github: https://github.com/skewlab/go-back
--
-- Description:
--


CREATE EXTENSION pgcrypto;

-- --------------------------------------------------------

--
-- Table structure for table users
--

DROP TABLE Users;

CREATE TABLE Users (
  ID uuid NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
  EMAIL text NOT NULL,
  PASSWORD text NOT NULL
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
-- User profile first name
--

CREATE TABLE Alias (
	USERID uuid,
	ALIAS varchar(255),
	PUBLIC boolean
);

--
-- Comments:
--
--

-- --------------------------------------------------------

--
-- User profile email
--

CREATE TABLE Email (
	USERID uuid,
	EMAIL varchar(255),
	PUBLIC boolean
);

--
-- Comments:
-- These are optional additional email
-- Primary email is connected to the user in users table
--

-- --------------------------------------------------------

--
-- User profile birth date
--

CREATE TABLE UserBirthdate (
	USERID uuid,
	BIRTHDATE timestamp,
	PUBLIC boolean
);

--
-- Comments:
-- Optional
--

-- --------------------------------------------------------

--
-- User profile picture
--

CREATE TABLE UserAvatar (
	USERID uuid,
	IMGPATH text,
	PUBLIC boolean
);

--
-- Comments:
--
--

-- --------------------------------------------------------

--
-- User description
--

CREATE TABLE Description (
	USERID uuid,
	DESCRIPTION text,
	PUBLIC boolean
);

--
-- Comments:
--
--

-- --------------------------------------------------------

--
-- Table structure for table article
--

DROP TABLE Article;

CREATE TABLE Article (
  ID SERIAL,
  TITLE varchar(255),
  HTML text,
  DATE_CREATED timestamp,
  DATE_UPDATED timestamp
);

--
-- Comments:
--
--

-- --------------------------------------------------------
