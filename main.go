package main

import (
	"RomanCalcurate/service"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "VII"
	stringarray := strings.Split(input, "")
	fmt.Println("input : \"" + input + "\"")
	strArr, err := service.ConvertRoman(stringarray)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Output :", service.CalculateRoman(strArr))
	fmt.Print("Explanation : ")
	for _, v := range strArr {
		fmt.Print(v + " = " + strconv.Itoa(service.ConvertValues(v)) + " ")
	}
	fmt.Println("")
}
