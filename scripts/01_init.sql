-- CREATE USER bypass;
-- CREATE DATABASE bypass;
-- GRANT ALL PRIVILEGES ON DATABASE bypass TO bypass;

CREATE SCHEMA logbook;

CREATE TABLE logbook.books (
id SERIAL PRIMARY KEY,
title VARCHAR(100),
author VARCHAR(100),
year VARCHAR(4),
publisher VARCHAR(100),
readtime TIMESTAMP,
rating VARCHAR(3),
comments VARCHAR(300),
language VARCHAR(50),
genre VARCHAR(50),
isbn VARCHAR(50)
);