package logic

import "testing"

func TestDoBusiness(t *testing.T) {
	expected := "Business Logic Executed"
	result := DoBusiness()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
