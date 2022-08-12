package search

import (
	"github.com/kennethdyer/versc/logger"
	"io/ioutil"
	"strings"
)

func read(f string) []*Line {
	lines := []*Line{}
	text, err := ioutil.ReadFile(f)
	if err != nil {
		logger.Error("Unable to read file ", f, " due to an error:\n", err)
		return lines
	}
	strLines := strings.Split(string(text), "\n")
	for num, strLine := range strLines {
		lines = append(lines,
			&Line{
				Path:    f,
				Linenum: num,
				Raw:     strLine,
			})
	}

	return lines
}
