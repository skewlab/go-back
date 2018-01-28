--
-- Init script for database
-- Adding test users and relations among the users
--

-- --------------------------------------------------------

-- Empty table
DELETE FROM Users;

--
-- Add users
--

-- Simon Garfunkel
INSERT INTO Users ( id, email, password, alias, avatar, description, website, phonenumber )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
	'simon@garfunkel.com',
	crypt( 'secret-simon', gen_salt( 'bf', 8 ) ),
	'Simon Garfunkel',
	'http://www.fakepersongenerator.com/Face/male/male20161085952333985.jpg',
	'I like to play music.',
	'www.google.com',
	'0701928374'
);

-- Pussy Pet
INSERT INTO Users ( id, email, password, alias, avatar, description, website, phonenumber )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaab',
	'pussy@pet.com',
	crypt( 'secret-pussy', gen_salt( 'bf', 8 ) ),
	'Pussy Pet',
	'http://www.fakepersongenerator.com/Face/male/male1085432314820.jpg',
	'I love pussy pussy.',
	'www.google.com',
	'0701928374'
);

-- Joe Man
INSERT INTO Users ( id, email, password, alias, avatar, description, website, phonenumber )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaac',
	'joe@man.org',
	crypt( 'secret-joe', gen_salt( 'bf', 8 ) ),
	'Joe Man',
	'http://www.fakepersongenerator.com/Face/male/male20111086094401184.jpg',
	'Hi Im Joe. Who are you? I like fishing and conversations about men.',
	'www.google.com',
	'0701928374'
);

-- Leila Skypetalker
INSERT INTO Users ( id, email, password, alias, avatar, description, website, phonenumber )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaad',
	'leila_skypetalker@galaxies.net',
	crypt( 'secret-leila', gen_salt( 'bf', 8 ) ),
	'Leila Skypetalker',
	'http://www.fakepersongenerator.com/Face/female/female102241340711.jpg',
	'I didn´t do it with my brother or with the big fat slimy hot dog looking fellow.',
	'http://galaxies.net/',
	'0701928374'
);

--
INSERT INTO Users ( id, email, password, alias, avatar, description, website, phonenumber )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaae',
	'gerly@gerls.org',
	crypt( 'secret-gerly', gen_salt( 'bf', 8 ) ),
	'Gerly Gerl',
	'http://www.fakepersongenerator.com/Face/female/female20161025522236115.jpg',
	'I like to read books and lie down in the leaves and get my knitted sweater dirty.',
	'',
	'0701928374'
);

-- Bear Bones
INSERT INTO Users ( id, email, password, alias, avatar, description, website, phonenumber )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaf',
	'bearbones@forestsandthat.com',
	crypt( 'secret-bear', gen_salt( 'bf', 8 ) ),
	'Bear Bones',
	'http://www.fakepersongenerator.com/Face/male/male20151083626635132.jpg',
	'I don´t have any firends... Please hug me.',
	'http://hello.world',
	'0701928374'
);

-- --------------------------------------------------------

-- Empty table
DELETE FROM UserConnections;

--
-- Insert relations
--

INSERT INTO UserConnections( RequestingUser, RespondingUser, Accepted )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaab',
	true
);

INSERT INTO UserConnections( RequestingUser, RespondingUser, Accepted )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaac',
	true
);

INSERT INTO UserConnections( RequestingUser, RespondingUser, Accepted )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaad',
	true
);

INSERT INTO UserConnections( RequestingUser, RespondingUser, Accepted )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaae',
	false
);

INSERT INTO UserConnections( RequestingUser, RespondingUser, Accepted )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaf',
	false
);

INSERT INTO UserConnections( RequestingUser, RespondingUser, Accepted )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaab',
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaac',
	false
);

INSERT INTO UserConnections( RequestingUser, RespondingUser, Accepted )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaab',
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaad',
	true
);

INSERT INTO UserConnections( RequestingUser, RespondingUser, Accepted )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaac',
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaf',
	false
);

-- --------------------------------------------------------

INSERT INTO Posts ( userId, content, date_created, date_updated )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
	'Post number one from simon garfunkel',
	'2017-10-30 14:39:41.17292',
	'2017-10-30 14:39:41.17292'
);

INSERT INTO Posts ( userId, content, date_created, date_updated )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
	'Post number two from simon garfunkel',
	'2017-10-30 14:39:41.17292',
	'2017-10-30 14:39:41.17292'
);

INSERT INTO Posts ( userId, content, date_created, date_updated )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
	'Post number three from simon garfunkel',
	'2017-10-30 14:39:41.17292',
	'2017-10-30 14:39:41.17292'
);

INSERT INTO Posts ( userId, content, date_created, date_updated )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaab',
	'Post number one from pussy pet',
	'2017-10-30 14:39:41.17292',
	'2017-10-30 14:39:41.17292'
);

INSERT INTO Posts ( userId, content, date_created, date_updated )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaab',
	'Post number two from pussy pet',
	'2017-10-30 14:39:41.17292',
	'2017-10-30 14:39:41.17292'
);

INSERT INTO Posts ( userId, content, date_created, date_updated )
Values(
	'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaac',
	'Post number one from Joe',
	'2017-10-30 14:39:41.17292',
	'2017-10-30 14:39:41.17292'
);