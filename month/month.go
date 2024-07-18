package month

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"
)

func parseMonthString(input string) int {
	if !slices.Contains([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "01", "02", "03", "04", "05", "06", "07", "08", "09"}, input) {
		return 0
	}
	if input[0] == '0' {
		input = input[1:]
	}
	m, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}
	return m
}

// GetMonthInput is a pretty dumb way of retrieving a month input, that only
// accepts a number between 1 and 12
func GetMonthInput() int {

	s := bufio.NewScanner(os.Stdin)
	var month int = 0

	fmt.Print("please select billing month (e.g. 6 for june) > ")
	for month == 0 {
		s.Scan()
		if m := parseMonthString(s.Text()); m == 0 {
			fmt.Print("invalid month, please try again (e.g. 6 for june) > ")
		} else {
			month = m
		}
	}
	return month
}

func beginningOfMonth(month time.Month) time.Time {
	n := time.Now()
	d := time.Date(2024, month, 1, 0, 0, 0, 0, n.Location())
	return d
}

func endOfMonth(month time.Month) time.Time {
	return beginningOfMonth(month).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

func FormatBeginningOfMonth(month time.Month) string {
	d := beginningOfMonth(month)
	return d.Format(time.DateOnly)
}

func FormatEndOfMonth(month time.Month) string {
	d := endOfMonth(month)
	return d.Format(time.DateOnly)
}
