package db

import (
	"context"
	"fmt"

	"playground.io/another-pet-store/model"
)

type StickerRepository interface {
	FindAllStickers() ([]model.Sticker, error)
}

type SimpleStickerRepository struct {
	connection Connection
}

func NewStickerRepository(connection Connection) StickerRepository {
	return &SimpleStickerRepository{connection: connection}
}

func (r *SimpleStickerRepository) FindAllStickers() ([]model.Sticker, error) {
	conn, err := r.connection.GetConnection()

	if err != nil {
		return nil, err
	}

	const query = `SELECT s.id, label, uri
	FROM sticker as s JOIN kit as k ON s.kit_id = k.id;`

	rows, err := conn.Query(context.Background(), query)

	if err != nil {
		return nil, fmt.Errorf("error during find offers command %w", err)
	}

	defer rows.Close()
	var stickers []model.Sticker

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, fmt.Errorf("error during obtaining result rows values for offers: %w", err)
		}

		id := values[0].(int64)
		label := values[1].(string)
		uri := values[2].(string)

		sticker := model.Sticker{ID: int(id), KitName: label, Uri: uri}
		stickers = append(stickers, sticker)
	}
	return stickers, nil
}
