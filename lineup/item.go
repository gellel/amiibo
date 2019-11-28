package lineup

import (
	"path/filepath"
	"strings"
)

// Item is a snapshot of a Nintendo Amiibo product provided from resource.Lineup.
// Item contains data provided as-is from Nintendo with a mixture of content describing
// a Nintendo Amiibo product. Items contain less verbose details than the lineup.Amiibo struct
// but contains details not captured in the aforementioned.
type Item struct {
	Description  string `json:"description"`
	LastModified int64  `json:"lastModified"`
	Path         string `json:"path"`
	Title        string `json:"title"`
	URL          string `json:"url"`
}

// Key returns a reliable ID.
func (i *Item) Key() string {
	return strings.TrimSuffix(filepath.Base(i.URL), ".html")
}
