package cron

import (
	"time"
)

// Next returns the next time based on the given time and cron schedule
func Next(from time.Time, schedule string) (time.Time, error) {
	expression, err := ParseSchedule(schedule)
	if err != nil {
		return time.Time{}, err
	}

	nextTime := time.Now()

	if from.Month() > time.Month(expression["month"][len(expression["month"])-1]) {
		nextTime = time.Date(nextTime.Year()+1, nextTime.Month(), nextTime.Day(), nextTime.Hour(), nextTime.Minute(), nextTime.Second(), nextTime.Nanosecond(), time.Local)
	}
	for _, month := range expression["month"] {
		nextTime = time.Date(nextTime.Year(), time.Month(month), from.Day(), from.Hour(), from.Minute(), 0, 0, time.Local)
		for _, day := range expression["date"] {
			nextTime = time.Date(nextTime.Year(), time.Month(month), day, from.Hour(), from.Minute(), 0, 0, time.Local)
			for _, hour := range expression["hour"] {
				nextTime = time.Date(nextTime.Year(), time.Month(month), day, hour, from.Minute(), 0, 0, time.Local)
				for _, minute := range expression["minute"] {
					nextTime = time.Date(nextTime.Year(), time.Month(month), day, hour, minute, 0, 0, time.Local)

					if nextTime.After(from) {
						return nextTime, nil
					}
				}
			}
		}
	}

	return time.Time{}, nil
}
