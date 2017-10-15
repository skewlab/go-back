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
  USERNAME text NOT NULL,
  PASSWORD text NOT NULL
);

/*
Comments:
Insert query:
  INSERT INTO
  Users ( name, password )
  VALUES ( 'admin', crypt( 'password', gen_salt( 'bf', 8 ) ) );
*/

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

-- --------------------------------------------------------
