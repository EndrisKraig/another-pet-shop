package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"playground.io/another-pet-store/model"
)

type AnimalRepository interface {
	FindAnimalById(ID int) (*model.Animal, error)
	AddAnimal(animal model.Animal) error
	UpdateAnimal(animal model.Animal) error
	FindAllAnimals(offset int, limit int) ([]model.Animal, int, error)
	SellAnimal(animalId, profileId int, balanceCalc func(int, int) (int, error)) error
}

type SimpleAnimalRepository struct {
}

func NewAnimalRepository() AnimalRepository {
	return &SimpleAnimalRepository{}
}

func (repository *SimpleAnimalRepository) FindAnimalById(ID int) (*model.Animal, error) {
	conn, err := GetConnection()

	if err != nil {
		return nil, err
	}

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
	err = conn.QueryRow(context.Background(), query, ID).Scan(&id, &nickname, &breed, &animaltype, &price, &imageUrl, &age, &title)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error animal with id %d doesn't exist: %v", ID, err)
	}
	return &model.Animal{ID: id, Nickname: nickname, Breed: breed, Price: price, Age: age, ImageUrl: imageUrl, Title: title, Type: animaltype}, nil
}

func (repository *SimpleAnimalRepository) AddAnimal(animal model.Animal) error {
	conn, err := GetConnection()

	if err != nil {
		return err
	}
	//TODO fix insert, new constraints require type as well
	const query = `INSERT INTO animal (nickname, breed_id, type_id, price, age, image_url, title)
				   VALUES ($1, (SELECT id FROM breed WHERE label = $2), (SELECT id FROM animal_type WHERE label = $3), $4, $5,$6,$7)`
	_, err = conn.Exec(context.Background(), query, animal.Nickname, animal.Breed, animal.Type, animal.Price, animal.Age, animal.ImageUrl, animal.Title)

	if err != nil {
		return fmt.Errorf("error execute insert command: %w", err)
	}
	return nil
}

func (repository *SimpleAnimalRepository) UpdateAnimal(animal model.Animal) error {
	conn, err := GetConnection()

	if err != nil {
		return err
	}

	const query = `UPDATE animal
				   SET buyer_id=$1
				   WHERE id = $2`
	_, err = conn.Exec(context.Background(), query, animal.BuyerId, animal.ID)

	if err != nil {
		return fmt.Errorf("error execute insert command: %w", err)
	}
	return nil
}

func (repository *SimpleAnimalRepository) FindAllAnimals(offset int, limit int) ([]model.Animal, int, error) {
	conn, err := GetConnection()

	if err != nil {
		return nil, 0, err
	}
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

func (repository *SimpleAnimalRepository) SellAnimal(animalId, profileId int, balanceCalc func(int, int) (int, error)) error {
	conn, err := GetConnection()

	if err != nil {
		return err
	}

	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		} else {
			tx.Commit(context.Background())
		}
	}()
	const query = `UPDATE animal
				   SET buyer_id=$1
				   WHERE id = $2`

	_, err = tx.Exec(context.Background(), query, profileId, animalId)

	if err != nil {
		return err
	}

	var animalPrice int64
	var balance int32

	const animalQuery = `SELECT price, balance
	FROM animal JOIN user_profile ON user_profile.id = animal.buyer_id
	WHERE animal.id=$1`

	tx.QueryRow(context.Background(), animalQuery, animalId).Scan(&animalPrice, &balance)
	newBalance, err := balanceCalc(int(animalPrice), int(balance))

	if err != nil {
		return err
	}

	const updateBalanceQuery = `UPDATE user_profile
	SET balance = $1
	WHERE id = $2;`
	fmt.Println(newBalance)
	_, err = tx.Exec(context.Background(), updateBalanceQuery, newBalance, profileId)

	return err
}
