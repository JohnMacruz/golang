package assert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	assert.Equal(t, 123, 123, "they should be equal")

	assert.NotEqual(t, 123, 148, "they should not be equal")

}
