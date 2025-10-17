package dbrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/vadim-vep/booking/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int

	statement := `INSERT INTO reservations (first_name, last_name, email, phone, start_date, end_date, 
                          room_id, created_at, updated_at) 
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	err := m.DB.QueryRowContext(ctx, statement,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}
	return newID, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `INSERT INTO room_restrictions (room_id, start_date, end_date, reservation_id, restriction_id,
                               created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.DB.ExecContext(ctx, statement,
		r.RoomID,
		r.StartDate,
		r.EndDate,
		r.ReservationID,
		r.RestrictionID,
		time.Now(), //created_at
		time.Now()) //updated_at
	if err != nil {
		return err
	}
	return nil
}

// SearchAvailabilityByDates returns true if availability exists for roomID and false if not
func (m *postgresDBRepo) SearchAvailabilityByDates(start, end time.Time, roomID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT COUNT(id) 
			  FROM room_restrictions 
			  WHERE room_id=$1 AND
			        $2<end_date AND $3>start_date`

	var numRows int

	row := m.DB.QueryRowContext(ctx, query, roomID, start, end)
	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}
	if numRows == 0 {
		return true, nil
	}
	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of all rooms that are not reserved for a given time period
func (m *postgresDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT r.id, r.room_name
			FROM rooms r
			WHERE r.id NOT IN 
			(SELECT rr.room_id FROM room_restrictions rr WHERE $1<rr.end_date AND $2>rr.start_date)`

	var rooms []models.Room
	rows, err := m.DB.QueryContext(ctx, query, start, end)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room
		err := rows.Scan(&room.ID, &room.RoomName)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)

	}
	fmt.Println("__________________")
	return rooms, nil
}

// GetRoomByID returns a room name by its ID
func (m *postgresDBRepo) GetRoomByID(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var room models.Room
	query := `SELECT r.id, r.room_name, r.created_at, r.updated_at
			FROM rooms r
			WHERE r.id = $1`

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&room.ID,
		&room.RoomName,
		&room.CreatedAt,
		&room.UpdatedAt,
	)
	if err != nil {
		return room, err
	}
	return room, nil
}
