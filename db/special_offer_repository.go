package db

import (
	"context"
	"fmt"
	"time"

	"playground.io/another-pet-store/model"
)

type SpecialOfferRepository interface {
	FindAllSpecialsOffers() (*model.SpecialOffers, error)
}

type SimpleSpecialOfferRepository struct {
}

func NewSpecialOfferRepository() SpecialOfferRepository {
	return &SimpleSpecialOfferRepository{}
}

func (repository *SimpleSpecialOfferRepository) FindAllSpecialsOffers() (*model.SpecialOffers, error) {
	conn, err := GetConnection()

	if err != nil {
		return nil, err
	}

	const query = `SELECT sp.id, nickname, label, image_url, price, begin_date, end_date, conditions
	FROM special_offer AS sp JOIN animal AS a on sp.animal_id=a.id JOIN breed AS b on b.id = a.breed_id
	WHERE begin_date <= CURRENT_TIMESTAMP AND end_date > CURRENT_TIMESTAMP;`

	rows, err := conn.Query(context.Background(), query)

	if err != nil {
		return nil, fmt.Errorf("error during find offers command %w", err)
	}

	defer rows.Close()
	var specialOffer []model.SpecialOffer

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, fmt.Errorf("error during obtaining result rows values for offers: %w", err)
		}

		id := values[0].(int64)
		nickname := values[1].(string)
		breed := values[2].(string)
		imageUrl := values[3].(string)
		price := values[4].(int32)
		beginDate := values[5].(time.Time)
		endDate := values[6].(time.Time)
		conditions := values[7].(string)

		offer := model.SpecialOffer{ID: int(id), Nickname: nickname, Breed: breed, Price: int(price), ImageUrl: imageUrl, Conditions: conditions, EndDate: endDate, BeginDate: beginDate}
		specialOffer = append(specialOffer, offer)
	}

	return &model.SpecialOffers{Offers: specialOffer}, nil
}
