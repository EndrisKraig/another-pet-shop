BEGIN;

ALTER TABLE
    messages DROP column format_id;

DROP TABLE IF EXISTS message_format; 

DROP TABLE IF EXISTS sticker;

DROP TABLE IF EXISTS kit;

COMMIT;