CREATE TABLE users (
	user_id int UNIQUE NOT NULL,
	created_at varchar(150) NOT NULL,
	updated_at varchar(150) NOT NULL,
	
	PRIMARY KEY (user_id)
);

CREATE TABLE slugs (
	slug_id serial NOT NULL,
	slug_name varchar(150) UNIQUE NOT NULL,
	created_at varchar(150) NOT NULL,
	updated_at varchar(150) NOT NULL,
	deleted_at varchar(150),
	disabled int NOT NULL,
	
	PRIMARY KEY (slug_id)
);

CREATE TABLE users_slugs (
	user_user_id integer NOT NULL,
	slug_slug_id integer NOT NULL,
	created_at varchar(150),
	updated_at varchar(150),
	deleted_at varchar(150),
	until_date varchar(150),
	
	PRIMARY KEY(user_user_id, slug_slug_id),
	FOREIGN KEY (user_user_id) REFERENCES users(user_id),
	FOREIGN KEY (slug_slug_id) REFERENCES slugs(slug_id)
);

CREATE TABLE histories (
	operation_id serial NOT NULL,
	user_user_id integer NOT NULL,
	slug_name varchar(150) NOT NULL,
	operation varchar(50) NOT NULL,
	date_info varchar(150) NOT NULL,
	
	PRIMARY KEY (operation_id)
);
