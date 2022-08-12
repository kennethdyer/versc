package search

import (
	"github.com/kennethdyer/versc/logger"
	"regexp"
)

func FileWorker(id int, in <-chan string, out chan<- *Line) {
	logger.Trace("Starting File Worker ", id)
	defer logger.Trace("Closing File Worker ", id)

	for f := range in {
		lines := read(f)
		for _, line := range lines {
			if line.Raw != "" {
				out <- line
			}
		}
	}
}

func LineWorker(id int, res []*regexp.Regexp, vers []string, in <-chan *Line, out chan<- string) {
	logger.Trace("Starting Line Worker ", id)
	defer logger.Trace("Closing Line Worker ", id)

	for line := range in {
		for _, match := range find(line, res, vers) {
			out <- match
		}
	}
}
