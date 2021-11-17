package common

import "errors"

func ValidateAccountId(accountID string) error {

	if accountID == "" {
		return errors.New("invalid account id")
	}
	return nil
}

func ValidateVersion(version int) error {

	if version < 0 {
		return errors.New("invalid version")
	}
	return nil
}
