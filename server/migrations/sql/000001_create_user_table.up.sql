CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL,
    use TEXT NOT NULL,
    avatar varchar(100) NOT NULL,
    description TEXT NOT NULL,
    rating SMALLINT DEFAULT 10000
);