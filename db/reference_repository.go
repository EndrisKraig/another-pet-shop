package db

import (
	"context"
	"fmt"

	"playground.io/another-pet-store/model"
)

var tables = map[string]string{
	"breed": "breed",
	"type":  "animal_type",
}

type ReferenceRepository interface {
	GetReferences(name string) (*model.References, error)
}

type SimpleReferenceRepository struct {
}

func NewReferenceRepository() ReferenceRepository {
	return &SimpleReferenceRepository{}
}

func (repository SimpleReferenceRepository) GetReferences(name string) (*model.References, error) {
	conn, err := GetConnection()

	if err != nil {
		return nil, err
	}

	tableName, ok := tables[name]

	if !ok {
		return nil, fmt.Errorf("table %v not exists", name)
	}

	query := "SELECT id, label FROM " + tableName
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("error find dictionary with name %v: %w", name, err)
	}

	defer rows.Close()

	var references []model.Reference

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, fmt.Errorf("error during obtaining result rows values: %w", err)
		}
		references = append(references, model.Reference{Id: int(values[0].(int64)), Label: values[1].(string)})
	}
	return &model.References{References: references}, nil

}
