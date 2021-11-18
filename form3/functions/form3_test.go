package functions

import (
	"testing"

	"github.com/rohan-sharma92/accountapi-client-lib/form3/core"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	f3 := Init("http://0.0.0.0:8080/")

	assert.IsType(t, &Form3{}, f3)
}

func TestInitWithClient(t *testing.T) {
	c := core.CreateClient("http://0.0.0.0:8080/")
	f3 := InitWithClient(*c)

	assert.IsType(t, &Form3{}, f3)
	assert.IsType(t, "http://0.0.0.0:8080/", f3.Client.HostURL)
}
