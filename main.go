package RomanCalcurate

import (
	"RomanCalcurate/domain"
	"fmt"
)

func main() {
	convert := domain.NewConvertRoman()
	stringArray := []string{"M", "C", "C"}
	r, i, err := convert.GroupingRoman(stringArray, "C", "D", "M")
	fmt.Println(r)
	fmt.Println(i)
	fmt.Println(err)
}
