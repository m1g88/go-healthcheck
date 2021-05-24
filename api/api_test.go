package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDoRequest_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "Bearer 12345", r.Header["Authorization"][0])
		assert.Equal(t, "application/json", r.Header["Accept"][0])
	}))
	defer ts.Close()

	DoRequest(ts.URL, "12345", "Body")
}

func TestDoChannelRequest_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	defer ts.Close()

	ch := make(chan Response)
	go DoChannelRequest(ts.URL, "5", ch)

	res := <-ch

	assert.True(t, res.Success)
}

func TestDoChannelRequest_Timeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(200)
	}))
	defer ts.Close()

	ch := make(chan Response)
	go DoChannelRequest(ts.URL, "1", ch)

	res := <-ch

	assert.False(t, res.Success)
}
