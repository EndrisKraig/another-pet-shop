INSERT INTO cats(nickname, breed, price, image_url, user_description, title, age)

SELECT name_t.nickname, breed_t.title, (random()*4000 + 1000)::int, breed_t.link, md5(random()::text), md5(random()::text), (random()*12)::int 
FROM (

WITH  cat_name (id, nickname) AS (VALUES
(0, 'Fluffy'),
(1,'Loki'),
(2,'Oliver'),
(3, 'Bella'),
(4, 'Luna')
) 
SELECT id, nickname, seq FROM (SELECT (random()*4)::int AS rnd, seq  FROM generate_series(1,1000) AS seq) AS f JOIN cat_name ON rnd = cat_name.id

) AS name_t JOIN (

WITH breed (id, title, link) AS (VALUES 
(0 ,'Persian', 'https://upload.wikimedia.org/wikipedia/commons/1/15/White_Persian_Cat.jpg'),
(1, 'Maine Coon', 'https://upload.wikimedia.org/wikipedia/commons/5/5a/Maine_Coon_cat_by_Tomitheos.JPG'),
(2, 'Norwegian Forest cat', 'https://upload.wikimedia.org/wikipedia/en/thumb/7/70/Norwegian_Forest_Cat_in_snow_%28closeup%29_%28cropped%29.jpg/800px-Norwegian_Forest_Cat_in_snow_%28closeup%29_%28cropped%29.jpg')
)
SELECT id, title, link, seq FROM (SELECT (random()*2)::int AS rnd, seq FROM generate_series(1,1000) AS seq) AS t JOIN breed ON rnd = breed.id

) AS breed_t ON breed_t.seq = name_t.seq;

