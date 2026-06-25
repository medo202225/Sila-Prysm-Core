package flags

import "strings"

const SilaAPIModule string = "sila"
const EthAPIModule string = "eth"

func EnableHTTPSilaAPI(httpModules string) bool {
	return enableAPI(httpModules, SilaAPIModule)
}

func EnableHTTPEthAPI(httpModules string) bool {
	return enableAPI(httpModules, EthAPIModule)
}

func enableAPI(httpModules, api string) bool {
	for m := range strings.SplitSeq(httpModules, ",") {
		if strings.EqualFold(m, api) {
			return true
		}
	}
	return false
}
