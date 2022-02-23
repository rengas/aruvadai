CREATE TABLE IF NOT EXISTS product (
	id serial PRIMARY KEY,
	name VARCHAR (256) UNIQUE NOT NULL,
	five_kg int,
	ten_kg int,
	twenty_five_kg int,
	updated_at timestamp  without time zone default (now() at time zone 'utc'),
    created_at timestamp without time zone default (now() at time zone 'utc')
);

INSERT INTO product (name,five_kg,ten_kg,twenty_five_kg)
VALUES
    ('அரிசி',10,2,3),
	('கடலை',0,1,5),
	('உளுந்து',2,20,5),
	('கடலை எண்ணெய்',2,0 ,0),
	('நல்ல எண்ணெய்',4,0,0),
	('தேங்காய் எண்ணெய்',4,0,0);