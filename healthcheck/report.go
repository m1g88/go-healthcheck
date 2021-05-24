package healthcheck

import (
	"encoding/json"
	"fmt"
)

type Report struct {
	Total_websites int   `json:"total_websites"`
	Success        int   `json:"success"`
	Failure        int   `json:"failure"`
	Total_time     int64 `json:"total_time"`
}

func (r *Report) ToJson() string {
	b, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}
