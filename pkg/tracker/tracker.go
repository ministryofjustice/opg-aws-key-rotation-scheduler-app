package tracker

import (
	"encoding/json"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/opgapp"
	"opg-aws-key-rotation-scheduler-app/pkg/storage"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// == configured by init
var (
	directory   string
	lockfile    string
	currentfile string
	lifetime    string
)

const (
	lockfileName             string = "lock.v1"
	currentFileName          string = "current.v1"
	rotationPreferencesKey   string = "rotation_frequency"
	rotationFallbackLifetime string = "15m"
)

// Track struct holds timestamp and lifetime for a key
// and provides methods to check its validity
type Track struct {
	Timestamp time.Time `json:"timestamp"`
	Lifetime  string    `json:"lifetime"`
}

// ExpiresAt returns the timestamp + lifetime as a time
//   - used for displaying when the rotation will happen
func (tr *Track) ExpiresAt() (t time.Time) {
	dur, _ := time.ParseDuration(tr.Lifetime)
	t = tr.Timestamp.Add(dur)
	defer debugger.Log("Track.ExpiresAt()", debugger.VERBOSE, t)()
	return
}

// Valid checks if the time since the timestamp is greater
// than the lifetime
func (tr *Track) Valid() (v bool) {
	now := time.Now().UTC()
	ex := tr.ExpiresAt()
	v = !now.After(ex)
	defer debugger.Log("Track.Valid()", debugger.INFO, "now:\t\t"+now.String(), "expires:\t"+ex.String(), "valid:\t\t"+strconv.FormatBool(v))()
	return
}

// OlderThan compares the time since the timestamp against
// the duration and determines if its longer ago than that
// gap
func (tr *Track) OlderThan(duration string) (aged bool) {
	dur, _ := time.ParseDuration(duration)
	aged = time.Since(tr.Timestamp) > dur
	defer debugger.Log("Track.OlderThan()", debugger.VERBOSE, aged)()
	return
}
func (tr *Track) Older(duration time.Duration) (aged bool) {
	aged = time.Since(tr.Timestamp) > duration
	defer debugger.Log("Track.Older()", debugger.VERBOSE, aged)()
	return
}

// Json converts the struct into a []byte of json ready to
// write to a file
func (tr *Track) Json() (content []byte) {
	content, _ = json.Marshal(tr)
	return
}

// ===

// GetCurrent tries to read and unmarsal the current
// track file finto a Track struct for use
//   - returns an error from either the readfile or unmarshaling
func GetCurrent() (tr Track, err error) {
	content, err := os.ReadFile(currentfile)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(content), &tr)
	defer debugger.Log("tracker.GetCurrent()", debugger.VERBOSE, tr, "error:", err)()
	return
}

// SetCurrent writes the tr passed into the standard current file
//   - returns err from WriteFile
func SetCurrent(tr Track) (t Track, err error) {
	err = os.WriteFile(currentfile, tr.Json(), storage.StoragePermissionMode)
	t = tr
	defer debugger.Log("tracker.SetCurrent()", debugger.VERBOSE, tr, "error:", err)()
	return
}

// GetLock returns the lock file if it exists as a struct
//   - returns error from either readfile or unmarshal
func GetLock() (tr Track, err error) {
	content, err := os.ReadFile(lockfile)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(content), &tr)
	defer debugger.Log("tracker.GetLock()", debugger.VERBOSE, tr, "error:", err)()
	return
}

// SetLock creates the lock file using the track item passed
func SetLock(tr Track) (err error) {
	err = os.WriteFile(lockfile, tr.Json(), storage.StoragePermissionMode)
	defer debugger.Log("tracker.SetLock()", debugger.VERBOSE, tr, "error:", err)()
	return
}

// Unlock removes lock file
func Unlock() (err error) {
	err = os.Remove(lockfile)
	defer debugger.Log("tracker.Unlock()", debugger.VERBOSE, "error:", err)()
	return
}

func Clean() (tr Track) {
	lifetime = opgapp.Preferences.StringWithFallback(rotationPreferencesKey, rotationFallbackLifetime)
	tr = Track{Timestamp: time.Now().UTC(), Lifetime: lifetime}
	return
}

// New creates a fresh tracker using the current time
// and the lifetime (as a string formatted duration)
func New() (tr Track) {
	tr = Clean()
	if track, err := GetCurrent(); err == nil {
		tr = track
		defer debugger.Log("tracker.New()", debugger.VERBOSE, "Found current tracker.", tr)()
	}
	defer debugger.Log("tracker.New()", debugger.VERBOSE, tr)()
	return
}

// init fetches storage directory details from the `storage` module
// directly
func init() {
	directory = storage.StorageDirectory()
	lockfile = filepath.Join(directory, lockfileName)
	currentfile = filepath.Join(directory, currentFileName)
}
