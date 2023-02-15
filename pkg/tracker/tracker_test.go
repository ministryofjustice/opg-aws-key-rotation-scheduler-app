package tracker

import (
	"opg-aws-key-rotation-scheduler-app/pkg/pref"
	"testing"
	"time"
)

func TestNewTracker(t *testing.T) {

	pref.PREFERENCES = pref.New("test-app", "{}")

	// replace the file
	SetCurrent(Clean())

	tr := New()
	if tr.Lifetime != pref.Fbs["rotation_frequency"] {
		t.Errorf("lifetime error: %v", tr)
	}

	newT, _ := GetCurrent()

	if newT.Timestamp != tr.Timestamp {
		t.Errorf("failed to load from file with correct details")
	}

}

func TestExpiresAt(t *testing.T) {
	now := time.Now().UTC()
	d, _ := time.ParseDuration(pref.Fbs["rotation_frequency"])
	expected := now.Add(d)
	tr := Track{Timestamp: now, Lifetime: pref.Fbs["rotation_frequency"]}

	if tr.ExpiresAt() != expected {
		t.Errorf("expiry does not matched, expected [%v], actual [%v]", expected, tr.ExpiresAt())
	}

}

func TestValid(t *testing.T) {
	d, _ := time.ParseDuration(pref.Fbs["rotation_frequency"])
	ts := time.Now().UTC().Add(-2 * d)

	tr := Track{Timestamp: ts, Lifetime: pref.Fbs["rotation_frequency"]}

	if tr.Valid() {
		t.Errorf("should be expired, actual [%v]", tr.Valid())
	}

}
