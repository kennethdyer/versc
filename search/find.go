package search

import (
	"fmt"
	"regexp"
	"strings"
)

func matchVers(ver string, vers []string) bool {
	for _, match := range vers {
		if strings.Compare(match, ver) == 0 {
			return true
		}
	}
	return false
}

func formatLine(ver string, line *Line) string {
	return fmt.Sprintf(
		"%s\t-\t%s:%d\t%s",
		ver,
		line.Path,
		line.Linenum+1,
		strings.TrimSpace(line.Raw),
	)
}

func find(line *Line, res []*regexp.Regexp, vers []string) []string {
	ret := []string{}

	hit := false
	var v string
	for !hit {
		for _, re := range res {
			patt := re.FindStringSubmatch(line.Raw)
			if len(patt) > 1 {
				v = string(patt[1])
				if !matchVers(v, vers) && !hit {
					ret = append(ret, formatLine(v, line))
					hit = true
				}
			}
		}
		hit = true
	}
	return ret
}
