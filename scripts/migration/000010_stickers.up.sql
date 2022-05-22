BEGIN;

CREATE TABLE message_format (
    id BIGSERIAL primary key,
    label VARCHAR(64)
);

INSERT INTO message_format(label) VALUES ('text'), ('sticker');

ALTER TABLE
    messages
ADD
    column format_id INTEGER;

UPDATE messages
SET format_id = (SELECT id FROM message_format WHERE label = 'text')
WHERE 1 = 1;


ALTER TABLE messages
ADD FOREIGN KEY (format_id) REFERENCES message_format(id);

ALTER TABLE messages alter column format_id SET NOT NULL;

CREATE TABLE kit(
    id BIGSERIAL primary key,
    label VARCHAR(64)
);

INSERT INTO kit(label) VALUES ('basic');

CREATE TABLE sticker(
    id BIGSERIAL primary key,
    kit_id INTEGER NOT NULL,
    uri VARCHAR(256) NOT NULL,
    CONSTRAINT fk_stiker_kit FOREIGN KEY(kit_id) REFERENCES kit(id)
);

COMMIT;