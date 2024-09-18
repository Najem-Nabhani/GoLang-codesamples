const (
	timeLayout = "15:04"
)

type Slot struct {
	ID                      *int64
	StartTime               string
	EndTime                 string
}

// CalculateEndDatetime calculates the end datetime based on the start datetime, knowing the start and end times
func (s Slot) CalculateEndDatetime(startDatetime int64) (int64, error) {

	startTime, err := time.Parse(timeLayout, s.StartTime)
	if err != nil {
		return 0, fmt.Errorf("err parsing start time: %w", err)
	}

	endTime, err := time.Parse(timeLayout, s.EndTime)
	if err != nil {
		return 0, fmt.Errorf("err parsing end time: %w", err)
	}

	if endTime.Before(startTime) {
		// Adding 24 hours to endTime to represent the next day
		endTime = endTime.Add(24 * time.Hour)
	}

	shiftDuration := endTime.Sub(startTime)
	endDatetime := time.UnixMilli(startDatetime).Add(shiftDuration).UnixMilli()

	return endDatetime, nil
}
