BEGIN;

ALTER TABLE
    animal
ADD
    column age INTEGER;

ALTER TABLE
    animal
ADD
    column image_url VARCHAR(256);

ALTER TABLE
    animal
ADD
    column title VARCHAR(128);

/* Migration failed cats were lost :), happily it was only test data...*/
DELETE
FROM animal
WHERE 1=1;

COMMIT;