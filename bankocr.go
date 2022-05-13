package bankocr

import "strings"

// _  _     _  _  _  _  _
// | _| _||_||_ |_   ||_||_|
// ||_  _|  | _||_|  ||_| _|

const (
	ERROR = "ERR"
	ILLEGAL = "ILL"

c0 = " _ " +
      "| |" +
      "|_|"

c1 = "   " + 
      "  |" +
      "  |"


c2 = " _ " + 
      " _|" +
	  "|_ "

c3 = " _ " + 
      " _|" +
	  " _|"

c4 = "   " + 
      "|_|" + 
	  "  |"

c5 = " _ " + 
      "|_ " + 
	  " _|"

c6 = " _ " + 
      "|_ " + 
	  "|_|"

c7 = " _ " + 
      "  |" + 
	  "  |"

c8 = " _ " + 
      "|_|" + 
	  "|_|"

c9 = " _ " +
     "|_|" + 
     " _|"

)



// Parse returns a numeric string for the input.
func Parse(input string) string {
	// if len(input) != 31 {
	// 	return ILLEGAL
	// }

	stringToNum := map[string]string{
		c0: "0",
		c1: "1",
		c2: "2",
		c3: "3",
		c4: "4",
		c5: "5",
		c6: "6",
		c7: "7",
		c8: "8",
		c9: "9",
	}
	

	lines := strings.Split(input, "\n")
	if len(lines) != 4 {
		return ILLEGAL
	}

	var parsedNums [9]string

	for _, line := range lines[1:] {

		for i := 0; i < 9; i += 1 {
			leftBound := i * 3
			rightBound := leftBound + 3
			parsedNums[i] += line[leftBound:rightBound]
		}
	}

	var result string
	for _, c := range parsedNums{
		result += stringToNum[c]
	}

	return result
}
