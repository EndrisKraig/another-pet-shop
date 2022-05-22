package service

import (
	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/dto"
)

type StickerService interface {
	GetAllStickers() (*dto.Kits, error)
}

type SimpleStickerService struct {
	stickerRepository db.StickerRepository
}

func NewStickerService(stickerRepository db.StickerRepository) StickerService {
	return &SimpleStickerService{stickerRepository: stickerRepository}
}

func (s *SimpleStickerService) GetAllStickers() (*dto.Kits, error) {
	var repository = s.stickerRepository
	var stickers, err = repository.FindAllStickers()
	if err != nil {
		return nil, err
	}
	kits := make(map[string][]dto.Sticker)
	for _, a := range stickers {
		stickers, ok := kits[a.KitName]
		if !ok {
			stickers = make([]dto.Sticker, 0)
		}
		stickers = append(stickers, dto.Sticker{ID: a.ID, Uri: a.Uri})
		kits[a.KitName] = stickers
	}
	return &dto.Kits{Kits: kits}, nil
}
