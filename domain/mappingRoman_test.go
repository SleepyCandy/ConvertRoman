package domain

import (
	"fmt"
	"testing"
)

func TestMappingRoman(t *testing.T) {
	convert := NewConvertRoman()
	stringArray := []string{"M", "C", "C"}
	r, i, err := convert.GroupingRoman(stringArray, "C", "D", "M")
	fmt.Println(r)
	fmt.Println(i)
	fmt.Println(err)
}
