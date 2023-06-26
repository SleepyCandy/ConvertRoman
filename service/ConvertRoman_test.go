package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertRoman(t *testing.T) {
	t.Run("TestSuccessCase1", func(t *testing.T) {
		stringArray := []string{"M", "C", "M", "X", "C", "I", "V"}
		r, err := ConvertRoman(stringArray)
		assert.Equal(t, r[0], "M")
		assert.NoError(t, err)
	})
	t.Run("TestSuccessCase2", func(t *testing.T) {
		stringArray := []string{"C", "M"}
		r, err := ConvertRoman(stringArray)
		assert.Equal(t, r[0], "CM")
		assert.NoError(t, err)
	})
	t.Run("TestSuccessCase3", func(t *testing.T) {
		stringArray := []string{"I", "I"}
		r, err := ConvertRoman(stringArray)
		assert.Equal(t, r[0], "II")
		assert.NoError(t, err)
	})
	t.Run("TestSuccessCase4", func(t *testing.T) {
		stringArray := []string{"X", "X", "V", "I", "I"}
		r, err := ConvertRoman(stringArray)
		assert.Equal(t, r[0], "XX")
		assert.Equal(t, r[1], "VII")
		assert.NoError(t, err)
	})
	t.Run("TestErrorWhenOver4", func(t *testing.T) {
		stringArray := []string{"I", "I", "I", "I"}
		_, err := ConvertRoman(stringArray)
		assert.Error(t, err)
	})
	t.Run("TestErrorWhenPrefixMinor", func(t *testing.T) {
		stringArray := []string{"C", "C", "M"}
		_, err := ConvertRoman(stringArray)
		assert.Error(t, err)
	})
}
