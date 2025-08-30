package utils

import "time"

func ParseDate(dateLayout string, dateStr string) (time.Time, error) {
	if dateStr != "" {
		parsedDate, err := time.Parse(dateLayout, dateStr)
		return parsedDate, err
	} else {
		return time.Now(), nil
	}
}
