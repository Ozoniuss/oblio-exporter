package month

import (
	"fmt"
	"testing"
	"time"
)

func TestFormatMonth(t *testing.T) {
	type testcase struct {
		month             int
		expectedBeginning string
		expectedEnd       string
	}
	testcases := []testcase{
		{
			month:             1,
			expectedBeginning: "2024-01-01",
			expectedEnd:       "2024-01-31",
		},
		{
			month:             2,
			expectedBeginning: "2024-02-01",
			expectedEnd:       "2024-02-29",
		},
		{
			month:             3,
			expectedBeginning: "2024-03-01",
			expectedEnd:       "2024-03-31",
		},
		{
			month:             4,
			expectedBeginning: "2024-04-01",
			expectedEnd:       "2024-04-30",
		},
		{
			month:             5,
			expectedBeginning: "2024-05-01",
			expectedEnd:       "2024-05-31",
		},
		{
			month:             6,
			expectedBeginning: "2024-06-01",
			expectedEnd:       "2024-06-30",
		},
		{
			month:             7,
			expectedBeginning: "2024-07-01",
			expectedEnd:       "2024-07-31",
		},
		{
			month:             8,
			expectedBeginning: "2024-08-01",
			expectedEnd:       "2024-08-31",
		},
		{
			month:             9,
			expectedBeginning: "2024-09-01",
			expectedEnd:       "2024-09-30",
		},
		{
			month:             10,
			expectedBeginning: "2024-10-01",
			expectedEnd:       "2024-10-31",
		},
		{
			month:             11,
			expectedBeginning: "2024-11-01",
			expectedEnd:       "2024-11-30",
		},
		{
			month:             12,
			expectedBeginning: "2024-12-01",
			expectedEnd:       "2024-12-31",
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("formatting month %v", tc.month), func(t *testing.T) {
			b := FormatBeginningOfMonth(time.Month(tc.month))
			e := FormatEndOfMonth(time.Month(tc.month))
			if b != tc.expectedBeginning {
				t.Errorf("expected %s, got %s for beginning of month\n", tc.expectedBeginning, b)
			}
			if e != tc.expectedEnd {
				t.Errorf("expected %s, got %s for end of month\n", tc.expectedEnd, e)
			}
		})
	}
}
