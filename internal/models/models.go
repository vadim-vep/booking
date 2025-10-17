package models

import "time"

// Users is a users model
type User struct {
	FirstName   string
	LastName    string
	ID          int
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Rooms is a rooms model
type Room struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Restrictions is a restrictions model
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Reservation is a reservations model
type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Room
}

// RoomRestriction is a room restrictions model
type RoomRestriction struct {
	ID            int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	StartDate     time.Time
	EndDate       time.Time
	RoomID        int
	Room          Room
	Reservation   Reservation
	Restriction   Restriction
	ReservationID int
	RestrictionID int
}
