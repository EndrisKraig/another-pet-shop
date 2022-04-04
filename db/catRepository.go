package db

import (
	"context"
	"log"
	"time"

	"playground.io/another-pet-store/model"
)

func FindCatById(ID int) model.Cat {
	var conn = getConnection()
	defer conn.Close(context.Background())
	var nickname string
	var id int64
	var breed string
	var price int32

	err := conn.QueryRow(context.Background(), "select id, nickname, breed, price from cats where id=$1", ID).Scan(&id, &nickname, &breed, &price)
	if err != nil {
		log.Fatal("QueryRow failed: $1\n", err)
	}
	return model.Cat{ID: id, Nickname: nickname, Breed: breed, Price: price}
}

func AddCat(cat model.Cat) {
	var conn = getConnection()
	defer conn.Close(context.Background())
	_, err := conn.Exec(context.Background(), "INSERT INTO cats (nickname, breed, price) VALUES ($1, $2, $3)", cat.Nickname, cat.Breed, cat.Price)

	if err != nil {
		log.Fatal("QueryRow failed: $1\n", err)
	}
}

func FindAllCats(offset int, limit int) ([]model.Cat, int) {
	var conn = getConnection()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT id, nickname, breed, price, create_at, image_url, title, age, count(*) OVER() AS full_count FROM cats ORDER BY id DESC OFFSET $1 LIMIT $2", offset, limit)

	if err != nil {
		log.Fatal("QueryRow failed: $1\n", err)
	}

	var cats []model.Cat

	var full_count int64 = 0

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		id := values[0].(int64)
		nickname := values[1].(string)
		breed := values[2].(string)
		price := values[3].(int32)
		createAt := values[4].(time.Time)
		imageUrl := values[5].(string)
		title := values[6].(string)
		age := values[7].(int32)
		full_count = values[8].(int64)
		cats = append(cats, model.Cat{ID: id, Nickname: nickname, Breed: breed, Price: price, CreateAt: createAt.Format("2006-01-02 15:04:05"), ImageUrl: imageUrl, Title: title, Age: age})
	}
	return cats[:], int(full_count)
}
