package db

import (
	"context"
	"fmt"

	"playground.io/another-pet-store/model"
)

type AnimalRepository interface {
	FindAnimalById(ID int) (*model.Animal, error)
	AddAnimal(animal model.Animal) error
	UpdateAnimal(animal model.Animal) error
	FindAllAnimals(offset int, limit int) ([]model.Animal, int, error)
}

type SimpleAnimalRepository struct {
}

func NewAnimalRepository() AnimalRepository {
	return &SimpleAnimalRepository{}
}

func (repository *SimpleAnimalRepository) FindAnimalById(ID int) (*model.Animal, error) {
	var conn = getConnection()
	defer conn.Close(context.Background())

	var id int64
	var nickname string
	var breed string
	var animaltype string
	var price int32
	var imageUrl string
	var age int32
	var title string
	const query = `SELECT a.id, nickname, b.label, at.label, price, image_url, age, title
				   FROM animal AS a JOIN breed AS b ON a.breed_id = b.id JOIN animal_type AS at ON type_id = at.id
				   WHERE a.is_deleted = '0' AND a.id = $1`
	err := conn.QueryRow(context.Background(), query, ID).Scan(&id, &nickname, &breed, &animaltype, &price, &imageUrl, &age, &title)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error animal with id %d doesn't exist: %v", ID, err)
	}
	return &model.Animal{ID: id, Nickname: nickname, Breed: breed, Price: price, Age: age, ImageUrl: imageUrl, Title: title, Type: animaltype}, nil
}

func (repository *SimpleAnimalRepository) AddAnimal(animal model.Animal) error {
	var conn = getConnection()
	defer conn.Close(context.Background())
	//TODO fix insert, new constraints require type as well
	const query = `INSERT INTO animal (nickname, breed_id, price)
				   VALUES ($1, (SELECT id FROM breed WHERE label $2), $3)`
	_, err := conn.Exec(context.Background(), query, animal.Nickname, animal.Breed, animal.Price)

	if err != nil {
		return fmt.Errorf("error execute insert command: %w", err)
	}
	return nil
}

func (repository *SimpleAnimalRepository) UpdateAnimal(animal model.Animal) error {
	var conn = getConnection()
	defer conn.Close(context.Background())
	const query = `UPDATE animal
				   SET nickname=$1, breed_id=$(SELECT id FROM breed WHERE label = $2), price=$3, buyer_id=$4, type=$5
				   WHERE id = $6`
	_, err := conn.Exec(context.Background(), query, animal.Nickname, animal.Breed, animal.Price, animal.BuyerId, animal.Type, animal.ID)

	if err != nil {
		return fmt.Errorf("error execute insert command: %w", err)
	}
	return nil
}

func (repository *SimpleAnimalRepository) FindAllAnimals(offset int, limit int) ([]model.Animal, int, error) {
	var conn = getConnection()
	defer conn.Close(context.Background())
	const query = `SELECT a.id, nickname, b.label, at.label, price, image_url, age, title, count(*) OVER() AS full_count 
				   FROM animal AS a JOIN breed AS b ON a.breed_id = b.id JOIN animal_type AS at ON type_id = at.id
				   WHERE a.is_deleted = '0' ORDER BY id DESC OFFSET $1 LIMIT $2`
	rows, err := conn.Query(context.Background(), query, offset, limit)

	if err != nil {
		return nil, 0, fmt.Errorf("error during update command %w", err)
	}

	defer rows.Close()
	var animals []model.Animal

	var full_count int64 = 0

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, 0, fmt.Errorf("error during obtaining result rows values: %w", err)
		}

		id := values[0].(int64)
		nickname := values[1].(string)
		breed := values[2].(string)
		animalType := values[3].(string)
		price := values[4].(int32)
		imageUrl := values[5].(string)
		age := values[6].(int32)
		title := values[7].(string)
		full_count = values[8].(int64)
		animals = append(animals, model.Animal{ID: id, Nickname: nickname, Breed: breed, Price: price, ImageUrl: imageUrl, Title: title, Age: age, Type: animalType})
	}
	return animals[:], int(full_count), nil
}
