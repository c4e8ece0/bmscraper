// Package pinnacle contents implementations of type.Bookmaker
package pinnacle

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

func ParseXML(r io.Reader) (XMLFeed, error) {
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel
	v := XMLFeed{}
	e := d.Decode(&v)
	return v, e
}

/*
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

<period>
	<period_number>0</period_number>
	<period_description>Game</period_description>
	<periodcutoff_datetimeGMT>2016-02-19 01:00</periodcutoff_datetimeGMT>
	<period_status>I</period_status>
	<period_update>open</period_update>
	<spread_maximum>2000</spread_maximum>
	<moneyline_maximum>1000</moneyline_maximum>
	<total_maximum>1000</total_maximum>
	<moneyline>
		<moneyline_visiting>121</moneyline_visiting>
		<moneyline_home>-136</moneyline_home>
	</moneyline>
	<spread>
		<spread_visiting>2</spread_visiting>
		<spread_adjust_visiting>-102</spread_adjust_visiting>
		<spread_home>-2</spread_home>
		<spread_adjust_home>-110</spread_adjust_home>
	</spread>
</period>

*/

// PINNACLE XML NOTICE
// IMPORTANT CHANGE - June, 2008

// Some xml consumers have reported incorrect Soccer Total (Asian Handicaps) and others have reported missed updates.

// A release performed on June 24th, 2008 addresses the incorrect odds reported for Soccer Total (Asian Handicaps).

// As for the missed updates, due to the complexity of our platform, a small schema change along with a new way to make incremental calls has been developed.

// It is important to note that the previous functionality, using the last parameter, has been grandfathered so that current users do not need to worry about this upgrade breaking their parsing software.  At some point, we will remove this functionality, so please upgrade your software in a timely fashion.

// The feed will now contain two new elements, lastGame and lastContest.

// If you are interested in games only you can use the value of the element lastGame to make an incremental call.

// For example.
// Get the static file - http://xml.pinnaclesports.com/pinnaclefeed.asp
// Make an incremental call to get up to date - http://xml.pinnaclesports.com/pinaclefeed.asp?lastGame=42145789 (the value for lastGame is found in the static file and subsequent incremental updates)

// If you are also interested in contests, simply add the parameter lastContest to the query string with the value found in the static file.

// For example.
// Get the static file - http://xml.pinnaclesports.com/pinnaclefeed.asp
// Make an incremental call to get up to date - http://xml.pinnclesports.com/pinnaclefeed.asp?lastGame=42145789&lastContest=4678

// You can also mix and match sportType, sportSubType with lastGame if you are only interested in one sport (see examples way below, mixing sporttype, sportSubType and last).

// Attention:  lastContest does not work without lastGame.

// Attention:  last supercedes the usage of either lastGame or lastContest.  If you issue both last and either of lastGame or lastContest in the same call, only last will be used.

// Warning:  The usage of SportType and SportSubtype is being abused.  The intent is to only be used by those who are only interested in one or two specific leagues.  We have observed, from the same IP, calls to all individual soccer leagues (more than 800 of them) in rapid succession.  This is a situation that calls for getting the static file, and then getting incremental updates without a sportType/sportSubtype being specified.  This type of abuse puts a lot of stress on our backend servers.  We are monitoring and will limit access to our systems from IP exhibiting this type of behavior.

// IMPORTANT CHANGE

// Pinnacle Sports has released a new and updated version of our live lines feed. A working knowledge of XML is required for integration purposes. Pinnacle Sports regrets that it is unable to provide additional programming support.

// USAGE:

// -  Get the file produced by http://xml.pinnaclesports.com/pinnacleFeed.asp. The server will produce a static file which is updated every 10 minutes.
// -  Make an incremental call to the same URL to receive the latest updates since the static file was produced by using the PinnacleFeedTime found in the previous call (example http://xml.pinnaclesports.com/pinnacleFeed.asp?last=1196199051920)
// -  The PinnacleFeedTime is a critical part of this system and represents the number of milliseconds since the epoch. To compensate for the various latencies between the time at which the PinnacleFeedTime value was generated and the production of the static file, it is perfectly acceptable to deduct a certain number of seconds (in milliseconds) from this value.  Please note however, that this can result in duplicate updates which you will need to be able to resolve.

// PARAMETERS:

// last: This is used to receive incremental updates using the value of PinncelFeedTime. Always use the value of PinnacleFeedTime found in the previous evocation of the xml feed. A PinnacleFeedTime older than 60 minutes, in the absence of the sportType parameter, will have no effect. If you do not include the PinnacleFeedTime, you will receive the same static file as if you had issued a call without it (unless the sporttype parameter was used, see next section).

