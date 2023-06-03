package jakpat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	result := ValidatePhoneNumber("81234567890")
	assert.Equal(t, "0812-3456-7890", result)
}

func TestCase2(t *testing.T) {
	result := ValidatePhoneNumber("081-234-567-890")
	assert.Equal(t, "0812-3456-7890", result)
}
