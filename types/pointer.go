package types

import "time"

func String(s string) *string {
	return &s
}

func Int(i int) *int {
	return &i
}

func Bool(b bool) *bool {
	return &b
}

func Float(f float64) *float64 {
	return &f
}

func Int64(i int64) *int64 {
	return &i
}

func Time(t time.Time) *time.Time {
	return &t
}
