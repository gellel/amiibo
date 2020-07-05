package amiibo

import "net/http"

// ENGChartURL is the URL for the Nintendo of America Nintendo Amiibo compatibility chart.
const ENGChartURL string = "https://www.nintendo.com/content/noa/en_US/amiibo/compatibility/jcr:content/root/responsivegrid/compatibility_chart.model.json"

// ENGChart is the unfettered English language Nintendo Amiibo product and game support information.
//
// ENGChart is provided by Nintendo of America.
//
// https://www.nintendo.com/content/noa/en_US/amiibo/compatibility/jcr:content/root/responsivegrid/compatibility_chart.model.json
type ENGChart struct {

	// AmiiboList is a collection of Nintendo Amiibo product information.
	AmiiboList []ENGChart `json:"amiiboList"`

	// GameList is a collection of Nintendo Amiibo supported games.
	GameList []ENGChartGame `json:"gameList"`

	// Items is a collection of additional Nintendo Amiibo product information.
	Items []ENGChartItem `json:"items"`
}

// GetENGChart gets the ENGChart from nintendo.com.
func GetENGChart() (req *http.Request, res *http.Response, v ENGChart, err error) {
	return
}
