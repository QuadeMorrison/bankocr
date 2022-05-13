package bankocr

import "strings"

// _  _     _  _  _  _  _
// | _| _||_||_ |_   ||_||_|
// ||_  _|  | _||_|  ||_| _|

const (
	ERROR = "ERR"
	ILLEGAL = "ILL"
	ZERO =
		" _ " +
		"| |" +
		"|_|"
)


// Parse returns a numeric string for the input.
func Parse(input string) string {
	if len(input) != 27 {
		return ILLEGAL
	}

	stringToNum := map[string]string{
		ZERO: "0",
	}
	var result string

	lines := strings.Split(input, "\n")
	if len(lines) != 4 {
		return ILLEGAL
	}

	var parsedNums [9]string

	for line := range lines[:3] {
		for i := 0; i < len(line); i += 3 {

		}
	}

	return input
}
