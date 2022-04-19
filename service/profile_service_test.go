package service

import (
	"testing"

	"playground.io/another-pet-store/db"
)

var profileService = NewProfileService(db.NewProfileRepository())

func TestBalanceChange(t *testing.T) {
	balance, err := profileService.ChangeBalance(5, 10)
	want := 5

	if err != nil {
		t.Errorf("Get error, wanted %q : %s", want, err.Error())
	}

	if balance != 5 {
		t.Errorf("Balance is %q, wanted %q", balance, want)
	}

}

func TestNoEnoughBalance(t *testing.T) {
	_, err := profileService.ChangeBalance(15, 10)

	if err == nil {
		t.Errorf("Should be a error with not enough balance")
	}
}
