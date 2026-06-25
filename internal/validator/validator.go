package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidators(v *validator.Validate) error {
	return v.RegisterValidation("expire_range", expireRangeValidator)
}

// expire_range validates that a date string in YYYY-MM-DD format
// is between 5 and 30 days from today (inclusive).
func expireRangeValidator(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	expiredAt, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return false
	}

	now := time.Now().Truncate(24 * time.Hour)
	days := int(expiredAt.Sub(now).Hours() / 24)

	return days >= 5 && days <= 30
}
