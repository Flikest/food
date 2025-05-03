CREATE TABLE users(
    id INT NOT NULL,
    use TEXT NOT NULL,
    avatar BLOB,
    description TEXT NOT NULL,
    rating SMALLINT NOT NULL DEFAULT 10000
)