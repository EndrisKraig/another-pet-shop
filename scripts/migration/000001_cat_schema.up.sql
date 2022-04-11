CREATE TABLE cats (
    id BIGSERIAL primary key,
    nickname VARCHAR(64),
    breed VARCHAR(64),
    price integer,
    create_at TIMESTAMP default now()
);