package tracker

import (
	"testing"
	"time"
)

func TestNewTracker(t *testing.T) {
	// replace the file
	SetCurrent(Clean())

	tr := New()
	if tr.Lifetime != rotationFallbackLifetime {
		t.Errorf("lifetime error: %v", tr)
	}

	newT, _ := GetCurrent()

	if newT.Timestamp != tr.Timestamp {
		t.Errorf("failed to load from file with correct details")
	}

}

func TestExpiresAt(t *testing.T) {
	now := time.Now().UTC()
	d, _ := time.ParseDuration(rotationFallbackLifetime)
	expected := now.Add(d)
	tr := Track{Timestamp: now, Lifetime: rotationFallbackLifetime}

	if tr.ExpiresAt() != expected {
		t.Errorf("expiry does not matched, expected [%v], actual [%v]", expected, tr.ExpiresAt())
	}

}

func TestValid(t *testing.T) {
	d, _ := time.ParseDuration(rotationFallbackLifetime)
	now := time.Now().UTC().Add(2 * d)

	tr := Track{Timestamp: now, Lifetime: rotationFallbackLifetime}

	if tr.Valid() {
		t.Errorf("should be expired, actual [%v]", tr.Valid())
	}

}
