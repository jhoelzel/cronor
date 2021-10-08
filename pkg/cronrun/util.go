//package cronrun defines the cronrunner and its methods
package cronrun

import "time"

//inTimeSpan returns if the current time is between two times
func inTimeSpan(start, end time.Time, now time.Time) bool {
	return now.After(start) && now.Before(end)
}

func int32Ptr(i int32) *int32 { return &i }
