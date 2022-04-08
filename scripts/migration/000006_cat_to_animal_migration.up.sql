BEGIN;

CREATE TABLE animal_type (
    id BIGSERIAL primary key,
    label VARCHAR(64) NOT NULL,
    is_deleted boolean default '0'
);

INSERT INTO
    animal_type(label)
VALUES
    ('cat');

CREATE TABLE breed (
    id BIGSERIAL primary key,
    label VARCHAR(64) NOT NULL,
    is_deleted boolean default '0'
);

INSERT INTO
    breed (label)
SELECT
    DISTINCT breed
FROM
    cats;

CREATE TABLE custom_field(
    id BIGSERIAL primary key,
    field_data JSONB
);

CREATE TABLE animal (
    id BIGSERIAL primary key,
    nickname VARCHAR(64),
    breed_id INTEGER NOT NULL,
    type_id INTEGER NOT NULL,
    price INTEGER,
    custom_field_id INTEGER,
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP,
    is_deleted boolean default '0',
    buyer_id INTEGER,
    CONSTRAINT fk_animal_breed FOREIGN KEY(breed_id) REFERENCES breed(id),
    CONSTRAINT fk_animal_type FOREIGN KEY(type_id) REFERENCES animal_type(id),
    CONSTRAINT fk_custom_field FOREIGN KEY(custom_field_id) REFERENCES custom_field(id),
    CONSTRAINT fk_buyer_id FOREIGN KEY(buyer_id) REFERENCES user_profile(id)
);

INSERT INTO
    animal (nickname, breed_id, price, created_at, type_id)
SELECT
    nickname,
    breed.id,
    price,
    create_at,
    (
        SELECT
            id
        FROM
            animal_type
        WHERE
            label = 'cat'
    )
FROM
    cats
    JOIN breed ON cats.breed = breed.label;

ALTER TABLE
    user_profile
ADD
    column nickname VARCHAR(64);

INSERT into
    user_profile(nickname)
SELECT
    username
FROM
    user_profile
    JOIN users on user_profile.user_id = users.id;

COMMIT;