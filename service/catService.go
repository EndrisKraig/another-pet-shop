package service

import (
	"fmt"
	"strconv"

	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/model"
)

type CatService interface {
	FindCatById(id string) (*dto.Cat, error)
	FindAllCats(page, limit int) (*dto.CatsResponse, error)
	AddCat(cat *dto.Cat) error
}

type CatServiceInstance struct {
}

func (c CatServiceInstance) FindCatById(id string) (*dto.Cat, error) {
	idAsInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("Couldn't pars %s as string: %w", id, err)
	}
	cat, err := db.FindCatById(idAsInt)
	if err != nil {
		return nil, fmt.Errorf("Error during find cat with id %d: %w", idAsInt, err)
	}
	return &dto.Cat{ID: cat.ID, Nickname: cat.Nickname, Breed: cat.Breed, Price: cat.Price}, nil
}

func (c CatServiceInstance) FindAllCats(page, limit int) (*dto.CatsResponse, error) {

	// page starts from 1, offset from 0
	var offset = (page - 1) * limit
	cats, catsNum, err := db.FindAllCats(offset, limit)

	if err != nil {
		return nil, fmt.Errorf("Error looking for cats: %w", err)
	}

	var dtoCats []dto.Cat
	for _, cat := range cats {
		dtoCats = append(dtoCats, dto.Cat{ID: cat.ID, Nickname: cat.Nickname, Breed: cat.Breed, Price: cat.Price, CreateAt: cat.CreateAt, ImageUrl: cat.ImageUrl, Title: cat.Title, Age: cat.Age})
	}
	var maxPage = 0
	if catsNum%limit == 0 {
		maxPage = catsNum / limit
	} else {
		maxPage = catsNum/limit + 1
	}
	return &dto.CatsResponse{Cats: dtoCats, MaxPage: maxPage}, nil
}

func (c CatServiceInstance) AddCat(cat *dto.Cat) error {
	var newCat model.Cat = model.Cat{Nickname: cat.Nickname, Breed: cat.Breed, Price: cat.Price}
	err := db.AddCat(newCat)
	if err != nil {
		return fmt.Errorf("Couldn't add cat: %w", err)
	}
	// cat.ID = addedCat.ID
	return nil
}
