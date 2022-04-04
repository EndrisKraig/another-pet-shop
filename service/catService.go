package service

import (
	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/model"
	"playground.io/another-pet-store/utils"
)

type CatService interface {
	FindCatById(id string) model.Cat
	FindAllCats(page, limit string) dto.CatsResponse
	AddCat(cat *dto.Cat) dto.Cat
}

type CatServiceInstance struct {
}

func (c *CatServiceInstance) FindCatById(id string) dto.Cat {
	idAsInt := utils.StringToInt(id)
	var cat model.Cat = db.FindCatById(idAsInt)
	return dto.Cat{ID: cat.ID, Nickname: cat.Nickname, Breed: cat.Breed, Price: cat.Price}
}

func (c *CatServiceInstance) FindAllCats(pageStr, limitStr string) dto.CatsResponse {
	page := utils.StringToInt(pageStr)
	limit := utils.StringToInt(limitStr)
	// page starts from 1, offset from 0
	var offset = (page - 1) * limit
	cats, catsNum := db.FindAllCats(offset, limit)
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
	return dto.CatsResponse{Cats: dtoCats, MaxPage: maxPage}
}

func (c *CatServiceInstance) AddCat(cat *dto.Cat) dto.Cat {
	var newCat model.Cat = model.Cat{Nickname: cat.Nickname, Breed: cat.Breed, Price: cat.Price}
	db.AddCat(newCat)
	// cat.ID = addedCat.ID
	return *cat
}
