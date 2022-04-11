INSERT INTO
    animal(
        nickname,
        breed_id,
        type_id,
        price,
        image_url,
        title,
        age
    )
SELECT
    name_t.nickname,
    breed_t.breed_id,
    (
        SELECT
            id
        FROM
            animal_type
        WHERE
            label = 'cat'
    ),
    (random() * 2000 + 1000) :: int,
    breed_t.link,
    md5(random() :: text),
    (random() * 12) :: int
FROM
    (
        WITH cat_name (id, nickname) AS (
            VALUES
                (0, 'Fluffy'),
                (1, 'Loki'),
                (2, 'Oliver'),
                (3, 'Bella'),
                (4, 'Luna')
        )
        SELECT
            id,
            nickname,
            seq
        FROM
            (
                SELECT
                    (random() * 4) :: int AS rnd,
                    seq
                FROM
                    generate_series(1, 1000) AS seq
            ) AS f
            JOIN cat_name ON rnd = cat_name.id
    ) AS name_t
    JOIN (
        WITH breed_tmp (id, title, link) AS (
            VALUES
                (
                    0,
                    'Persian',
                    'https://upload.wikimedia.org/wikipedia/commons/1/15/White_Persian_Cat.jpg'
                ),
                (
                    1,
                    'Maine Coon',
                    'https://upload.wikimedia.org/wikipedia/commons/5/5a/Maine_Coon_cat_by_Tomitheos.JPG'
                ),
                (
                    2,
                    'Norwegian Forest cat',
                    'https://upload.wikimedia.org/wikipedia/en/thumb/7/70/Norwegian_Forest_Cat_in_snow_%28closeup%29_%28cropped%29.jpg/800px-Norwegian_Forest_Cat_in_snow_%28closeup%29_%28cropped%29.jpg'
                )
        )
        SELECT
            breed_tmp.id,
            title,
            link,
            seq,
            breed.id breed_id
        FROM
            (
                SELECT
                    (random() * 2) :: int AS rnd,
                    seq
                FROM
                    generate_series(1, 1000) AS seq
            ) AS t
            JOIN breed_tmp ON rnd = breed_tmp.id
            JOIN breed ON breed_tmp.title = breed.label
    ) AS breed_t ON breed_t.seq = name_t.seq;