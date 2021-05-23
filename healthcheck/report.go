package healthcheck

import (
	"fmt"
)

type Report struct {
	Total_websites int
	Success        int
	Failure        int
	Total_time     int64
}

func (r *Report) Print() {
	fmt.Printf("Checked webistes: %v\n", r.Total_websites)
	fmt.Printf("Successful websites: %v\n", r.Success)
	fmt.Printf("Failure websites: %v\n", r.Failure)
	fmt.Printf("Total times to finished checking website: %v [time to proceed (ms/sec/minutes]\n", r.Total_time)
}
