package opgapp

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
	"time"
)

var accessKeyFileMode fs.FileMode = 0644

type AccessKeyTracker struct {
	Timestamp time.Time               `json:"timestamp"`
	Cfg       *AccessKeyConfiguration `json:"cfg"`
}

// RotateAt returns a timestamp of when this AccessKeyTracker should
// trigger the next rotation command
func (k *AccessKeyTracker) RotateAt(frequency string) time.Time {
	dur, _ := time.ParseDuration(frequency)
	return k.Timestamp.Add(dur)
}

// JsonBytes converts this struct into marshaled []byte
// ready to write to a file
func (k *AccessKeyTracker) JsonBytes() (content []byte) {
	content, _ = json.Marshal(k)
	return
}

// Lock generates the lock file to mark that the key rotation is happening
func (k *AccessKeyTracker) Lock() {
	lockfile := k.Cfg.LockFile()
	ioutil.WriteFile(lockfile, k.JsonBytes(), accessKeyFileMode)
}

// Unlock remove the lock file once the rotation is complete
func (k *AccessKeyTracker) Unlock() error {
	lockfile := k.Cfg.LockFile()
	return os.Remove(lockfile)
}

// Locked checks if the lock file exists
func (k *AccessKeyTracker) Locked() (locked bool) {
	lockfile := k.Cfg.LockFile()
	locked = false
	if _, err := os.Stat(lockfile); err == nil {
		locked = true
	}
	return
}

func (k *AccessKeyTracker) LockIsOld() bool {
	lf := k.Cfg.Lock()
	return (time.Since(lf.Timestamp) > time.Hour)
}

// Rotate updates self to be a new version with own timestamp
func (k *AccessKeyTracker) Rotate() *AccessKeyTracker {
	newK, _ := NewAccessKey(k.Cfg).Save()
	return &newK
}

// Save doesnt use pointer so it is chained (NewAccessKey().Save())
func (k AccessKeyTracker) Save() (AccessKeyTracker, error) {
	err := ioutil.WriteFile(k.Cfg.CurrentFile(), k.JsonBytes(), accessKeyFileMode)
	return k, err
}

// == new and current
func CurrentAccessKey(s *Settings) (key *AccessKeyTracker, err error) {
	key = s.AccessKeys.Current()
	return
}

func NewAccessKey(cfg *AccessKeyConfiguration) AccessKeyTracker {
	return AccessKeyTracker{Timestamp: time.Now().UTC(), Cfg: cfg}
}
