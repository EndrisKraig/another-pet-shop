BEGIN;

ALTER TABLE animal DROP column age;
ALTER TABLE animal DROP column image_url;
ALTER TABLE animal DROP column title;

COMMIT;