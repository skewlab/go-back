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
--
--

-- --------------------------------------------------------
