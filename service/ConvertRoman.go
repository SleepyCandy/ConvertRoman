package service

func ConvertRoman(stringArray []string) (result []string, err error) {
	convert := NewConvertRoman()
	r, i, err := convert.GroupingRoman(stringArray, 0, "C", "D", "M")
	if err != nil {
		return
	}
	result = append(result, r...)
	if i < len(stringArray) {
		r, i, err = convert.GroupingRoman(stringArray, i, "X", "L", "C")
		if err != nil {
			return
		}
		result = append(result, r...)
	}
	if i < len(stringArray) {
		r, i, err = convert.GroupingRoman(stringArray, i, "I", "V", "X")
		if err != nil {
			return
		}
		result = append(result, r...)
	}
	return result, err
}
