package healthcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJson_Success(t *testing.T) {
	r := Report{
		Total_websites: 1,
		Success:        1,
		Failure:        0,
		Total_time:     1234,
	}

	jsonStr := r.ToJson()

	assert.Equal(t, `{"total_websites":1,"success":1,"failure":0,"total_time":1234}`, jsonStr)
}
