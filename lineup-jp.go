package amiibo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// lineupJPN is the unfettered Nintendo Amiibo lineup information provided by nintendo.co.jp.
//
// lineupJPN contains the product properties related to Nintendo Amiibo products.
//
// lineupJPN is assumed to be in Japanese Hiragana.
type lineupJPN struct {

	// jp is a composed property.
	jp

	// XMLName is the xml node.
	XMLName xml.Name `xml:"items"`

	// Item is a collection of Nintendo Amiibo products containing their product information in Japanese.
	Item []lineupItemJPN `xml:"item"`

	// SeriesItem is a collection of Nintendo Amiibo product auxiliary information.
	SeriesItem []lineupSeriesItemJP `xml:"series_item"`
}

// lineupItemJPN is the unfettered Nintendo Amiibo product lineup information from nintendo.co.jp.
type lineupItemJPN struct {

	// XMLName is the xml node.
	XMLName xml.Name `xml:"item"`

	// BigSize is a integer representation of a boolean.
	//
	// BigSize relates to the scale of the Nintendo Amiibo product.
	BigSize int `xml:"bigsize"`

	// Chart is a integer representation of a boolean.
	//
	// Chart relates to the occurrence of the Nintendo Amiibo product in the chart XML.
	Chart int64 `xml:"chart"`

	// Code is the English ID for the Nintendo Amiibo product from the Japanese CDN.
	Code string `xml:"code"`

	// Date is the YYYYMMDD expression of the Nintendo Amiibo product release date.
	Date string `xml:"date"`

	// DisplayDate is the Japanese Hiragana expression of the Nintedo Amiibo product release date.
	//
	// DisplayDate currently (5/07/2020 (DD-MM-YYYY)) has a typo on the nintendo.co.jp and exists
	// as dispalydate.
	DisplayDate string `xml:"displayDate"`

	// Limited is a integer representation of a boolean.
	//
	// Limited relates to the rareness of the Nintendo Amiibo product.
	Limited int `xml:"limited"`

	// Name is the name of the Nintendo product in Japanese Hiragana.
	Name string `xml:"name"`

	// NameKana is the name of the Nintendo Amiibo product in Japanese Hiragana.
	NameKana string `xml:"nameKana"`

	// New is a integer representation of a boolean.
	//
	// New relates to the newness of the Nintendo Amiibo product.
	New int `xml:"new"`

	// Price is the price of the Nintendo Amiibo product in Japanese Yen.
	Price string `xml:"price"`

	// Priority is the numerical rank of the Nintendo Amiibo product.
	Priority int64 `xml:"priority"`

	// Series is the Japanese Hiragana for the Nintendo product that the Nintendo Amiibo product is affiliated with.
	//
	// Series will need to be translated from Japanese to English.
	Series string `xml:"series"`
}

// lineupSeriesItemJP is the unfettered Nintendo Amiibo auxiliary metadata provided by nintendo.co.jp.
//
// lineupSeriesItemJP is assumed to be in Japanese Hiragana.
type lineupSeriesItemJP struct {

	// XMLName is the xml node.
	XMLName xml.Name `xml:"series_item"`

	// BGColor is the hexidecimal code for the Nintendo Amiibo product.
	BGColor string `xml:"bgcolor"`

	// Color is the hexidecimal code for the Nintendo Amiibo product.
	Color string `xml:"color"`

	// Name is the name of the Nintendo product in Japanese Hiragana.
	Name string `xml:"name"`
}

// getLineupJPN gets the http.Response from nintendo.co.jp.
//
// getLineupJPN returns an error on the following:
//
// http.Request is nil or errors.
//
// http.Response is nil or errors.
//
// http.Response.StatusCode is not http.StatusOK.
func getLineupJPN() (req *http.Request, res *http.Response, err error) {
	const URL = "https://www.nintendo.co.jp/hardware/amiibo/chart/data/lineup.xml"
	req, err = http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return
	}
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf(("http: %d"), res.StatusCode)
	}
	if err != nil {
		return
	}
	return
}

// getLineupJPNXML creates a new lineupJPN from getLineupJPN.
func getLineupJPNXML() (v lineupJPN, err error) {
	var b ([]byte)
	var req *http.Request
	var res *http.Response
	req, res, err = getLineupJPN()
	if err != nil {
		return
	}
	v.Cookies = res.Cookies()
	v.Status = res.Status
	v.StatusCode = res.StatusCode
	v.URL = req.URL
	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = xml.Unmarshal(b, &v)
	if err != nil {
		return
	}
	return
}