// sporttype: (maximum size of 20 characters) This can be used to restrict the feed to a particular sport (examples of a valid sporttype are Baseball, Football (for American Football), Handball, Hockey, Soccer etc.). With the exception of soccer, providing a sporttype parameter will result in a dynamic feed with lines current at the time of the call. Due to the number of individual leagues in soccer, please refer to the Standard usage example below to receive up to date lines for soccer. A list of valid "sports" can be found at http://xml.pinnaclesports.com/leagues.asp.

// sportsubtype:  (maximum size of 12 characters) This can be used to restrict the feed to a chosen League only for a particular sport. When using a specific sportsubtype, you must include a valid sporttype to identify the sport and league. For example, NCAA could refer to either Basketball or Football in American college sports. Refer to http://xml.pinnaclesports.com/leagues.asp for the complete list of leagues that are currently available together with the matching 12 character abbreviation.

// contest: This can be used to prevent the display of contests using the following name/value pair in the query string: contest=no.

// EXAMPLES:

// Standard usage
// 1. Get the complete feed with http://xml.pinnaclesports.com/pinnacleFeed.asp
// 2. Make an incremental call to receive the latest changes with http://xml.pinnaclesports.com/pinnacleFeed.asp?last=1196336347638
// 3. Make a second incremental call 60 seconds later to receive the latest changes with http://xml.pinnaclesports.com/pinnacleFeed.asp?last=1196336407638
// 4. In this example it was assumed that the complete feed included a value of 1196336347638 for the PinnacleFeedTime. Please ensure to always use the value present in the feed obtained in the previous call.

// Interested in Hockey only
// 1.  Get the full list of hockey games with http://xml.pinnaclesports.com/pinnacleFeed.asp?sportType=Hockey&contest=no
// 2.  Make an incremental call to receive the latest changes in Hockey with ttp://xml.pinnaclesports.com/pinnacleFeed.asp?sportType=Hockey&last=1196336347639&contest=no
// 3.  Make another incremental call 60 seconds later to receive the latest changes with http://xml.pinnaclesports.com/pinnacleFeed.asp?sportType=Hockey&last=1196336407639&contest=no
// 4.  In this example it was assumed that the complete feed included a value of 1196336347639 for the PinnacleFeedTime. Please ensure to always use the value present in the feed obtained in the previous call.

// Interested in Soccer only
// 1. Get the complete feed with http://xml.pinnaclesports.com/pinnacleFeed.asp
// 2. Make an incremental call to receive the latest changes with http://xml.pinnaclesports.com/pinnacleFeed.asp?sportType=Soccer&last=1196336347640
// 3. Make another incremental call 60 seconds later to receive the latest changes with http://xml.pinnaclesports.com/pinnacleFeed.asp?sportType=Hockey&last=1196336407640&contest=no
// 4. In this example it was assumed that the complete feed included a value of 1196336347640 for the PinnacleFeedTime. Please ensure to always use the value present in the feed obtained in the previous call.

// IMPORTANT NOTICE:

// Please refrain from using a very high frequency of calls to the xml feed. Pinnacle Sports reserves the right to monitor usage of the XML feed and block the IP address range of any user that abuses this service. 1 call per minute is considered an acceptable usage of the feed.

// UPDATE September, 2010:

// There is a new element ‘IsLive’, it indicates whether the event is dealt live or not. All live soccer spread lines are in the 'full game' format. Please note that on the website the spread odds are in the 'rest of the game' format

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

type XMLPeriod struct {
	period_number            string       `xml:"period_number"`
	period_description       string       `xml:"period_description"`
	periodcutoff_datetimeGMT string       `xml:"periodcutoff_datetimeGMT"`
	period_status            string       `xml:"period_status"`
	period_update            string       `xml:"period_update"`
	spread_maximum           string       `xml:"spread_maximum"`
	moneyline_maximum        string       `xml:"moneyline_maximum"`
	total_maximum            string       `xml:"total_maximum"`
	MoneyLine                XMLMoneyline `xml:"moneyline"`
	Spread                   XMLSpread    `xml:"spread"`
	Total                    XMLTotal     `xml:"total"`
}

type XMLOdd struct {
	MoneylineValue string `moneyline_value"`
	ToBase         string `to_base"`
}

type XMLParticipant struct {
	Name             string   `xml:"participant_name"`
	ContestantNum    string   `xml:"contestantnum"`
	RotNum           string   `xml:"rotnum"`
	VisitingHomeDraw string   `xml:"visiting_home_draw"` // Visiting or Home
	Odds             []XMLOdd `xml:"odds"`
	Pitcher          string   `xml:"pitcher"`
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

type XMLFeed struct {
	FeedTime    string     `xml:"PinnacleFeedTime"`
	LastContest string     `xml:"lastContest"`
	LastGame    string     `xml:"lastGame"`
	Events      []XMLEvent `xml:"events>event"`
}
