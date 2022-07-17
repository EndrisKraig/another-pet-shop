package service_test

import (
	"fmt"
	"testing"

	"playground.io/another-pet-store/model"
	"playground.io/another-pet-store/service"
)

type StubStickerRepository struct {
	stickers []model.Sticker
}

func (r *StubStickerRepository) FindAllStickers() ([]model.Sticker, error) {
	return r.stickers, nil
}

func TestGetStickers(t *testing.T) {
	standardKit := "standard"
	specialKit := "special"
	stickers := []model.Sticker{createSticker(standardKit), createSticker(standardKit), createSticker(standardKit), createSticker(specialKit), createSticker(specialKit)}
	stickerService := service.NewStickerService(&StubStickerRepository{stickers: stickers})
	resp, err := stickerService.GetAllStickers()
	assertNoError(err, t)
	if resp == nil || resp.Kits == nil {
		t.Errorf("Expected sticker response but didn't get one")
	}
	x := resp.Kits[standardKit]
	standardNum := len(x)
	assertStickersNum(3, standardNum, t)
	s := resp.Kits[specialKit]
	specialNum := len(s)
	assertStickersNum(2, specialNum, t)
}

func assertStickersNum(want, get int, t *testing.T) {
	t.Helper()
	if want != get {
		t.Errorf("Expected %d stickers, but got %d", want, get)
	}
}

var id int

func createSticker(kitName string) model.Sticker {
	id++
	return model.Sticker{ID: id, KitName: kitName, Uri: fmt.Sprintf("http://stickers.io/%d", id)}
}
