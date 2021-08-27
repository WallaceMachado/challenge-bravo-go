package validations

import (
	"challeng-bravo/src/models"
	"errors"
	"strings"
)

func ValidateCurrency(currency *models.Currency) error {
	if currency.Code == "" {
		return errors.New("The code is mandatory and cannot be blank.")
	}

	if len(currency.Code) > 10 || len(currency.Code) < 3 {
		return errors.New("The code must be between 3 and 10 characters.")
	}

	if currency.Name == "" {
		return errors.New("The name is mandatory and cannot be blank.")
	}

	if len(currency.Name) > 50 || len(currency.Name) < 3 {
		return errors.New("The name must be between 3 and 50 characters.")
	}

	if currency.ValueInUSD <= 0 {
		return errors.New("the ValueInUSD must be greater than zero.")
	}

	currency.Code = strings.ToUpper(strings.TrimSpace(currency.Code))
	currency.Name = strings.ToUpper(strings.TrimSpace(currency.Name))

	return nil
}

func ValidateConversionCurrency(to string, from string, amount float64) error {
	if to == "" {
		return errors.New("The 'to' currency is mandatory and cannot be blank.")
	}

	if len(to) > 10 || len(to) < 3 {
		return errors.New("The 'to' currency must be between 3 and 10 characters.")
	}
	if from == "" {
		return errors.New("The 'from' currency is mandatory and cannot be blank.")
	}

	if len(from) > 10 || len(from) < 3 {
		return errors.New("The 'from' currency must be between 3 and 10 characters.")
	}

	if amount <= 0 {
		return errors.New("the 'amount' must be greater than zero.")
	}

	return nil
}
