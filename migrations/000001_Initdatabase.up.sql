CREATE TABLE IF NOT EXISTS product (
	id serial PRIMARY KEY,
	name VARCHAR (255) UNIQUE NOT NULL,
	stock JSON,
	updated_at timestamp  without time zone default (now() at time zone 'utc'),
    created_at timestamp without time zone default (now() at time zone 'utc')
);