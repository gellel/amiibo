package amiibo

// amiibo is the normalized amiibo data scraped from a rawAmiibo.
type amiibo struct{}

// amiiboCompatability is the unfettered Nintendo Amiibo compatability information provided by nintendo.com.
//
// amiiboCompatability contains the relationship information between Nintendo Amiibo products
// and the games or applications that can be used within.
type amiiboCompatability struct {
	AuthorMode    bool        `json:"authorMode"`
	AmiiboList    []rawAmiibo `json:"amiiboList"`
	ComponentPath string      `json:"componentPath"`
	GameList      []rawGame   `json:"gameList"`
	Items         []rawItem   `json:"items"`
	Language      string      `json:"language"`
	Mode          string      `json:"mode"`
}

// rawAmiibo is the unfettered Nintendo Amiibo product data provided by nintendo.com.
//
// rawAmiibo describes the abbreviated compatability information for a specific Nintendo Amiibo figurine or card.
//
// rawAmiibo contains varying levels of completeness relative to the release status of the product.
type rawAmiibo struct{}

// rawGame is the unfettered game information related to a Nintendo Amiibo product provided by nintendo.com.
//
// rawGame describes the abbreviated game product information that has known Nintendo Amiibo support.
//
// rawGame contains varying levels of accurracy relative to the release status of Nintendo Amiibo products.
type rawGame struct{}

// rawItem is the unfettered auxilliary information related to a Nintendo Amiibo product provided by nintendo.com.
//
// rawItem describes the additional miscellaneous information that relates to Nintendo game that supports a Nintendo Amiibo product.
//
// rawItem contains varying levels of completeness relative to the release status of Nintendo Amiibo products or game titles.
type rawItem struct{}
