package models

// User maps to the "users" table.
type User struct {
	Base
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"not null;uniqueIndex" json:"email"`
	Password string `gorm:"not null" json:"-"`
	Phone    string `gorm:"type:varchar" json:"phone,omitempty"`
	Role     string `gorm:"not null;default:customer" json:"role"`

	// Relationships
	Bookings []Booking `gorm:"foreignKey:UserID" json:"bookings,omitempty"`
}

func (User) TableName() string {
	return "users"
}
