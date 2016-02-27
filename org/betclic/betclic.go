// package betclic contents implementations of type.Bookmaker
package betclic

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

func ParseXML(r io.Reader) (XMLSports, error) {
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel
	v := XMLSports{}
	e := d.Decode(&v)
	return v, e
}

type XMLSports struct {
	Sport []XMLSport `xml:"sport"`
}

type XMLSport struct {
	Name   string     `xml:"name,attr"`
	Id     int        `xml:"id,attr"`
	Events []XMLEvent `xml:"event"`
}

type XMLEvent struct {
	Name    string     `xml:"name,attr"`
	Id      int        `xml:"id,attr"`
	Matches []XMLMatch `xml:"match"`
}

type XMLMatch struct {
	Name      string   `xml:"name,attr"`
	Id        int      `xml:"id,attr"`
	StartDate string   `xml:"start_date,attr"`
	LiveId    string   `xml:"live_id,attr"`
	Streaming string   `xml:"streaming,attr"`
	Bets      []XMLBet `xml:"bets>bet"`
}

type XMLBet struct {
	Code    string      `xml:"code,attr"`
	Name    string      `xml:"name,attr"`
	Id      int         `xml:"id,attr"`
	Choices []XMLChoice `xml:"choice"`
}

type XMLChoice struct {
	Name string `xml:"name,attr"`
	Id   int    `xml:"id,attr"`
	Odd  string `xml:"odd,attr"`
}
