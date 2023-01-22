package types

import "time"

func StringValue(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

func IntValue(i *int) int {
	if i == nil {
		return 0
	}

	return *i
}

func BoolValue(b *bool) bool {
	if b == nil {
		return false
	}

	return *b
}

func FloatValue(f *float64) float64 {
	if f == nil {
		return 0
	}

	return *f
}

func Int64Value(i *int64) int64 {
	if i == nil {
		return 0
	}

	return *i
}

func TimeValue(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}

	return *t
}
