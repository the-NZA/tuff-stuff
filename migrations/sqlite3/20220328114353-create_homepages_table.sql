
-- +migrate Up
CREATE TABLE IF NOT EXISTS homepages(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    lang TEXT UNIQUE NOT NULL,
    content TEXT NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS homepages;