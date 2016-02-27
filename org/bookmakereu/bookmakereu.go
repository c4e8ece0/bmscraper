// package bookmakereu contents implementations of type.Bookmaker
package bookmakereu

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

func ParseXML(r io.Reader) (XMLLeagues, error) {
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel
	v := XMLLeagues{}
	e := d.Decode(&v)
	return v, e
}

type XMLLeagues struct {
	Leagues []XMLLeague `xml:"Leagues>league"`
}

type XMLLeague struct {
	IdLeague    string      `xml:"IdLeague,attr"`
	IdSport     string      `xml:"IdSport,attr"`
	Description string      `xml:"Description,attr"`
	Banners     []XMLBanner `xml:"banner"`
	Games       []XMLGame   `xml:"game"`
}

type XMLBanner struct {
	Ab  string `xml:"ab,attr"`
	Vtm string `xml:"vtm,attr"`
	Htm string `xml:"htm,attr"`
}

type XMLGame struct {
	Idgm      string    `xml:"idgm,attr"`
	Gdesc     string    `xml:"gdesc,attr"`
	Idgmtyp   string    `xml:"idgmtyp,attr"`
	Gmdt      string    `xml:"gmdt,attr"`
	Idlg      string    `xml:"idlg,attr"`
	Gmtm      string    `xml:"gmtm,attr"`
	Idspt     string    `xml:"idspt,attr"`
	Vpt       string    `xml:"vpt,attr"`
	Hpt       string    `xml:"hpt,attr"`
	Vnum      string    `xml:"vnum,attr"`
	Hnum      string    `xml:"hnum,attr"`
	Evtyp     string    `xml:"evtyp,attr"`
	Idgp      string    `xml:"idgp,attr"`
	Gpd       string    `xml:"gpd,attr"`
	Vtm       string    `xml:"vtm,attr"`
	Htm       string    `xml:"htm,attr"`
	Stats     string    `xml:"stats,attr"`
	PropCount string    `xml:"propCount,attr"`
	Descgmtyp string    `xml:"descgmtyp,attr"`
	Lines     []XMLLine `xml:"line"`
}

type XMLLine struct {
	Voddst     string `xml:"voddst,attr"`
	Hoddst     string `xml:"hoddst,attr"`
	Ovt        string `xml:"ovt,attr"`
	Ovoddst    string `xml:"ovoddst,attr"`
	Unt        string `xml:"unt,attr"`
	Unoddst    string `xml:"unoddst,attr"`
	Vsprdt     string `xml:"vsprdt,attr"`
	Vsprdoddst string `xml:"vsprdoddst,attr"`
	Hsprdt     string `xml:"hsprdt,attr"`
	Hsprdoddst string `xml:"hsprdoddst,attr"`
	Vspt       string `xml:"vspt,attr"`
	Vspoddst   string `xml:"vspoddst,attr"`
	Hspt       string `xml:"hspt,attr"`
	Hspoddst   string `xml:"hspoddst,attr"`
	Voddsh     string `xml:"voddsh,attr"`
	Hoddsh     string `xml:"hoddsh,attr"`
	Vsprdh     string `xml:"vsprdh,attr"`
	Hsprdh     string `xml:"hsprdh,attr"`
	Ovh        string `xml:"ovh,attr"`
	Unh        string `xml:"unh,attr"`
	Vsph       string `xml:"vsph,attr"`
	Hsph       string `xml:"hsph,attr"`
	Voddshr    string `xml:"voddshr,attr"`
	Vsprdhr    string `xml:"vsprdhr,attr"`
	Ovhr       string `xml:"ovhr,attr"`
	Vsphr      string `xml:"vsphr,attr"`
	Btot       string `xml:"btot,attr"`
	Bsprd      string `xml:"bsprd,attr"`
	Bml        string `xml:"bml,attr"`
	Haschild   string `xml:"haschild,attr"`
	Related    string `xml:"related,attr"`
	EmptyGame  string `xml:"EmptyGame,attr"`
	Tmnum      string `xml:"tmnum,attr"`
	Tmname     string `xml:"tmname,attr"`
	Odds       string `xml:"odds,attr"`
	Oddsh      string `xml:"oddsh	,attr"`
}
