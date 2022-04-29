
-- +migrate Up
CREATE TABLE IF NOT EXISTS how_works_cards(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    lang TEXT NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS how_works_cards;