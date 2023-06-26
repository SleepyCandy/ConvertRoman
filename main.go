package main

import (
	"RomanCalcurate/service"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "MCMXCIV"
	stringarray := strings.Split(input, "")
	fmt.Println("input : \"" + input + "\"")
	strArr, err := service.ConvertRoman(stringarray)
	if err != nil {
		panic(err)
	}
	fmt.Println("Output :", service.CalculateRoman(strArr))
	fmt.Print("Explanation : ")
	for _, v := range strArr {
		fmt.Print(v + " = " + strconv.Itoa(service.ConvertValues(v)) + " ")
	}
}
