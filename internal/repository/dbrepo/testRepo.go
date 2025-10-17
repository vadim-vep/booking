package dbrepo

import (
	"errors"
	"time"

	"github.com/vadim-vep/booking/internal/models"
)

func (m *testPostgresRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *testPostgresRepo) InsertReservation(res models.Reservation) (int, error) {
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (m *testPostgresRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	return nil
}

// SearchAvailabilityByDates returns true if availability exists for roomID and false if not
func (m *testPostgresRepo) SearchAvailabilityByDates(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of all rooms that are not reserved for a given time period
func (m *testPostgresRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {

	var rooms []models.Room
	return rooms, nil
}

// GetRoomByID returns a room name by its ID
func (m *testPostgresRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("room not found")
	}
	return room, nil
}
