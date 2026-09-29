package validators

import (
	"houshold-app/internal/model"

	"github.com/go-playground/validator/v10"
)

func IsValidHouseNumber(fl validator.FieldLevel) bool {
	houseNumber := fl.Field().Interface().(model.House)
	switch houseNumber {
	case model.L1:
		return true
	case model.L2:
		return true
	default:
		return false
	}
}
