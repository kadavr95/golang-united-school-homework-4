package string_sum

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	sum := 0
	totalValues := 0
	currentValue := 0
	digits := 0
	positive := true
	incorrectOperand := ""

	for _, s := range input {
		// Incorrect operand checks (fast exit)
		if incorrectOperand != "" {
			// Add new symbols to the operand value until complete value will be parsed
			if s == '-' || s == '+' || s == ' ' {
				_, conversionError := strconv.Atoi(incorrectOperand)
				return "", fmt.Errorf("an error occurred: %w", conversionError)
			} else {
				incorrectOperand += string(s)
			}
		}
		// Check for wrong operands number (fast exit)
		if totalValues > 2 {
			return "", fmt.Errorf("an error occurred: %w", errorNotTwoOperands)
		}
		if s == ' ' {
			continue
		}
		if s == '-' || s == '+' {
			if digits > 0 {
				// Summation with next value
				if positive {
					sum += currentValue
				} else {
					sum -= currentValue
				}
				totalValues++
				currentValue = 0
				digits = 0
				positive = true
			}
			positive = s == '+'
			continue
		}
		// Read new digit
		n, conversionError := strconv.Atoi(string(s))
		if conversionError == nil {
			currentValue = currentValue*10 + n
			digits++
		} else {
			// Manage wrong symbols in the operand
			incorrectOperand = strconv.Itoa(currentValue) + string(s)
		}

	}
	// Finishing operations when string ends
	if digits > 0 {
		if positive {
			sum += currentValue
		} else {
			sum -= currentValue
		}
		totalValues++
	}
	// Checks for the incorrect operand at the end
	if incorrectOperand != "" {
		_, conversionError := strconv.Atoi(incorrectOperand)
		return "", fmt.Errorf("an error occurred: %w", conversionError)
	}
	// Check for wrong operands number
	if totalValues == 0 {
		return "", fmt.Errorf("an error occurred: %w", errorEmptyInput)
	}
	if totalValues != 2 {
		return "", fmt.Errorf("an error occurred: %w", errorNotTwoOperands)
	}

	return strconv.Itoa(sum), nil
}
