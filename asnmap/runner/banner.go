package runner

import (
	updateutils "github.com/projectdiscovery/utils/update"
)

const banner = ``

// version is the current version of asnmap
const version = `v1.0.6`

// showBanner is used to show the banner to the user
func showBanner() {}

// GetUpdateCallback returns a callback function that updates asnmap
func GetUpdateCallback() func() {
	return func() {
		showBanner()
		updateutils.GetUpdateToolCallback("asnmap", version)()
	}
}
