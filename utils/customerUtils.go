package utils

import (
	"time"

	"golang.org/x/exp/rand"
)

func GenerateRandomDate() time.Time {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)

	sec := rand.Int63n(end.Unix()-start.Unix()) + start.Unix()
	randomDate := time.Unix(sec, 0)

	return time.Date(randomDate.Year(), randomDate.Month(), randomDate.Day(), 0, 0, 0, 0, time.UTC)
}
