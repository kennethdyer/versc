package search

// Line represents a line found in a given file.  It is used as an interal structure to store
// the path, raw string, and line number.  In the event that the Raw string matches a 
// version descriptor, the data in this struct is used to generate the return text.
type Line struct {
	Path     string   `json:"path"`
	Raw      string   `json:"raw"`
	//Versions []string `json:"versions"`
	Linenum  int      `json:"line_number"`
}
