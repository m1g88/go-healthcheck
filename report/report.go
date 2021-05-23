package report

import (
	"fmt"
)

type Report struct {
	total_websites int
	success        int
	failure        int
	total_time     int64
}

func (r *Report) Success() {
	r.success++
}

func (r *Report) Failure() {
	r.failure++
}

func (r *Report) Print() {
	fmt.Printf("Checked webistes: %v\n", r.total_websites)
	fmt.Printf("Successful websites: %v\n", r.success)
	fmt.Printf("Failure websites: %v\n", r.failure)
	fmt.Printf("Total times to finished checking website: %v [time to proceed (ms/sec/minutes]\n", r.total_time)
}
