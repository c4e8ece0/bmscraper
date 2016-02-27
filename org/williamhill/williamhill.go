// package williamhill contents implementations of type.Bookmaker
package williamhill

// TODO: perfect export to std format
// Hide public structures

import (
	"encoding/xml"
	"io"

	"golang.org/x/net/html/charset"

	// "github.com/c4e8ece0/bmscraper/org"
	"github.com/c4e8ece0/bmscraper/types"
)

type Unit struct {
	XML
}

func init() {
	// org.Register(types.Bookmaker{new(Unit)})
}

// XML-version of scraper
type XML struct {
	types.Bookmaker
}

// Parse XML from io.Reader
// func (x *XML) Get(r io.Reader) (XMLSports, error) {
// 	return ParseXML(r)
// }

func ParseXML(r io.Reader) (XMLClasses, error) {
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel
	v := XMLClasses{}
	e := d.Decode(&v)
	return v, e
}

type XMLClasses struct {
	Sport []XMLClass `xml:"response>williamhill>class"`
}

type XMLClass struct {
	Id         string    `xml:"id,attr"`
	Name       string    `xml:"name,attr"`
	MaxRepDate string    `xml:"maxRepDate,attr"`
	MaxRepTime string    `xml:"maxRepTime,attr"`
	Types      []XMLType `xml:"type"`
}

type XMLType struct {
	Id             string      `xml:"id,attr"`
	Name           string      `xml:"name,attr"`
	Url            string      `xml:"url,attr"`
	LastUpdateDate string      `xml:"lastUpdateDate,attr"`
	LastUpdateTime string      `xml:"lastUpdateTime,attr"`
	Markets        []XMLMarket `xml:"market"`
}

type XMLMarket struct {
	Id                       string           `xml:"id,attr"`
	Name                     string           `xml:"name,attr"`
	Url                      string           `xml:"url,attr"`
	Date                     string           `xml:"date,attr"`
	Time                     string           `xml:"time,attr"`
	BetTillDate              string           `xml:"betTillDate,attr"`
	BetTillTime              string           `xml:"betTillTime,attr"`
	LastUpdateDate           string           `xml:"lastUpdateDate,attr"`
	LastUpdateTime           string           `xml:"lastUpdateTime,attr"`
	PlaceAvailable           string           `xml:"placeAvailable,attr"`
	ForcastAvailable         string           `xml:"forcastAvailable,attr"`
	TricastAvailable         string           `xml:"tricastAvailable,attr"`
	EachwayAvailable         string           `xml:"eachwayAvailable,attr"`
	CashoutAvailable         string           `xml:"cashoutAvailable,attr"`
	StartingPriceAvailable   string           `xml:"startingPriceAvailable,attr"`
	LivePriceAvailable       string           `xml:"livePriceAvailable,attr"`
	GuarenteedPriceAvailable string           `xml:"guarenteedPriceAvailable,attr"`
	FirstPriceAvailable      string           `xml:"firstPriceAvailable,attr"`
	Participants             []XMLParticipant `xml:"participant"`
}

type XMLParticipant struct {
	Name           string `xml:"name,attr"`
	Id             string `xml:"id,attr"`
	Odds           string `xml:"odds,attr"`
	OddsDecimal    string `xml:"oddsDecimal,attr"`
	LastUpdateDate string `xml:"lastUpdateDate,attr"`
	LastUpdateTime string `xml:"lastUpdateTime,attr"`
	Handicap       string `xml:"handicap,attr"`
}
