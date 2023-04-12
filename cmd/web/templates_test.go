package main

import (
	"testing"
	"time"
)

func TestNormalizeDate(t *testing.T) {
	tm := time.Date(2023, 4, 12, 15, 0, 0, 0, time.UTC)
	normalized := NormalizeDate(tm)

	if normalized != "12 Apr 2023 at 15:00" {
		t.Errorf("got %v; want %v", normalized, "12 Apr 2023 at 15:00")
	}
}
