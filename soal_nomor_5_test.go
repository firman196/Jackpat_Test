package jakpat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaseEncription1(t *testing.T) {
	result := Encription("aku wibu dan aku bangga")
	assert.Equal(t, "aban kukg udug waba ina", result)
}

func TestCaseEncription2(t *testing.T) {
	result := Encription("have a nice day")
	assert.Equal(t, "hae and via ecy", result)
}
