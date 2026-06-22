package models

import "github.com/google/uuid"

// TicketDetail maps to the "ticket_details" table.
type TicketDetail struct {
	Base
	BookingID        uuid.UUID `gorm:"type:uuid;not null" json:"booking_id"`
	TicketCategoryID uuid.UUID `gorm:"type:uuid;not null" json:"ticket_category_id"`
	TicketCode       string    `gorm:"not null;uniqueIndex" json:"ticket_code"`
	HolderName       string    `gorm:"not null" json:"holder_name"`
	HolderIdentityNo string    `gorm:"not null" json:"holder_identity_no"`
	IsUsed           bool      `gorm:"not null;default:false" json:"is_used"`

	// Relationships
	Booking        *Booking        `gorm:"foreignKey:BookingID" json:"booking,omitempty"`
	TicketCategory *TicketCategory `gorm:"foreignKey:TicketCategoryID" json:"ticket_category,omitempty"`
}

func (TicketDetail) TableName() string {
	return "ticket_details"
}
