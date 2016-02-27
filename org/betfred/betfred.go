// package betfred contents implementations of type.Bookmaker
package betfred

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
	Name          string     `xml:"name,attr"`
	TimeGenerated string     `xml:"timeGenerated,attr"`
	Bookmaker     string     `xml:"Bookmaker,attr"`
	Events        []XMLEvent `xml:"event"`
}

type XMLEvent struct {
	Name     string       `xml:"name,attr"`
	Eventid  string       `xml:"eventid,attr"`
	Date     string       `xml:"date,attr"`
	Time     string       `xml:"time,attr"`
	Meeting  string       `xml:"meeting,attr"`
	Venue    string       `xml:"venue,attr"`
	BetTypes []XMLBetType `xml:"bettype"`
}

type XMLBetType struct {
	BetStartDate string   `xml:"bet-start-date,attr"`
	BetStartTime string   `xml:"bet-start-time,attr"`
	Ewreduction  string   `xml:"ewreduction,attr"`
	Ewplaceterms string   `xml:"ewplaceterms,attr"`
	Eachway      string   `xml:"eachway,attr"`
	Suspended    string   `xml:"suspended,attr"`
	Name         string   `xml:"name,attr"`
	Inrunning    string   `xml:"inrunning,attr"`
	Bettypeid    string   `xml:"bettypeid,attr"`
	Handicap     string   `xml:"handicap,attr"`
	Bets         []XMLBet `xml:"bet"`
}

type XMLBet struct {
	Name             string `xml:"name,attr"`
	ShortName        string `xml:"short-name,attr"`
	Id               string `xml:"id,attr"`
	Price            string `xml:"price,attr"`
	PriceDecimal     string `xml:"priceDecimal,attr"`
	PriceUS          string `xml:"priceUS,attr"`
	ActivePriceTypes string `xml:"active-price-types,attr"`
	HadValue         string `xml:"had-value,attr"`
	HandicapId       string `xml:"handicap-id,attr"`
	Handicap         string `xml:"handicap,attr"`
}

// КАКАЯ-ТО ДИЧ, НАДО ПЕРЕПРОВЕРИТЬ ВСЕ ТЕМПЛЕЙТЫ НА ПРЕДМЕТ АТРИБУТОВ И ТЕГОВ
// !!! НУЖЕН PREFETCH ПО ВСЕМ ФАЙЛАМ
// Особенный темплейт: <template name="in_running" is4InRunningOnly="-1" type="Sport Type">
// Внимание на is4InRunningOnly="-1"

// <element sport="CRI-2020INTE"/>
// </template>
// <template name="asian_handicaps" is4InRunningOnly="0" type="Market Type">
// <element marketType="8742.10"/>
// <element marketType="6622.10"/>

// <template name="in_running" is4InRunningOnly="-1" type="Sport Type">
// <element sportType="ATHLETICS"/>
// <element sportType="AUSSIE FOOTY"/>
// <element sportType="BASEBALL"/>
// <element sportType="BASKETBALL"/>

func ParseTemplateXML(r io.Reader) (XMLTemplates, error) {
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel
	v := XMLTemplates{}
	e := d.Decode(&v)
	return v, e
}

type XMLTemplates struct {
	TimeGenerated string        `xml:"timeGenerated,attr"`
	Bookmaker     string        `xml:"Bookmaker,attr"`
	Templates     []XMLTemplate `xml:"template"`
}

type XMLTemplate struct {
	Name             string       `xml:"name,attr"`
	Is4InRunningOnly string       `xml:"is4InRunningOnly,attr"`
	Type             string       `xml:"type,attr"`
	Elements         []XMLElement `xml:"element"`
}

type XMLElement struct {
	MarketType string `xml:"marketType,attr"`
	Sport      string `xml:"sport,attr"`
	SportType  string `xml:"sportType,attr"`
}
