
-- +migrate Up
CREATE TABLE IF NOT EXISTS commission_cards(
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   title TEXT NOT NULL,
   content TEXT NOT NULL,
   lang TEXT NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS commission_cards;
