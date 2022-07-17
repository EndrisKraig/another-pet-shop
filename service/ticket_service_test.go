package service_test

import (
	"testing"

	"playground.io/another-pet-store/service"
)

func TestTicket(t *testing.T) {
	ticketService := service.NewTicketService()

	t.Run("create and read ticket", func(t *testing.T) {
		profileId := 123
		token := ticketService.CreateTicket(profileId)
		profile, err := ticketService.ReadTicket(token)
		if profile != profileId {
			t.Errorf("Wanted %d but got %d", profileId, profile)
		}
		assertNoError(err, t)
	})

	t.Run("error on non existing ticket", func(t *testing.T) {
		_, err := ticketService.ReadTicket("")
		assertError(err, t)
	})

}
