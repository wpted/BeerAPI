CREATE USER beer_fellow WITH PASSWORD 'lovebeer';
ALTER USER beer_fellow WITH SUPERUSER;
ALTER ROLE beer_fellow CREATEROLE CREATEDB;

CREATE DATABASE beer_server;
GRANT ALL PRIVILEGES ON DATABASE beer_server to beer_fellow;

CREATE TABLE beers(
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(50) NOT NULL,
    Price int NOT NULL,
    Company VARCHAR(50) NOT NULL
);