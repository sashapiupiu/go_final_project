package repeat

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const DateLayout = "20060102"

func afterNow(a, b time.Time) bool {
	a = time.Date(a.Year(), a.Month(), a.Day(), 0, 0, 0, 0, a.Location())
	b = time.Date(b.Year(), b.Month(), b.Day(), 0, 0, 0, 0, b.Location())

	return a.After(b)
}

func NextDate(now time.Time, date string, repeat string) (string, error) {
	t, err := time.Parse(DateLayout, date)
	if err != nil {
		return "", err
	}

	if repeat == "" {
		return "", errors.New("empty repeat")
	}

	if repeat == "y" {
		for {
			t = t.AddDate(1, 0, 0)

			if afterNow(t, now) {
				break
			}
		}

		return t.Format(DateLayout), nil
	}

	parts := strings.Fields(repeat)

	if len(parts) != 2 {
		return "", errors.New("invalid repeat")
	}

	if parts[0] != "d" {
		return "", errors.New("invalid repeat")
	}

	interval, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", errors.New("invalid repeat")
	}

	if interval < 1 || interval > 400 {
		return "", errors.New("invalid repeat")
	}

	for {
		t = t.AddDate(0, 0, interval)

		if afterNow(t, now) {
			break
		}
	}

	return t.Format(DateLayout), nil
}
