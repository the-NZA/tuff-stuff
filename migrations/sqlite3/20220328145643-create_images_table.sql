
-- +migrate Up
CREATE TABLE IF NOT EXISTS images(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    url TEXT NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS images;