BEGIN;

ALTER TABLE
    user_profile DROP column nickname;

DROP TABLE IF EXISTS animal;
DROP TABLE IF EXISTS breed;
DROP TABLE IF EXISTS animal_type;
DROP TABLE IF EXISTS custom_field;


COMMIT;