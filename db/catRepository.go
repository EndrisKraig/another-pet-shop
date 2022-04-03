package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"playground.io/another-pet-store/model"
)

func FindCatById(ID int) model.Cat {
	var conn = getConnection()
	defer conn.Close(context.Background())
	var nickname string
	var id int64
	var breed string
	var price int32
	/* SELECT *, count(*) OVER() AS full_count
	   FROM   tbl
	   WHERE
	   ORDER  BY col1
	   OFFSET ?
	   LIMIT  ? */

	err := conn.QueryRow(context.Background(), "select id, nickname, breed, price from cats where id=$1", ID).Scan(&id, &nickname, &breed, &price)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	return model.Cat{ID: id, Nickname: nickname, Breed: breed, Price: price}
}

func AddCat(cat *model.Cat) {
	var conn = getConnection()
	defer conn.Close(context.Background())
	_, err := conn.Exec(context.Background(), "INSERT INTO cats (nickname, breed, price) VALUES ($1, $2, $3)", cat.Nickname, cat.Breed, cat.Price)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

}

func FindAllCats(page string, limit string) ([]model.Cat, int64) {
	var conn = getConnection()
	defer conn.Close(context.Background())

	pageInt, _ := strconv.ParseInt(page, 10, 64)
	limitInt, _ := strconv.ParseInt(limit, 10, 64)
	//TODO move pagination logic to service
	offset := pageInt*limitInt - limitInt

	rows, err := conn.Query(context.Background(), "SELECT id, nickname, breed, price, count(*) OVER() AS full_count FROM cats ORDER BY id DESC OFFSET $1 LIMIT $2", offset, limitInt)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
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
		full_count = values[4].(int64)
		cats = append(cats, model.Cat{ID: id, Nickname: nickname, Breed: breed, Price: price})
	}
	return cats[:], full_count
}
