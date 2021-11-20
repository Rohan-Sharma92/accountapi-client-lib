package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	f3 := Init("http://0.0.0.0:8080/")

	assert.IsType(t, &Form3{}, f3)
}
