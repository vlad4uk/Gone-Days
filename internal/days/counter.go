package days

import (
	"fmt"
	"time"
)

func CountDaysFrom(dateStr string) (int, error) {
	const layout = "02.01.2006"

	startDate, err := time.Parse(layout, dateStr)
	if err != nil {
		return 0, fmt.Errorf("Не удалось разобрать дату %q: %w", dateStr, err)
	}

	today := time.Now()
	duration := today.Sub(startDate)
	return int(duration.Hours() / 24), nil
}
