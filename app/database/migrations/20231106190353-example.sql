-- +migrate Up
CREATE TABLE example (
    id INTEGER PRIMARY KEY
);
-- +migrate Down
DROP TABLE example;
