package validator

import (
	"errors"
	"houshold-app/internal/model"
	"net/mail"
	"regexp"
)

func ValidateUser(userModel *model.CreateUser) error {
	validateFieldAreNonNullError := validateFieldAreNonNull(userModel)
	validateHouseNumbersError := validateHouseNumbers(userModel.House)
	validateEmailError := validateEmail(userModel.Email)
	validateMobilePhoneNumberError := validateMobileNumber(userModel.Mobile)

	if validateFieldAreNonNullError != nil {
		return validateFieldAreNonNullError
	}
	if validateHouseNumbersError != nil {
		return validateHouseNumbersError
	}
	if validateEmailError != nil {
		return validateEmailError
	}
	if validateMobilePhoneNumberError != nil {
		return validateMobilePhoneNumberError
	}

	return nil

}

func ValidateUpdatedUser(updateUserModel *model.UpdatedUser) error {
	if updateUserModel.Name != nil && *updateUserModel.Name == "" {
		return errors.New("Name field cannot be empty!")
	}
	if updateUserModel.Email != nil {
		validateEmailError := validateEmail(*updateUserModel.Email)
		if validateEmailError != nil {
			return validateEmailError
		}
	}
	if updateUserModel.Mobile != nil {
		validateMobilePhoneNumberError := validateMobileNumber(*updateUserModel.Mobile)
		if validateMobilePhoneNumberError != nil {
			return validateMobilePhoneNumberError
		}
	}
	if updateUserModel.House != nil {
		validateHouseNumbersError := validateHouseNumbers(*updateUserModel.House)
		if validateHouseNumbersError != nil {
			return validateHouseNumbersError
		}
	}

	return nil

}

func validateFieldAreNonNull(userModel *model.CreateUser) error {
	if userModel.Name == "" || userModel.Email == "" || userModel.Mobile == "" || userModel.House == "" {
		return errors.New("Fields cannot be empty!")
	}
	return nil
}

func validateHouseNumbers(houseNumber model.House) error {
	switch houseNumber {
	case model.L1:
		return nil
	case model.L2:
		return nil
	default:
		return errors.New("Unsupported value! House number can be 'L1' or 'L2'")
	}

}

func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}
	return nil
}

func validateMobileNumber(mobileNumber string) error {
	phoneRegex := regexp.MustCompile(`^\+[1-9]\d{7,14}$`)
	if !phoneRegex.MatchString(mobileNumber) {
		return errors.New("Mobile phone number format must follow the E.164 format (+ and max 15 char)")
	}
	return nil
}
