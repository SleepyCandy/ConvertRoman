package service

import (
	"strings"
)

func CalculateRoman(list []string) int {
	sum := 0
	var listInt []int
	for _, v := range list {
		listInt = append(listInt, ConvertValues(v))
	}
	for _, valueInt := range listInt {
		sum += valueInt
	}
	return sum
}

func ConvertValues(s string) int {
	sum := 0
	var listInt []int
	strList := strings.Split(s, "")
	for i := 0; i < len(strList)-1; i++ {
		if ConvertValue(strList[i]) < ConvertValue(strList[i+1]) {
			listInt = append(listInt, ConvertValue(strList[i])*-1)
		} else {
			listInt = append(listInt, ConvertValue(strList[i]))
		}
	}
	listInt = append(listInt, ConvertValue(strList[len(strList)-1]))
	for _, valueInt := range listInt {
		sum += valueInt
	}
	return sum
}

func ConvertValue(s string) int {
	var result int
	switch s {
	case "I":
		result = 1
	case "V":
		result = 5
	case "X":
		result = 10
	case "L":
		result = 50
	case "C":
		result = 100
	case "D":
		result = 500
	case "M":
		result = 1000
	}
	return result
}
