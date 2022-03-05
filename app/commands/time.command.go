package commands

import (
	"fmt"
	"time"
)

func getCurrentTimeByTimezone(timezone string) (time.Time, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}

	return time.Now().In(loc), nil
}

func GetCurrentTimeByTimezone12hr(timezone string) (string, error) {
	currentTime, err := getCurrentTimeByTimezone(timezone); if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", currentTime.Format("3:04pm")), nil
}

func GetCurrentTimeByTimezone24hr(timezone string) (string, error) {
	currentTime, err := getCurrentTimeByTimezone(timezone); if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", currentTime.Format("15:04")), nil
}
