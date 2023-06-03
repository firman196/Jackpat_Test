package jakpat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaseChat1(t *testing.T) {
	result := SearchChat("Ari")
	assert.Equal(t, []string{"backend", "frontend"}, result)
}

func TestCaseChat2(t *testing.T) {
	result := SearchChat("lorem")
	assert.Equal(t, []string{"frontend"}, result)
}
