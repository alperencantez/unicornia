// strings.go
//
// Author: Alperen Cantez <alperen.cantez@outlook.com>
// Created on: 23-10-2023

package unicornia

import (
	"os/exec"
	"regexp"
	"strings"
)

func StringUuid() string {
	uuid, err := exec.Command("uuidgen").Output()

	if err != nil {
		return err.Error()
	}

	str := string(uuid[:])
	return str
}

// Converts "normal text" to "snake_case"
func StringToSnakeCase(value string) string {
	var wordsSnakeCase []string

	words := strings.Split(value, " ")
	for _, element := range words {
		wordsSnakeCase = append(wordsSnakeCase, element)
	}

	return strings.Join(wordsSnakeCase[:], "_")
}

// Converts "normal text" to "kebab-case"
func StringToKebabCase(value string) string {
	var wordsKebabCase []string

	words := strings.Split(value, " ")
	for _, element := range words {
		wordsKebabCase = append(wordsKebabCase, element)
	}

	return strings.Join(wordsKebabCase[:], "-")
}

// Truncates text after the given index.
func StringTruncate(value string, truncationIndex int) string {
	return value[0:truncationIndex] + "..."
}

// Multiplies given string <times> times and returns.
func StringRepeat(value string, times int) string {
	if times < 2 {
		return value
	}

	var res string

	for i := 0; i < times; i++ {
		res += value
	}

	return res
}

// Checks if <value> starts with <substring>.
func StringStartsWith(value string, substring string) bool {
	checker := value[0:len(substring)]

	if checker == substring {
		return true
	}

	return false
}

// Checks if <value> ends with <substring>.
func StringEndsWith(value string, substring string) bool {
	checker := value[len(value)-len(substring):]

	if checker == substring {
		return true
	}

	return false
}

// Checks if given string is alphanumeric.
func StringIsAlphaNum(value string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(value)
}

// sluggifies-a-string-just-like-this
func StringSlugify(value string) string {
	var res string

	valueAsArray := strings.Split(strings.ToLower(value), " ")
	res = strings.Join(valueAsArray, "-")

	return res
}
