package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Validate_AccountID(t *testing.T) {
	assert.Equal(t, ValidateAccountId("").Error(), "invalid account id")
}

func Test_Validate_AccountID_With_Valid_Value(t *testing.T) {
	assert.Equal(t, ValidateAccountId("test"), nil)
}

func Test_Validate_Version(t *testing.T) {
	assert.Equal(t, ValidateVersion(-1).Error(), "invalid version")
}

func Test_Validate_Version_With_Valid_Value(t *testing.T) {
	assert.Equal(t, ValidateVersion(1), nil)
}
