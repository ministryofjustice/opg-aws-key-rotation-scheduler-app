package opgapp

import (
	"fmt"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"github.com/k0kubun/pp"
)

// UpdateMenu is called frequently via a go routine to check the status
// of access key, trigger gui updates and funcs to update the key
func UpdateMenu(
	info *fyne.MenuItem,
	menu *fyne.Menu,
	key *AccessKeyTracker,
	s *Settings,
	mu *sync.Mutex,
) *AccessKeyTracker {

	now := time.Now().UTC()
	at := key.RotateAt(s.RotationFrequency)

	pp.Printf("[%s] next rotation at [%s]\n", now, at)

	mu.Lock()
	// if the lock file is old, remove it and trigger process
	if key.Locked() && key.LockIsOld() {
		MenuKeyOldLock(info, menu, key, s)
	}

	if key.Locked() {
		MenuKeyLocked(info, menu, key, s)
	} else if now.After(at) {
		key = MenuRotate(info, menu, key, s)
	} else {
		at = key.RotateAt(s.RotationFrequency)
		info.Label = fmt.Sprintf(s.Labels.NextRotation, at.Format(s.DateTimeFormat))

	}
	menu.Refresh()

	mu.Unlock()
	return key
}

// MenuKeyOldLock deals with a lock file that is aged, so presume
// that the rotate failed or was cancelled and therefore
// cleanup the file and carry on
func MenuKeyOldLock(
	info *fyne.MenuItem,
	menu *fyne.Menu,
	key *AccessKeyTracker,
	s *Settings,
) {
	pp.Println("Key is locked and too old, so removing...")
	key.Unlock()
}

// MenuKeyLocked updates gui to show that there is a
// lock file present on the filesystem and there presume
// a rotate is in progress
func MenuKeyLocked(
	info *fyne.MenuItem,
	menu *fyne.Menu,
	key *AccessKeyTracker,
	s *Settings,
) {
	pp.Println("Key is locked...")
	info.Label = s.Labels.Locked
	// add icon change - LOCKED
	menu.Refresh()
}

// MenuRotate handles the gui changes and func calls to change
// a key and show the status of that change
func MenuRotate(
	info *fyne.MenuItem,
	menu *fyne.Menu,
	key *AccessKeyTracker,
	s *Settings,
) *AccessKeyTracker {

	pp.Println("Rotating key...")
	key.Lock()
	// add icon change - UPDATING
	info.Label = s.Labels.Rotating
	menu.Refresh()

	RotateCommand(s)

	key.Unlock()
	key = key.Rotate()
	// add icon change - NORMAL
	at := key.RotateAt(s.RotationFrequency)
	info.Label = fmt.Sprintf(s.Labels.NextRotation, at.Format(s.DateTimeFormat))
	return key
}
