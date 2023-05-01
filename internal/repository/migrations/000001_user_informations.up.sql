CREATE TABLE IF NOT EXISTS user_informations(
	id SERIAL PRIMARY KEY,
	first_name VARCHAR(30),
	last_name VARCHAR(30),
	address VARCHAR(100)
);
