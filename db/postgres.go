package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
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

func FindAllCats() []model.Cat {
	var conn = getConnection()
	defer conn.Close(context.Background())
	rows, err := conn.Query(context.Background(), "select id, nickname, breed, price from public.cats")

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	var cats []model.Cat

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		id := values[0].(int64)
		nickname := values[1].(string)
		breed := values[2].(string)
		price := values[3].(int32)

		cats = append(cats, model.Cat{ID: id, Nickname: nickname, Breed: breed, Price: price})
	}
	return cats[:]
}

//INSERT INTO "cat" ("nickname", "breed", "price")
//VALUES ('fluffy', 'siberian', '1000');

func getConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://supercat:meow_meow@localhost:5432/supercat")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}
