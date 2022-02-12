package utils

import "strings"

// ParseParams parses param for _Key?-Value pairs and returns a map
// @param param string is the initial string
// @param *map[string]string holds the key, value pairs
// @returns ret string the initial string with all the key, value pairs removed
func ParseParams(param string, mp *map[string]string, transform bool) string {

	m := *mp
	spl := strings.Split(param, "_")
	if len(spl) < 2 {
		return ""
	}
	ret := spl[0]
	spl = spl[1:]
	for _, cf := range spl {
		ss := strings.SplitN(cf, "-", 2)
		if len(ss) < 2 {
			m[ss[0]] = "1"
		} else {
			if transform {
				ss[1] = strings.ToUpper(ss[1])
			}
			m[ss[0]] = ss[1]
		}
	}
	return ret
}
