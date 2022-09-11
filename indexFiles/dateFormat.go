package indexFiles

import (
	"fmt"
	"time"
)

type ReleaseDatePrecision string

const (
	PRECISION_DATE  ReleaseDatePrecision = "date"
	PRECISION_MONTH ReleaseDatePrecision = "month"
	PRECISION_YEAR  ReleaseDatePrecision = "year"
)

type ReleaseDate struct {
	Year      int                  `json:"year"`
	Month     int                  `json:"month,omitempty"`
	Date      int                  `json:"date,omitempty"`
	Precision ReleaseDatePrecision `json:"precision"`
}

func (a *ReleaseDate) After(b *ReleaseDate) bool {
	if a.Year > b.Year {
		return true
	} else if a.Year < b.Year {
		return false
	}

	if a.Month > b.Month {
		return true
	} else if a.Month < b.Month {
		return false
	}

	return a.Date > b.Date
}

const ISODateLayout = "2006-01-02"
const YearLayout = "2006"
const YearAndMonthLayout = "2006-01"

func ParseDate(date string) (*ReleaseDate, error) {
	parsedIsoDate, err := time.Parse(ISODateLayout, date)

	if err == nil {
		return &ReleaseDate{
			Year:      parsedIsoDate.Year(),
			Month:     int(parsedIsoDate.Month()),
			Date:      parsedIsoDate.Day(),
			Precision: PRECISION_YEAR,
		}, nil
	}

	parsedYearMonthDate, err := time.Parse(YearAndMonthLayout, date)

	if err == nil {
		return &ReleaseDate{
			Year:      parsedYearMonthDate.Year(),
			Month:     int(parsedYearMonthDate.Month()),
			Precision: PRECISION_MONTH,
		}, nil
	}

	parsedYearDate, err := time.Parse(YearLayout, date)

	if err == nil {
		return &ReleaseDate{
			Year:      parsedYearDate.Year(),
			Precision: PRECISION_YEAR,
		}, nil
	}

	return nil, fmt.Errorf("failed to parse datestring %s", date)
}
