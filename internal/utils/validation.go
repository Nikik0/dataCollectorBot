package utils

import (
	"github.com/pkg/errors"
	"strconv"
)

func ValidateName(name string) (string, error) {
	//todo regexp to check name
	return name, nil
}

func ValidateSurname(surname string) (string, error) {
	//todo regexp to check name
	return surname, nil
}

func ValidateEmail(email string) (string, error) {
	//todo regexp to check name
	return email, nil
}

func ValidateDate(date string) (string, error) {
	//todo regexp to check name
	return date, nil
}

func ValidateAcceptedTerms(accepted string) (bool, error) {
	conf, err := strconv.ParseBool(accepted)
	if err != nil {
		return false, errors.Wrap(err, "Failed to parse message at current state.")
	}
	if conf == true {
		return true, nil
	} else {
		return false, errors.New("Confirmation was false, can't proceed.")
	}
}

func ValidateConfirmation(confirmation string) (bool, error) {
	conf, err := strconv.ParseBool(confirmation)
	if err != nil {
		return false, errors.Wrap(err, "Failed to parse message at current state.")
	}
	if conf == true {
		return true, nil
	} else {
		return false, errors.New("Confirmation was false, can't proceed.")
	}
}
