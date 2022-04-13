CREATE TABLE special_offer (
    id BIGSERIAL primary key,
    animal_id INTEGER not null,
    begin_date TIMESTAMP not null,
    end_date TIMESTAMP not null,
    conditions VARCHAR(256),
    CONSTRAINT fk_specila_offer_animal FOREIGN KEY(animal_id) REFERENCES animal(id)
);