// Package intertops contents implementations of type.Bookmaker
package intertops

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

func ParseXML(r io.Reader) (XMLOdds, error) {
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel
	v := XMLOdds{}
	e := d.Decode(&v)
	return v, e
}

type XMLLine struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type XMLBet struct {
	Type  string    `xml:"type,attr"`
	Id    string    `xml:"id,attr"`
	Lines []XMLLine `xml:"line"`
}

type XMLMatch struct {
	Id             string   `xml:"id,attr"`
	Name           string   `xml:"Name"`
	Rotation       string   `xml:"Rotation"`
	Date           string   `xml:"Date"`
	HasLiveBetting string   `xml:"hasLiveBetting"`
	Bets           []XMLBet `xml:"Bet"`
}

type XMLCompetition struct {
	Name    string     `xml:"Name"`
	Id      string     `xml:"id,attr"`
	Sprtyp  string     `xml:"sprtyp,attr"`
	Compno  string     `xml:"compno,attr"`
	Matches []XMLMatch `xml:"Match"`
}

type XMLOdds struct {
	Competition []XMLCompetition `xml:"Competition"`
}
