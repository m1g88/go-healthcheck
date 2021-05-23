package report

import "testing"

func TestSuccess_Success(t *testing.T) {
	r := Report{}

	r.Success()

	if r.success != 1 {
		t.Errorf("Success func invalid")
	}
}

func TestFailure_Success(t *testing.T) {
	r := Report{}

	r.Failure()

	if r.success != 1 {
		t.Errorf("Failure func invalid")
	}
}

func TestPrint_Success(t *testing.T) {
	r := Report{}

	for i := 0; i < 10; i++ {
		r.Success()
	}
	r.Failure()

	r.Print()
}
