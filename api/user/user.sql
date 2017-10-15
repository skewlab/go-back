CREATE EXTENSION pgcrypto;

CREATE TABLE Users (
  ID uuid NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
  USERNAME text NOT NULL,
  PASSWORD text NOT NULL
);

/*
Insert query:
INSERT INTO Users ( name, password ) VALUES ( 'admin', crypt( 'password', gen_salt( 'bf', 8 ) ) );
*/
