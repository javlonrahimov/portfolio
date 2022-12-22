package main

import "time"

func yearMonthTime(year int, month time.Month) time.Time {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
}

func monthYearDiff(a, b time.Time) (years, months int) {
	m := a.Month()
	for a.Before(b) {
		a = a.Add(time.Hour * 24)
		m2 := a.Month()
		if m2 != m {
			months++
		}
		m = m2
	}
	years = months / 12
	months = months % 12
	return
}
