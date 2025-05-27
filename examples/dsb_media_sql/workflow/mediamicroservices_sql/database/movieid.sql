CREATE TABLE IF NOT EXISTS movieid (
	movieid   VARCHAR(25), 
	title     VARCHAR(150) UNIQUE,
	PRIMARY KEY(movieid)
);
