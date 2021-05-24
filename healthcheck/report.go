package healthcheck

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Report struct {
	Total_websites int   `json:"total_websites"`
	Success        int   `json:"success"`
	Failure        int   `json:"failure"`
	Total_time     int64 `json:"total_time"`
}

func (r *Report) Print() {
	nanostr := strconv.FormatInt(r.Total_time, 10)
	nano, _ := time.ParseDuration(fmt.Sprintf("%s ns", nanostr))

	fmt.Printf("Checked webistes: %v\n", r.Total_websites)
	fmt.Printf("Successful websites: %v\n", r.Success)
	fmt.Printf("Failure websites: %v\n", r.Failure)
	fmt.Printf("Total times to finished checking website: %v [time to proceed (ms/sec/minutes\n", nano.Milliseconds())
}

func (r *Report) ToJson() string {
	b, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}
