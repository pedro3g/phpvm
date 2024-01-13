package handlers

import "runtime"

func CheckVersionAvailability(version string) (bool, string) {
	releases := ListAllVersions(false)

	for _, release := range releases {
		if release.Tag == version {
			arch := runtime.GOARCH

			if matchArch, ok := release.Source[arch]; ok {
				return true, matchArch
			}

			if release.Source["x64"] != "" {
				return true, release.Source["x64"]
			} else if release.Source["x86"] != "" {
				return true, release.Source["x86"]
			}

			return true, ""
		}
	}

	return false, ""
}
