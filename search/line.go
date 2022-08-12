package search

type Line struct {
    Path string `json:"path"`
    Raw string `json:"raw"`
    Versions []string `json:"versions"`
    Linenum int `json:"line_number"`
}

