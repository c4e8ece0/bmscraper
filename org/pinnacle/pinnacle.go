// Package pinnacle contents implementations of type.Bookmaker
package pinnacle

import (
	"encoding/xml"
	"io"

	"golang.org/x/net/html/charset"

	// "github.com/c4e8ece0/bmscraper/org"
	"github.com/c4e8ece0/bmscraper/types"
)

// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// READ XML-SOURCE FOR DETAILS ABOUT PARTIAL UPDATES
// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

type Unit struct{ XML }

func init() {
	// org.Register(types.Bookmaker{new(Unit)})
}

type XML struct {
	types.Bookmaker
}

func ParseXML(r io.Reader) (XMLFeed, error) {
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel
	v := XMLFeed{}
	e := d.Decode(&v)
	return v, e
}

type XMLFeed struct {
	FeedTime    string     `xml:"PinnacleFeedTime"`
	LastContest string     `xml:"lastContest"`
	LastGame    string     `xml:"lastGame"`
	Events      []XMLEvent `xml:"events>event"`
}

type XMLEvent struct {
	DatetimeGMT    string           `xml:"event_datetimeGMT"`
	GameNumber     string           `xml:"gamenumber"`
	SportType      string           `xml:"sporttype"`
	League         string           `xml:"league"`
	ContestMaximum string           `xml:"contest_maximum"`
	Description    string           `xml:"description"`
	IsLive         string           `xml:"IsLive"`
	Participants   []XMLParticipant `xml:"participants>participant"`
	Periods        []XMLPeriod      `xml:"periods>period"`
	Total          XMLTotal         `xml:"total"`
}

type XMLParticipant struct {
	Name             string   `xml:"participant_name"`
	ContestantNum    string   `xml:"contestantnum"`
	RotNum           string   `xml:"rotnum"`
	VisitingHomeDraw string   `xml:"visiting_home_draw"` // Visiting or Home
	Odds             []XMLOdd `xml:"odds"`
	Pitcher          string   `xml:"pitcher"`
}

type XMLOdd struct {
	MoneylineValue string `moneyline_value"`
	ToBase         string `to_base"`
}
type XMLPeriod struct {
	PeriodNumber            string       `xml:"period_number"`
	PeriodDescription       string       `xml:"period_description"`
	PeriodCutOffDatetimeGMT string       `xml:"periodcutoff_datetimeGMT"`
	PeriodStatus            string       `xml:"period_status"`
	PeriodUpdate            string       `xml:"period_update"`
	SpreadMaximum           string       `xml:"spread_maximum"`
	MoneylineMaximum        string       `xml:"moneyline_maximum"`
	TotalMaximum            string       `xml:"total_maximum"`
	MoneyLine               XMLMoneyline `xml:"moneyline"`
	Spread                  XMLSpread    `xml:"spread"`
	Total                   XMLTotal     `xml:"total"`
}

type XMLMoneyline struct {
	Visiting string `xml:"moneyline_visiting"`
	Home     string `xml:"moneyline_home"`
	Draw     string `xml:"moneyline_draw"`
}

type XMLSpread struct {
	Visiting       string `xml:"spread_visiting"`
	AdjustVisiting string `xml:"spread_adjust_visiting"`
	Home           string `xml:"spread_home"`
	AdjustHome     string `xml:"spread_adjust_home"`
}

type XMLTotal struct {
	TotalPoints string `xml:"total_points"`
	OverAdjust  string `xml:"over_adjust"`
	UnderAdjust string `xml:"under_adjust"`
	Units       string `xml:"units"`
}
