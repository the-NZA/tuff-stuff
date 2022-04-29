
-- +migrate Up
CREATE TABLE IF NOT EXISTS why_us_cards(
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   title TEXT NOT NULL,
   content TEXT NOT NULL,
   lang TEXT NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS why_us_cards;