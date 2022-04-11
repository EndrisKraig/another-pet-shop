CREATE TABLE user_profile (
    id BIGSERIAL primary key,
    user_id BIGSERIAL,
    balance NUMERIC default 10000,
    image_url VARCHAR(256),
    notes TEXT
);