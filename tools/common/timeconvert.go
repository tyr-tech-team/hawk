package common

import "time"

// ParseTimeFromRFC3339 function  
func ParseTimeFromRFC3339(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Time{}
	}
	return t
}
