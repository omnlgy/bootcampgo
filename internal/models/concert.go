package models

import "time"

// Concert maps to the "concerts" table.
type Concert struct {
	Base
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	Date        time.Time `gorm:"not null" json:"date"`
	Venue       string    `gorm:"not null" json:"venue"`
	Status      string    `gorm:"not null;default:active" json:"status"`

	// Relationships
	Bookings        []Booking        `gorm:"foreignKey:ConcertID" json:"bookings,omitempty"`
	TicketCategories []TicketCategory `gorm:"foreignKey:ConcertID" json:"ticket_categories,omitempty"`
}

func (Concert) TableName() string {
	return "concerts"
}
