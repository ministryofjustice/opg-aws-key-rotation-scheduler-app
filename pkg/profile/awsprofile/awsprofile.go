package awsprofile

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/pref"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"strings"
)

// these are related to aws cli commands, so unlikely to change
const (
	profileList        string = "configure list-profiles | grep  %s$"
	regionCheckDefault string = "configure get region --profile default"
	regionCheck        string = "configure get region --profile %s"
)

type AwsProfile struct{}

func (p *AwsProfile) Name() (ip string) {
	ip = pref.PREFERENCES.ProfileIdentity.Get()
	debugger.Log("AwsProfile.Name()", debugger.VERBOSE, ip)()
	return
}

// Supported needs to check that its installed, the profile `identity` exists and a region has been configured
// for it or for default in order to work
func (p *AwsProfile) Supported(sh shell.Shell) (installed bool, identityProfileFound bool, regionSet bool) {
	installed, identityProfileFound, regionSet = false, false, false
	command := pref.PREFERENCES.ProfileTool.Get()
	// installed
	path, _, _, err := sh.Search(command, true)
	if err == nil && len(path) > 0 {
		installed = true
	}
	debugger.Log("AwsProfile.Supported()", debugger.VERBOSE, "installed:", installed, "path:", path, "err:", err)()
	// can only check the next steps if aws tool is installed
	if installed {
		pl := fmt.Sprintf(profileList, p.Name())

		cmd := fmt.Sprintf("%s %s", path, pl)
		sOut, _, err := sh.Run([]string{cmd}, false)
		hasProfile := strings.Contains(sOut.String(), p.Name())

		debugger.Log("AwsProfile.Supported()", debugger.VERBOSE, "profile found:", hasProfile, "sOut:", sOut.String(), "err:", err)()
		if err == nil && hasProfile {
			identityProfileFound = true
		}

		// now look for the region
		if identityProfileFound {
			hasRegion := false
			rCheck := fmt.Sprintf(regionCheck, p.Name()) //
			// check the profile and default
			for _, c := range []string{regionCheckDefault, rCheck} {
				cmd = fmt.Sprintf("%s %s", path, c)
				sOut, _, err = sh.Run([]string{cmd}, false)
				hasRegion = strings.Contains(sOut.String(), "-")
				debugger.Log("AwsProfile.Supported()", debugger.VERBOSE, "regioncheck:", c, "hasRegion:", hasRegion)()
				if err == nil && hasRegion {
					regionSet = true
					break
				}
			}
			debugger.Log("AwsProfile.Supported()", debugger.VERBOSE, "region found:", hasRegion)()
		}

	}
	debugger.Log("AwsProfile.Supported()", debugger.VERBOSE, "installed", installed, "identityProfileFound", identityProfileFound, "regionSet", regionSet)()
	return

}
