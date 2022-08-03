package indexFiles_test

import (
	"goreact/indexFiles"
	"testing"
)

type TestCase struct {
	Input    string
	Expected *indexFiles.ReleaseDate
}

func TestDateParse(t *testing.T) {
	testCases := []TestCase{
		{"2014", &indexFiles.ReleaseDate{2014, 0, 0, "year"}},
		{"2014-01", &indexFiles.ReleaseDate{2014, 1, 0, "month"}},
		{"2014-01-03", &indexFiles.ReleaseDate{2014, 1, 3, "date"}},
	}

	for _, testCase := range testCases {
		date, err := indexFiles.ParseDate(testCase.Input)

		if err != nil {
			t.Errorf("Expected date to be successfully parsed, but got error %s", err)
		}

		if date.Year != testCase.Expected.Year ||
			date.Month != testCase.Expected.Month ||
			date.Date != testCase.Expected.Date {
			t.Errorf("Parsed date (%d-%d-%d) does not match expected date (%d-%d-%d)",
				date.Year,
				date.Month,
				date.Date,
				testCase.Expected.Year,
				testCase.Expected.Month,
				testCase.Expected.Date,
			)
		}
	}
}
