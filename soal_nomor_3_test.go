package jakpat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaseSearchPages1(t *testing.T) {
	result := SearchPage(3)
	assert.Equal(t, "From first page flip 1x", result)
}

func TestCaseSearchPages2(t *testing.T) {
	result := SearchPage(4)
	assert.Equal(t, "From first page flip 2x", result)
}

func TestCaseSearchPages3(t *testing.T) {
	result := SearchPage(230)
	assert.Equal(t, "From last page flip 0x", result)
}
