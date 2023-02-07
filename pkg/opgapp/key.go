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
	Timestamp time.Time `json:"timestamp"`
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
func (k *AccessKeyTracker) Lock(lockfile string) {
	ioutil.WriteFile(lockfile, k.JsonBytes(), accessKeyFileMode)
}

// Unlock remove the lock file once the rotation is complete
func (k *AccessKeyTracker) Unlock(lockfile string) error {
	return os.Remove(lockfile)
}

// Locked checks if the lock file exists
func (k *AccessKeyTracker) Locked(lockfile string) (locked bool) {
	locked = false
	if _, err := os.Stat(lockfile); err == nil {
		locked = true
	}
	return
}

func (k *AccessKeyTracker) LockedAt(lockfile string) time.Time {
	key := &AccessKeyTracker{}
	content, _ := ioutil.ReadFile(lockfile)
	json.Unmarshal([]byte(content), &key)
	return key.Timestamp
}

// Rotate updates self to be a new version with own timestamp
func (k *AccessKeyTracker) Rotate(file string) *AccessKeyTracker {
	newK, _ := NewAccessKey().Save(file)
	return &newK
}

// Save doesnt use pointer so it is chained (NewAccessKey().Save())
func (k AccessKeyTracker) Save(file string) (AccessKeyTracker, error) {
	err := ioutil.WriteFile(file, k.JsonBytes(), accessKeyFileMode)
	return k, err
}

// == new and current
func CurrentAccessKey(s *Settings) (key *AccessKeyTracker, err error) {
	key = &AccessKeyTracker{}
	content, err := ioutil.ReadFile(s.AccessKeys.CurrentFile())
	json.Unmarshal([]byte(content), &key)
	return
}

func NewAccessKey() AccessKeyTracker {
	return AccessKeyTracker{Timestamp: time.Now().UTC()}
}
