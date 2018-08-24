CREATE DATABASE test_db owner postgres encoding 'utf8';
\c test_db;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username character varying(40) NOT NULL UNIQUE,
    full_name character varying(217),
    password character varying(217),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);

INSERT INTO users values(1, 'ma1', 'Master Git 1', '1234');
INSERT INTO users values(2, 'ma2', 'Master Git 2', '1234');
INSERT INTO users values(3, 'ma3', 'Master Git 3', '1234');
INSERT INTO users values(4, 'ma4', 'Master Git 4', '1234');