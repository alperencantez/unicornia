// utils.go
//
// Author: Alperen Cantez <alperen.cantez@outlook.com>
// Created on: 23-10-2023

package unicornia

import (
	"strconv"
	"strings"
)

// Checks if given string is a valid credit card, returns the result & card issuer if <input> is a valid card number.
//
// DISCLAIMER: The results may vary by region & bank.
//
// See Luhn algorithm to understand the implementation: https://en.wikipedia.org/wiki/Luhn_algorithm
func UtilsIsCreditCard(input string) (isValid bool, issuer string) {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) < 13 || len(input) > 19 {
		return false, ""
	}

	var sum int
	double := false

	for i := len(input) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(input[i]))
		if err != nil {
			panic(err)
		}

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	mod := (10 - (sum % 10)) % 10
	var issuedBy string

	if string(input[0]) == "2" || string(input[0]) == "5" {
		issuedBy = "mastercard"
	} else if string(input[0]) == "4" {
		issuedBy = "visa"
	} else if string(input[0]) == "3" {
		issuedBy = "amex"
	} else {
		issuedBy = "unknown"
	}

	return mod == 0, issuedBy
}
