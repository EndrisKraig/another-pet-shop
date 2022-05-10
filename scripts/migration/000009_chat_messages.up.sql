BEGIN;

CREATE TABLE conversation_room(
    id BIGSERIAL primary key,
    room_type INTEGER not null,
    room_name VARCHAR(256) not null
);

CREATE TABLE messages(
    id BIGSERIAL primary key,
    profile_id INTEGER not null,
    room_id INTEGER not null,
    creation_date TIMESTAMP not null,
    text_body VARCHAR(512) not null,
    send_status INTEGER not null,
    CONSTRAINT fk_user_profile FOREIGN KEY(profile_id) REFERENCES user_profile(id),
    CONSTRAINT fk_messages_conversation_room FOREIGN KEY(room_id) REFERENCES conversation_room(id)
);

COMMIT;