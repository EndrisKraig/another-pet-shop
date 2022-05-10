package service

import (
	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/dto"
)

type ReferenceService interface {
	GetReferences(name string) (*dto.References, error)
}

type SimpleReferenceService struct {
	referenceRepository db.ReferenceRepository
}

func NewReferenceService(referenceRepository db.ReferenceRepository) ReferenceService {
	return &SimpleReferenceService{referenceRepository: referenceRepository}
}

func (service *SimpleReferenceService) GetReferences(name string) (*dto.References, error) {
	modelReferences, err := service.referenceRepository.GetReferences(name)
	if err != nil {
		return nil, err
	}
	var references []dto.Reference

	for _, v := range modelReferences.References {
		references = append(references, dto.Reference{Id: v.Id, Label: v.Label})
	}
	return &dto.References{References: references}, nil
}
