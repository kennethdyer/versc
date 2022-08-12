package search

import "strings"

func sortStrings(rets []string) []string {
	lenRets := len(rets)
	var tmp string
	for i := 0; i < lenRets; i++ {
		for j := 0; j < lenRets; j++ {
			if strings.Compare(rets[i], rets[j]) > 0 {
				tmp = rets[i]
				rets[i] = rets[j]
				rets[j] = tmp
			}
		}
	}

	return rets
}
