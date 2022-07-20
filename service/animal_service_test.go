package service_test

import (
	"fmt"
	"testing"

	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/model"
	"playground.io/another-pet-store/service"
)

type StubProfileService struct {
}

func (s *StubProfileService) CreateProfile(userId int) error {
	return nil
}
func (s *StubProfileService) GetProfile(userId int) (*dto.Profile, error) {
	return &dto.Profile{}, nil
}
func (s *StubProfileService) ChangeBalance(animalPrice, profileCash int) (int, error) {
	return 0, nil
}
func (s *StubProfileService) ChangeBalanceFunc() func(int, int) (int, error) {
	return s.ChangeBalance
}

type StubAnimalRepository struct {
	animals map[int]model.Animal
}

func (r *StubAnimalRepository) FindAnimalById(ID int) (*model.Animal, error) {
	animal, ok := r.animals[ID]
	if !ok {
		return nil, fmt.Errorf("No such animal id=%d", ID)
	}
	return &animal, nil
}
func (r *StubAnimalRepository) AddAnimal(animal model.Animal) error {
	animal.ID = 10
	r.animals[10] = animal
	return nil
}
func (r *StubAnimalRepository) UpdateAnimal(animal model.Animal) error {
	r.animals[int(animal.ID)] = animal
	return nil
}
func (r *StubAnimalRepository) FindAllAnimals(offset int, limit int) ([]model.Animal, int, error) {
	animals := make([]model.Animal, 0)
	for _, v := range r.animals {
		animals = append(animals, v)
	}
	return animals, len(animals), nil
}
func (r *StubAnimalRepository) SellAnimal(animalId, profileId int, balanceCalc func(int, int) (int, error)) error {
	return nil
}

func TestAnimal(t *testing.T) {
	animalsMap := make(map[int]model.Animal)
	animalsMap[1] = model.Animal{ID: 1, Nickname: "Fluffy"}
	animalsMap[2] = model.Animal{ID: 2, Nickname: "Duffy"}
	animalsMap[3] = model.Animal{ID: 3, Nickname: "Rosy"}
	animalService := service.NewAnimalService(&StubProfileService{}, &StubAnimalRepository{animals: animalsMap})
	animals, err := animalService.FindAllAnimals(1, 1)
	assertNoError(err, t)
	if len(animals.Animals) != 3 {
		t.Errorf("expected %d animals, got %d", 3, len(animals.Animals))
	}
	err = animalService.AddAnimal(&dto.Animal{ID: 10})
	assertNoError(err, t)
	animal, err := animalService.FindAnimalById("10")
	assertNoError(err, t)
	if animal == nil {
		t.Fatalf("Nil animal, but wanted a real one")
	}
	if animal.ID != 10 {
		t.Errorf("expected animal with ID %d, but got %d", 10, animal.ID)
	}
}
