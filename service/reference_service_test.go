package service_test

import (
	"fmt"
	"testing"

	"playground.io/another-pet-store/model"
	"playground.io/another-pet-store/service"
)

type StubReferenceRepository struct {
	references map[string]*model.References
}

func (r *StubReferenceRepository) GetReferences(name string) (*model.References, error) {
	value, ok := r.references[name]
	if !ok {
		return nil, fmt.Errorf("No reference for %s", name)
	}
	return value, nil
}

func TestReference(t *testing.T) {
	refName := "breed"
	testData := make(map[string]*model.References)
	testData[refName] = &model.References{References: []model.Reference{
		{Id: 1, Label: "bengal"},
		{Id: 2, Label: "persian"},
	}}

	referenceService := service.NewReferenceService(&StubReferenceRepository{references: testData})

	t.Run("get breed reference", func(t *testing.T) {
		breedReference, err := referenceService.GetReferences(refName)
		assertNoError(err, t)

		if breedReference == nil || breedReference.References == nil {
			t.Error("Wanted a reference but didn't get one")
		}
		num := len(breedReference.References)
		if num != 2 {
			t.Errorf("Wanted %d references, but get %d", 2, num)
		}
	})

	t.Run("get error on non existing reference", func(t *testing.T) {
		_, err := referenceService.GetReferences("terminator")
		assertError(err, t)
	})

}
