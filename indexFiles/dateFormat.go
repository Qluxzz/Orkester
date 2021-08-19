package indexFiles

import "time"

const DateLayout = "2006-01-02"

func ParseDateToIsoDate(date string) (time.Time, error) {
	parsedDate, err := time.Parse(DateLayout, date)

	if err != nil {
		return time.Time{}, err
	}

	return parsedDate, nil
}
