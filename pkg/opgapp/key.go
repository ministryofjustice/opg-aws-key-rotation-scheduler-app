package opgapp

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
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
func (k *AccessKeyTracker) RotateAt(frequency string) (t time.Time) {
	dur, _ := time.ParseDuration(frequency)
	t = k.Timestamp.Add(dur)
	defer debugger.Log("AccessKeyTracker.RotateAt()", debugger.VERBOSE, "frequency:", frequency, "nextrotate:", t)()
	return
}

// JsonBytes converts this struct into marshaled []byte
// ready to write to a file
func (k *AccessKeyTracker) JsonBytes() (content []byte) {
	defer debugger.Log("AccessKeyTracker.JsonBytes()", debugger.VERBOSE)()
	content, _ = json.Marshal(k)
	return
}

// Lock generates the lock file to mark that the key rotation is happening
func (k *AccessKeyTracker) Lock() {
	defer debugger.Log("AccessKeyTracker.Lock()", debugger.VERBOSE)()
	lockfile := k.Cfg.LockFile()
	ioutil.WriteFile(lockfile, k.JsonBytes(), accessKeyFileMode)
}

// Unlock remove the lock file once the rotation is complete
func (k *AccessKeyTracker) Unlock() error {
	defer debugger.Log("AccessKeyTracker.Unlock()", debugger.VERBOSE)()
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
	defer debugger.Log("AccessKeyTracker.Locked()", debugger.VERBOSE, "locked:", locked)()
	return
}

func (k *AccessKeyTracker) LockIsOld() (isOld bool) {
	lf := k.Cfg.Lock()
	isOld = (time.Since(lf.Timestamp) > time.Hour)
	defer debugger.Log("AccessKeyTracker.Locked()", debugger.VERBOSE, "old:", isOld, "\nlock:\n", lf)()
	return
}

// Rotate updates self to be a new version with own timestamp
func (k *AccessKeyTracker) Rotate() *AccessKeyTracker {
	newK, _ := NewAccessKey(k.Cfg).Save()
	defer debugger.Log("AccessKeyTracker.Rotate()", debugger.VERBOSE, "newkey:\n", newK)()
	return &newK
}

// Save doesnt use pointer so it is chained (NewAccessKey().Save())
func (k AccessKeyTracker) Save() (AccessKeyTracker, error) {
	err := ioutil.WriteFile(k.Cfg.CurrentFile(), k.JsonBytes(), accessKeyFileMode)
	defer debugger.Log("AccessKeyTracker.Save()", debugger.VERBOSE, "key:\n", k, "\nerr:\n", err)()
	return k, err
}

// == new and current
func CurrentAccessKey(s *Settings) (key *AccessKeyTracker, err error) {
	key = s.AccessKeys.Current()
	defer debugger.Log("CurrentAccessKey()", debugger.VERBOSE, "key:\n", key, "\nerr:\n", err)()
	return
}

func NewAccessKey(cfg *AccessKeyConfiguration) AccessKeyTracker {
	defer debugger.Log("NewAccessKey()", debugger.VERBOSE)()
	return AccessKeyTracker{Timestamp: time.Now().UTC(), Cfg: cfg}
}
