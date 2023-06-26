package service

import (
	"errors"
)

type convertRoman struct {
}

func NewConvertRoman() *convertRoman {
	return &convertRoman{}
}

func (s *convertRoman) GroupingRoman(list []string, remain int, minor, majorOne, majorTwo string) (result []string, remainIndex int, err error) {
	var str string
	filter := []string{majorTwo, majorOne, minor}

	for i := remain; i < len(list); i++ {
		//todo validate filter then next stage
		if !contains(filter, list[i]) {
			return result, i, nil
		}

		//todo minor not over 2 C
		if (list[i] == minor) && (list[i+1] == majorOne || list[i+1] == majorTwo) {
			str = list[i] + list[i+1]
			i++
			result = append(result, str)
			continue
		}

		if (list[i] == majorOne || list[i] == majorTwo) && list[i+1] == minor {
			str = list[i]
			i++
			str = str + list[i]
			if i+1 < len(list) && list[i+1] == minor {
				i++
				str = str + list[i]
				if i+1 < len(list) && minor == list[i+1] {
					i++
					str = str + list[i+1]
					if i+1 < len(list) && list[i] == list[i+1] {
						return result, i, errors.New("valid Roman Numeral")
					}
				} else if i+1 < len(list) {
					return result, i, errors.New("valid Roman Numeral")
				}

			} else {
				result = append(result, list[i-1])
				result = append(result, list[i]+list[i+1])
				i++
				remainIndex = i
				continue
			}
			result = append(result, str)
			remainIndex = i
			continue
		}

		if (list[i] == majorOne || list[i] == majorTwo) && (list[i+1] == majorOne || list[i+1] == majorTwo) {
			result = append(result, list[i])
			i++
			remainIndex = i
			continue
		}

		//todo add if same type and be minor
		if list[i] == minor && list[i+1] == minor {
			str = list[i]
			i++
			str = str + list[i]
			if i+1 < len(list) && list[i+1] == minor {
				i++
				str = str + list[i]
				if i+1 < len(list) && list[i] == list[i+1] {
					return result, i, errors.New("valid Roman Numeral ")
				}

			} else if i+1 < len(list) && (list[i+1] == majorOne || list[i+1] == majorTwo) {
				return result, i, errors.New("valid Roman Numeral ")
			}
			result = append(result, str)
			remainIndex = i
			continue
		}

	}

	return
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
