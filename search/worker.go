package search

import (
	"github.com/kennethdyer/versc/logger"
	"regexp"
)

// FileWorker initializes a File Worker goroutine.  It is used to search a channel of incoming files
// and passes the lines of that file on to a channel of Lines.
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

// LineWorker initializes a Line Worker goroutine.  It is used to perform Regular Expression matching on
// a channel of incoming Line pointers.  If any of them have matches, it sends the line back to a string
// channel for printing.
func LineWorker(id int, res []*regexp.Regexp, vers []string, in <-chan *Line, out chan<- string) {
	logger.Trace("Starting Line Worker ", id)
	defer logger.Trace("Closing Line Worker ", id)

	for line := range in {
		for _, match := range find(line, res, vers) {
			out <- match
		}
	}
}
