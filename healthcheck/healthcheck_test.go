package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckAll_Success(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// you can also inspect the contents of r (the request) to assert over it
	}))
	defer ts.Close()

	mockServerURL := ts.URL

	h := HealthCheck{
		CheckList: []string{mockServerURL},
	}

	h.CheckAll()

	assert.Equal(t, 1, h.Report.Success)
	assert.Equal(t, 0, h.Report.Failure)
	assert.Equal(t, 1, h.Report.Total_websites)
	assert.NotNil(t, h.Report.Total_time)
}

func TestSendReport_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// you can also inspect the contents of r (the request) to assert over it
		w.WriteHeader(200)
	}))
	defer ts.Close()

	mockServerURL := ts.URL

	h := HealthCheck{
		Config: Config{
			AccessToken: "123456789",
			EndpointApi: mockServerURL,
			Timeout:     "1",
		},
	}

	err := h.SendReport()

	assert.NoError(t, err)
}

func TestSendReport_Failure(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// you can also inspect the contents of r (the request) to assert over it
		w.WriteHeader(400)
	}))
	defer ts.Close()

	mockServerURL := ts.URL

	h := HealthCheck{
		Config: Config{
			AccessToken: "123456789",
			EndpointApi: mockServerURL,
			Timeout:     "1",
		},
	}

	err := h.SendReport()

	assert.Error(t, err)
}
