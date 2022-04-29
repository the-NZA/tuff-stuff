
-- +migrate Up
CREATE TABLE IF NOT EXISTS options(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE NOT NULL,
    title TEXT NOT NULL,
    value TEXT NOT NULL
);
-- +migrate Down
DROP TABLE IF EXISTS options;