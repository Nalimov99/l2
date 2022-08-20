package main

import (
	"testing"
	"time"
)

func TestNTPTime(t *testing.T) {
	ntpTime, err := NTPTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		t.Fatal(err)
	}

	diff := time.Since(ntpTime).Milliseconds()
	if diff > 100 {
		t.Fatal("Разница во времени больше 100мс:", diff)
	}
}
