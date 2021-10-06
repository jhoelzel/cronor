//package cronrun defines the cronrunner and its methods
package cronrun

import "time"

//inTimeSpan returns if the current time is between two times
func inTimeSpan(start, end time.Time) bool {
	check := time.Now()
	return check.After(start) && check.Before(end)
}

func int32Ptr(i int32) *int32 { return &i }
