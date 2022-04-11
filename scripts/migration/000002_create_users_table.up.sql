CREATE TABLE users (
    id BIGSERIAL primary key,
    username VARCHAR(64) NOT NULL UNIQUE,
    pass_hash TEXT NOT NULL,
    email VARCHAR(128) UNIQUE,
    create_at TIMESTAMP default now()
);