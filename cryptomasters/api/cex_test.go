package api_test

import (
	"cryptomasters.com/api"
	"testing"
)

func TestGetRate(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
