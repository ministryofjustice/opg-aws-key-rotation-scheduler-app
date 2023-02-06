package opgapp

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
	"time"

	"github.com/k0kubun/pp"
)

var accessKeyFileMode fs.FileMode = 0644

type AccessKeyTracker struct {
	Timestamp time.Time `json:"timestamp"`
}

func (k *AccessKeyTracker) RotateAt(frequency string) time.Time {
	dur, _ := time.ParseDuration(frequency)
	return k.Timestamp.Add(dur)
}

func (k *AccessKeyTracker) JsonBytes() (content []byte) {
	content, _ = json.Marshal(k)
	return
}

func (k *AccessKeyTracker) Lock(lockfile string) {
	pp.Println("write file - " + lockfile)
	ioutil.WriteFile(lockfile, k.JsonBytes(), accessKeyFileMode)
}
func (k *AccessKeyTracker) Unlock(lockfile string) error {
	return os.Remove(lockfile)
}

func (k *AccessKeyTracker) Locked(lockfile string) (locked bool) {
	locked = false
	if _, err := os.Stat(lockfile); err == nil {
		locked = true
	}
	return
}

func (k *AccessKeyTracker) Rotate(file string) (key *AccessKeyTracker) {
	newK, _ := NewAccessKey().Save(file)
	key = &newK
	return
}

// Save doesnt use pointer so it is chained (NewAccessKey().Save())
func (k AccessKeyTracker) Save(file string) (AccessKeyTracker, error) {
	err := ioutil.WriteFile(file, k.JsonBytes(), accessKeyFileMode)
	return k, err
}

// == new and current
func CurrentAccessKey(s Settings) (key AccessKeyTracker, err error) {
	key = AccessKeyTracker{}
	content, err := ioutil.ReadFile(s.CurrentAccessKeyFilepath())
	json.Unmarshal([]byte(content), &key)
	return
}

func NewAccessKey() AccessKeyTracker {
	return AccessKeyTracker{Timestamp: time.Now().UTC()}
}
