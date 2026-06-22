package models

import (
	"time"

	"github.com/google/uuid"
)

// Booking maps to the "bookings" table.
type Booking struct {
	Base
	BookingCode   string    `gorm:"not null;uniqueIndex" json:"booking_code"`
	UserID        uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	ConcertID     uuid.UUID `gorm:"type:uuid;not null" json:"concert_id"`
	TotalAmount   float64   `gorm:"type:numeric;not null" json:"total_amount"`
	Status        string    `gorm:"not null;default:pending" json:"status"`
	PaymentMethod string    `gorm:"type:varchar" json:"payment_method,omitempty"`
	BookingDate   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"booking_date"`

	// Relationships
	User          *User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Concert       *Concert       `gorm:"foreignKey:ConcertID" json:"concert,omitempty"`
	TicketDetails []TicketDetail `gorm:"foreignKey:BookingID" json:"ticket_details,omitempty"`
}

func (Booking) TableName() string {
	return "bookings"
}
