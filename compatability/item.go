package compatability

// Item is a snapshot of a Nintendo video-game product provided from https://www.nintendo.com/amiibo/compatability/.
// Item contains data provided as-is from Nintendo with a mixture of content describing
// a Nintendo video-game that is compatabile with an Nintendo Amiibo product.
// Items contain less verbose details than the compatability.Game struct
// but contains details not captured in the aforementioned.
type Item struct {
	Description  string `json:"description"`
	LastModified int64  `json:"lastModified"`
	Path         string `json:"path"`
	Title        string `json:"title"`
	URL          string `json:"url"`
}
