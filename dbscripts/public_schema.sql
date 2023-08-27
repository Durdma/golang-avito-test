SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET client_min_messages = warning;
SET row_security = off;
CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA pg_catalog;
SET search_path = public, pg_catalog;
SET default_tablespace = '';
--users
CREATE TABLE users (
	user_id int UNIQUE NOT NULL,
	
	PRIMARY KEY (user_id)
);
--slugs
CREATE TABLE slugs (
	slug_id serial NOT NULL,
	slug_name varchar(150) UNIQUE NOT NULL,
	
	PRIMARY KEY (slug_id)
);
--users_slugs
CREATE TABLE users_slugs (
	fk_user_id integer NOT NULL,
	fk_slug_id integer NOT NULL,
	
	PRIMARY KEY(fk_user_id, fk_slug_id),
	FOREIGN KEY(fk_user_id) REFERENCES users(user_id) ON DELETE NO ACTION ON UPDATE NO ACTION,
	FOREIGN KEY(fk_slug_id) REFERENCES slugs(slug_id) ON DELETE NO ACTION ON UPDATE NO ACTION
);