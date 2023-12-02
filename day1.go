package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func findFirstAndLastSingleDigits(s string) (int, bool) {
	// Captures any digit with a capturing group (the parenthesis is the capturing group)
	firstDigitPattern := `(\d)`
	// Capture a capturing group containing a digit and which is followed by any number of non-digits until the end of the line
	lastDigitPattern := `(\d)[^\d]*$`

	firstDigitRegex := regexp.MustCompile(firstDigitPattern)
	lastDigitRegex := regexp.MustCompile(lastDigitPattern)

	firstDigitMatches := firstDigitRegex.FindStringSubmatch(s)
	lastDigitMatches := lastDigitRegex.FindStringSubmatch(s)

	if len(firstDigitMatches) == 0 {
		return 0, false // No digits found
	}

	firstDigit, _ := strconv.Atoi(firstDigitMatches[1])
	// We duplicate the firstDigit as the lastDigit in case we don't find a last digit
	lastDigit := firstDigit

	// But if we do find one, we set it to lastDigit
	if len(lastDigitMatches) > 0 {
		lastDigit, _ = strconv.Atoi(lastDigitMatches[1])
	}

	// Then we combine these digits to from a two-digit number
	combinedNumberStr := fmt.Sprintf("%d%d", firstDigit, lastDigit)
	combinedNumber, err := strconv.Atoi(combinedNumberStr)
	if err != nil {
		return 0, false
	}

	return combinedNumber, true
}

func main() {
	// open file
	f, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Close the file!
	defer f.Close()

	// Read line-by-line using Scanner
	scanner := bufio.NewScanner(f)

	var sum_total int = 0
	for scanner.Scan() {

		line := scanner.Text()
		number, err := findFirstAndLastSingleDigits(line)

		if err != true {
			continue
		}
		sum_total = sum_total + number
	}
	fmt.Println(sum_total)
}
