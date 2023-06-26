package domain

import (
	"errors"
)

type convertRoman struct {
}

func NewConvertRoman() *convertRoman {
	return &convertRoman{}
}

func (s *convertRoman) GroupingRoman(list []string, minor, majorOne, majorTwo string) (result []string, index int, err error) {
	var str string
	filter := []string{minor, majorOne, majorTwo}

	for i := 0; i < len(list); i++ {
		//todo validate filter then next stage
		if !contains(filter, list[i]) {
			return result, i, nil
		}
		//todo minor not over 2
		if i == 2 && (list[i] == minor) {
			result = append(result, list[i])
			return result, i, errors.New("wrong format")
		}
		//todo major not over 3
		if i == 3 && (list[i] == majorOne || list[i] == majorTwo) {
			result = append(result, list[i])
			return result, i, errors.New("wrong format")
		}

		//todo major not next minor
		if (list[i] == majorOne && majorOne == list[i+1]) || (list[i] == majorTwo && majorTwo == list[i+1]) {
			result = append(result, list[i+1])
			return result, i, errors.New("wrong format")
		}
		//todo major not lasted
		if (i == len(list) || i == 0) && (list[i] == majorOne || list[i] == majorTwo) {
			str = list[i]
			if list[i+1] == minor {
				i++
				str = str + list[i]
				if i+1 < len(list) && list[i+1] == minor {
					i++
					str = str + list[i]
				}
				if i+1 < len(list) && minor == list[i+1] {
					str = str + list[i+1]
					i++
					if i+1 < len(list) && list[i] == list[i+1] {
						return result, i, errors.New("over 4 charecter")
					}

				}
				result = append(result, str)
				index = i
				continue
			}
		}

		//todo add if same type and be minor
		if list[i] == minor && list[i] == list[i+1] {
			str = list[i] + list[i+1]
			i++
			if i+1 < len(list) && list[i] == list[i+1] {
				str = str + list[i+1]
				i++
				if i+1 < len(list) && list[i] == list[i+1] {
					return result, i, errors.New("over 4 charecter")
				}

			}
			result = append(result, str)
			index = i
			continue
		}
		//todo add if next minor is major
		if list[i] == minor && (list[i+1] == majorOne || list[i+1] == majorTwo) {
			str = list[i] + list[i+1]
			i++
			result = append(result, str)
			index = i
			continue
		}

		//todo add if next major is minor
		if (list[i] == majorOne || list[i] == majorTwo) && list[i+1] == minor {
			str = list[i] + list[i+1]
			i++
			result = append(result, str)
			index = i
			continue
		}
	}
	return
}

func remove(a []string, index int) []string {
	for i := 0; i < index; i++ {
		a[i] = a[len(a)-1] // Copy last element to index i.
		a[len(a)-1] = ""   // Erase last element (write zero value).
		a = a[:len(a)-1]   // Truncate slice.
	}
	return a
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
