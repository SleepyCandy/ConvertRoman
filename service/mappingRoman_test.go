package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMappingRoman(t *testing.T) {
	convert := NewConvertRoman()
	stringArray := []string{"M", "C", "M", "X", "C", "I", "V"}
	r, _, _ := convert.GroupingRoman(stringArray, 0, "C", "D", "M")
	assert.Equal(t, r[0], "M")
	assert.Equal(t, r[1], "CM")
}
