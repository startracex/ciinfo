package ciinfo

import (
	"os"
	"strings"
	"sync"

	"github.com/startracex/ciinfo/vendors"
)

type Info struct {
	IsPR    bool
	IsCI    bool
	ID      string
	Name    string
	Vendors map[string]bool
}

func EnvironMap(env []string) map[string]string {
	out := make(map[string]string, len(env))
	for _, envString := range env {
		if key, value, ok := strings.Cut(envString, "="); ok {
			out[key] = value
		}
	}
	return out
}

var GetInfo = sync.OnceValue(
	func() Info {
		return GetInfoFrom(EnvironMap(os.Environ()), vendors.All)
	},
)

func GetInfoFrom(env map[string]string, vendors []vendors.Vendor) Info {
	if isExplicitlyFalseLike(env["CI"]) {
		return Info{}
	}

	info := Info{
		Vendors: make(map[string]bool, 2),
	}

	for _, vendor := range vendors {
		if !vendor.Env.Match(env) {
			continue
		}

		info.Vendors[vendor.Constant] = true
		info.IsCI = true
		info.Name = vendor.Name
		info.ID = vendor.Constant

		if vendor.PR != nil {
			info.IsPR = vendor.PR.Match(env)
		}
	}

	if !info.IsCI {
		info.IsCI = fromCommonKeys(env)
	}

	return info
}

var commonKeys = []string{
	"BUILD_ID",
	"BUILD_NUMBER",
	"CI",
	"CI_APP_ID",
	"CI_BUILD_ID",
	"CI_BUILD_NUMBER",
	"CI_NAME",
	"CONTINUOUS_INTEGRATION",
	"RUN_ID",
}

func fromCommonKeys(env map[string]string) bool {
	for _, k := range commonKeys {
		if env[k] != "" {
			return true
		}
	}
	return false
}

func isExplicitlyFalseLike(s string) bool {
	return s == "false" || s == "0"
}
