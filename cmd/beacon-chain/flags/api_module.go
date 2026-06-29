package flags

import "strings"

const SilaAPIModule string = "sila"
const SilaCompatAPIModule string = "eth"

func EnableHTTPSilaAPI(httpModules string) bool {
	return enableAPI(httpModules, SilaAPIModule)
}

func EnableHTTPSilaCompatAPI(httpModules string) bool {
	return enableAPI(httpModules, SilaCompatAPIModule)
}

func enableAPI(httpModules, api string) bool {
	for m := range strings.SplitSeq(httpModules, ",") {
		if strings.EqualFold(m, api) {
			return true
		}
	}
	return false
}
