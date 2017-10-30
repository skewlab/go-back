select * from Users;--
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
	'https://cdn.pixabay.com/photo/2015/09/08/17/35/man-930397_960_720.jpg',
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
	'https://cdn.pixabay.com/photo/2016/06/22/21/18/cat-1474092_960_720.jpg',
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
	'https://cdn.pixabay.com/photo/2016/02/18/22/16/smile-1208203_960_720.jpg',
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
	'https://cdn.pixabay.com/photo/2015/04/20/13/32/woman-731377_960_720.jpg',
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
	'https://cdn.pixabay.com/photo/2017/01/23/19/40/beautiful-girl-2003647_960_720.jpg',
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
	'https://cdn.pixabay.com/photo/2017/09/21/09/25/teddy-bear-2771252_960_720.jpg',
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