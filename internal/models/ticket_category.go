package models

import "github.com/google/uuid"

// TicketCategory maps to the "ticket_categories" table.
type TicketCategory struct {
	Base
	ConcertID      uuid.UUID `gorm:"type:uuid;not null" json:"concert_id"`
	Name          string    `gorm:"not null" json:"name"`
	Price         float64   `gorm:"type:numeric;not null" json:"price"`
	TotalQuota    int       `gorm:"not null" json:"total_quota"`
	AvailableQuota int      `gorm:"not null" json:"available_quota"`

	// Relationships
	Concert      *Concert       `gorm:"foreignKey:ConcertID" json:"concert,omitempty"`
	TicketDetails []TicketDetail `gorm:"foreignKey:TicketCategoryID" json:"ticket_details,omitempty"`
}

func (TicketCategory) TableName() string {
	return "ticket_categories"
}
