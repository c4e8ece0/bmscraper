// package smarkets contents implementations of type.Bookmaker
package smarkets

import (
	"encoding/xml"
	"io"

	"golang.org/x/net/html/charset"

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

func ParseXML(r io.Reader) (XMLOdds, error) {
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel
	v := XMLOdds{}
	e := d.Decode(&v)
	return v, e
}

type XMLOdds struct {
	Timestamp string     `xml:"timestamp,attr"`
	Events    []XMLEvent `xml:"event"`
}

type XMLEvent struct {
	Date       string      `xml:"date,attr"`
	Id         string      `xml:"id,attr"`
	Name       string      `xml:"name,attr"`
	Parent     string      `xml:"parent,attr"`
	ParentSlug string      `xml:"parent_slug,attr"`
	Slug       string      `xml:"slug,attr"`
	Time       string      `xml:"time,attr"`
	Type       string      `xml:"type,attr"`
	Url        string      `xml:"url,attr"`
	Markets    []XMLMarket `xml:"market"`
}

type XMLMarket struct {
	Id        string        `xml:"id,attr"`
	Slug      string        `xml:"slug,attr"`
	Contracts []XMLContract `xml:"contract"`
}

type XMLContract struct {
	Id     string     `xml:"id,attr"`
	Name   string     `xml:"name,attr"`
	Slug   string     `xml:"slug,attr"`
	Bids   []XMLPrice `xml:"bids>price"`
	Offers []XMLPrice `xml:"offers>price"`
}

type XMLPrice struct {
	Decimal      string `xml:"decimal,attr"`
	Percent      string `xml:"percent,attr"`
	BackersStake string `xml:"backers_stake,attr"`
	Liability    string `xml:"liability,attr"`
}
