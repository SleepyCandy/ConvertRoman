package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertValues(t *testing.T) {
	t.Run("TestSuccessCase1", func(t *testing.T) {
		ans := ConvertValues("XX")
		assert.Equal(t, 20, ans)
	})
	t.Run("TestSuccessCase2", func(t *testing.T) {
		ans := ConvertValues("VI")
		assert.Equal(t, 6, ans)
	})
	t.Run("TestSuccessCase3", func(t *testing.T) {
		ans := ConvertValues("III")
		assert.Equal(t, 3, ans)
	})
	t.Run("TestSuccessCase4", func(t *testing.T) {
		ans := ConvertValues("IV")
		assert.Equal(t, 4, ans)
	})
	t.Run("TestSuccessCase4", func(t *testing.T) {
		ans := ConvertValues("M")
		assert.Equal(t, 1000, ans)
	})

}

func TestCalRoman(t *testing.T) {
	t.Run("TestSuccessCase1", func(t *testing.T) {
		stringArray := []string{"M", "CM", "XC", "IV"}
		ans := CalculateRoman(stringArray)
		assert.Equal(t, 1994, ans)
	})
	t.Run("TestSuccessCase2", func(t *testing.T) {
		stringArray := []string{"L", "V", "III"}
		ans := CalculateRoman(stringArray)
		assert.Equal(t, 58, ans)
	})
	t.Run("TestSuccessCase3", func(t *testing.T) {
		stringArray := []string{"III"}
		ans := CalculateRoman(stringArray)
		assert.Equal(t, 3, ans)
	})
}
