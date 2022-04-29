
-- +migrate Up
CREATE TABLE IF NOT EXISTS image_grid(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    image_id INTEGER NOT NULL,
    FOREIGN KEY(image_id) REFERENCES images(id)
);

-- +migrate Down
DROP TABLE IF EXISTS image_grid;
