package awsprofile

import (
	"fmt"
	"opg-aws-key-rotation-scheduler-app/pkg/debugger"
	"opg-aws-key-rotation-scheduler-app/pkg/shell"
	"strings"
)

const (
	command         string = "aws"
	identityProfile string = "identity"
)

// these are related to aws cli commands, so unlikely to change
const (
	profileList        string = "configure list-profiles | grep " + identityProfile + "$"
	regionCheckDefault string = "configure get region --profile default"
	regionCheck        string = "configure get region --profile " + identityProfile
)

type AwsProfile struct{}

func (p *AwsProfile) Name() string {
	debugger.Log("AwsProfile.Name()", debugger.VERBOSE, identityProfile)()
	return identityProfile
}

// Supported needs to check that its installed, the profile `identity` exists and a region has been configured
// for it or for default in order to work
func (p *AwsProfile) Supported(sh shell.Shell) (installed bool, identityProfileFound bool, regionSet bool) {
	installed, identityProfileFound, regionSet = false, false, false
	// installed
	path, _, _, err := sh.Search(command, true)
	if err == nil && len(path) > 0 {
		installed = true
	}
	debugger.Log("AwsProfile.Supported()", debugger.VERBOSE, "installed:", installed, "path:", path, "err:", err)()
	// can only check the next steps if aws tool is installed
	if installed {

		cmd := fmt.Sprintf("%s %s", path, profileList)
		sOut, _, err := sh.Run([]string{cmd}, false)
		hasProfile := strings.Contains(sOut.String(), identityProfile)

		debugger.Log("AwsProfile.Supported()", debugger.VERBOSE, "profile found:", hasProfile, "sOut:", sOut.String(), "err:", err)()
		if err == nil && hasProfile {
			identityProfileFound = true
		}

		// now look for the region
		if identityProfileFound {
			hasRegion := false
			// check the profile and default
			for _, c := range []string{regionCheckDefault, regionCheck} {
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
