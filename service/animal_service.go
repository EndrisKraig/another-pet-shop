package service

import (
	"fmt"
	"strconv"

	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/model"
)

type AnimalService interface {
	FindAnimalById(id string) (*dto.Animal, error)
	FindAllAnimals(page, limit int) (*dto.AnimalResponse, error)
	AddAnimal(animal *dto.Animal) error
	UpdateAnimal(animal string, token string) error
}

type AnimalServiceInstance struct {
	profileService   ProfileService
	animalRepository db.AnimalRepository
}

func CreateAnimalService(profileService ProfileService, animalRepository db.AnimalRepository) AnimalService {
	return &AnimalServiceInstance{profileService: profileService, animalRepository: animalRepository}
}

func (c AnimalServiceInstance) FindAnimalById(id string) (*dto.Animal, error) {
	idAsInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("Couldn't pars %s as string: %w", id, err)
	}
	animal, err := c.animalRepository.FindAnimalById(idAsInt)
	if err != nil {
		return nil, fmt.Errorf("Error during find animal with id %d: %w", idAsInt, err)
	}
	return &dto.Animal{ID: animal.ID, Nickname: animal.Nickname, Breed: animal.Breed, Price: animal.Price, ImageUrl: animal.ImageUrl, Age: animal.Age, Title: animal.Title}, nil
}

func (c AnimalServiceInstance) FindAllAnimals(page, limit int) (*dto.AnimalResponse, error) {

	// page starts from 1, offset from 0
	var offset = (page - 1) * limit
	animals, animalsNum, err := c.animalRepository.FindAllAnimals(offset, limit)

	if err != nil {
		return nil, fmt.Errorf("Error looking for animals: %w", err)
	}

	var dtoAnimals []dto.Animal
	for _, animal := range animals {
		dtoAnimals = append(dtoAnimals, dto.Animal{ID: animal.ID, Nickname: animal.Nickname, Breed: animal.Breed, Price: animal.Price, CreateAt: animal.CreateAt, ImageUrl: animal.ImageUrl, Title: animal.Title, Age: animal.Age})
	}
	var maxPage = 0
	if animalsNum%limit == 0 {
		maxPage = animalsNum / limit
	} else {
		maxPage = animalsNum/limit + 1
	}
	return &dto.AnimalResponse{Animals: dtoAnimals, MaxPage: maxPage}, nil
}

func (c AnimalServiceInstance) AddAnimal(animal *dto.Animal) error {
	var newAnimal model.Animal = model.Animal{Nickname: animal.Nickname, Breed: animal.Breed, Price: animal.Price}
	err := c.animalRepository.AddAnimal(newAnimal)
	if err != nil {
		return fmt.Errorf("Couldn't add animal: %w", err)
	}
	return nil
}

func (service AnimalServiceInstance) UpdateAnimal(animalId string, token string) error {
	profile, err := service.profileService.GetProfile(token)
	if err != nil {
		return fmt.Errorf("Error during updating animal: %w", err)
	}
	animal, err := service.FindAnimalById(animalId)
	if err != nil {
		return fmt.Errorf("No animal with id %s: %v", animalId, err)
	}
	//TODO transaction
	modelAnimal := model.Animal{ID: animal.ID, Nickname: animal.Nickname, Breed: animal.Breed, Price: animal.Price, BuyerId: profile.Id}
	err = service.animalRepository.UpdateAnimal(modelAnimal)
	if err != nil {
		return fmt.Errorf("Animal wasn't sold: %w", err)
	}
	err = service.profileService.ChangeBalance(float32(animal.Price), profile)
	if err != nil {
		return fmt.Errorf("Animal wasn't sold: %w", err)
	}
	return nil
}
