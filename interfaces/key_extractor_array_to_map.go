// ConvertArrayToMap generic function
func ConvertArrayToMap[T KeyExtractor[T]](list []T) map[string]T {
	result := make(map[string]T)
	for _, item := range list {
    // polymorphic behavior
		result[item.Key()] = item 
	}
	return result
}

// KeyExtractor interface to extract key
type KeyExtractor[T any] interface {
	Key() string
}


// A sruct from another file could implement KeyExtractor as:
type Attendance struct {
	ID         int64
	StartAt    time.Time
	EndAt      time.Time
	Percentage float32
}

func (a *Attendance) Key() string {
	return fmt.Sprintf("%d-%d", a.ID, a.StartAt.UnixMilli())
}
