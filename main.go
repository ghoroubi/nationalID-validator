/*
NationalID validator in golang,
Developed By Nima Ghoroubi,
October - 2019
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nationalCode1 := "2949850685" // valid
	nationalCode2 := "2949850686" // invalid
	isValid := NationalCodeIsValid(nationalCode1)
	fmt.Printf("%v validation is%v ", nationalCode1, isValid)
	isValid = NationalCodeIsValid(nationalCode2)
	fmt.Printf("%v validation is%v ", nationalCode2, isValid)
	
}

// NationalCodeIsValid , checks the validation of national code
func NationalCodeIsValid(code string) bool {
	var (
		digits      = make([]int, 0)
		posDigits   = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
		sum         = 0
		processCtrl int
	)
	
	nums := strings.Split(code, "")
	if len(nums) < 10 {
		return false
	}
	
	for _, str := range nums {
		num, _ := strconv.Atoi(str)
		digits = append(digits, num)
	}
	control := digits[len(digits)-1]   // last Digit
	digits = digits[0 : len(digits)-1] // remove last digit from array
	
	for i, v := range posDigits {
		sum += digits[i] * v
	}
	
	r := sum % 11
	if r < 2 {
		processCtrl = r
	} else {
		processCtrl = 11 - r
	}
	return control == processCtrl
}
