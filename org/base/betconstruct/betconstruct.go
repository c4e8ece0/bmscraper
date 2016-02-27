// Package betconstruct contains implementations of type.Unit{JSON}
// Betconstruct is popular developer of betting-sites and data-provider.

package betconstruct

import (
	"encoding/json"
	"io"

	"github.com/c4e8ece0/bmscraper/types"
)

type Unit struct{ JSON }

// func (u *Unit) Get(r io.Reader) {
//
// }

func ParseJSON(r io.Reader) (JSRoot, error) {
	d := json.NewDecoder(r)
	v := JSRoot{}
	e := d.Decode(&v)
	return v, e
}

type JSON struct {
	types.Bookmaker
}

type JSRoot struct {
	Error     string    `json:"Error"`
	ErrorCode int       `json:"ErrorCode"`
	Guid      string    `json:"Guid"`
	Id        int       `json:"Id"`
	Success   bool      `json:"Success"`
	Values    []JSValue `json:"Value"`
}

type JSValue struct {
	ConstId       int            `json:"ConstId"`       //3614811,
	DopInfo       string         `json:"DopInfo"`       //null,
	Events        []JSEvent      `json:"Events"`        //[],
	EventsCount   int            `json:"EventsCount"`   //655,
	Finished      bool           `json:"Finished"`      //false,
	GameType      string         `json:"GameType"`      //"",
	GameTypeId    int            `json:"GameTypeId"`    //1,
	GameVid       string         `json:"GameVid"`       //"",
	Id            int            `json:"Id"`            //66032487,
	LigaId        int            `json:"LigaId"`        //118587,
	MainGameId    int            `json:"MainGameId"`    //66032487,
	NameGame      string         `json:"NameGame"`      //"",
	Num           int            `json:"Num"`           //297,
	Period        int            `json:"Period"`        //0,
	PeriodName    string         `json:"PeriodName"`    //"",
	SportId       int            `json:"SportId"`       //1,
	Top           int            `json:"Top"`           //1000,
	Before        int            `json:"Before"`        //156597,
	Champ         string         `json:"Champ"`         //"Лига Чемпионов УЕФА",
	ChampEng      string         `json:"ChampEng"`      //"UEFA Champions League",
	ChampRus      string         `json:"ChampRus"`      //"Лига Чемпионов УЕФА",
	GameVidEng    string         `json:"GameVidEng"`    //"",
	GameVidRus    string         `json:"GameVidRus"`    //"",
	MainEvents    string         `json:"MainEvents"`    //null,
	Opp1          string         `json:"Opp1"`          //"Арсенал",
	Opp1Eng       string         `json:"Opp1Eng"`       //"Arsenal",
	Opp1Id        int            `json:"Opp1Id"`        //50679,
	Opp1Rus       string         `json:"Opp1Rus"`       //"Арсенал",
	Opp2          string         `json:"Opp2"`          //"Барселона",
	Opp2Eng       string         `json:"Opp2Eng"`       //"Barcelona",
	Opp2Id        int            `json:"Opp2Id"`        //3442,
	Opp2Rus       string         `json:"Opp2Rus"`       //"Барселона",
	SportName     string         `json:"SportName"`     //"Футбол",
	SportNameEng  string         `json:"SportNameEng"`  //"Football",
	SportNameRus  string         `json:"SportNameRus"`  //"Футбол",
	Start         int            `json:"Start"`         //1456256700,
	SubGames      []JSValueSmall `json:"SubGames"`      //[],
	TvChannel     string         `json:"TvChannel"`     //null,
	CurPeriodAsia string         `json:"curPeriodAsia"` //null
}

type JSValueSmall struct {
	ConstId     int       `json:"ConstId"`     //3614811,
	DopInfo     string    `json:"DopInfo"`     //null,
	Events      []JSEvent `json:"Events"`      //[],
	EventsCount int       `json:"EventsCount"` //655,
	Finished    bool      `json:"Finished"`    //false,
	GameType    string    `json:"GameType"`    //"",
	GameTypeId  int       `json:"GameTypeId"`  //1,
	GameVid     string    `json:"GameVid"`     //"",
	Id          int       `json:"Id"`          //66032487,
	LigaId      int       `json:"LigaId"`      //118587,
	MainGameId  int       `json:"MainGameId"`  //66032487,
	NameGame    string    `json:"NameGame"`    //"",
	Num         int       `json:"Num"`         //297,
	Period      int       `json:"Period"`      //0,
	PeriodName  string    `json:"PeriodName"`  //"",
	SportId     int       `json:"SportId"`     //1,
	Top         int       `json:"Top"`         //1000,
}

type JSEvent struct {
	B  bool    `json:"B"`
	C  float32 `json:"C"`
	CV string  `json:"CV"`
	G  int     `json:"G"` // group?
	// P  float32 `json:"P"`  // H1, H2, TL, TM  values
	Pl string `json:"Pl"` //
	T  int    `json:"T"`
}

// OLD INFO

type Values struct {
	B  string
	C  string
	CV string
	G  string // group?
	P  string // H1, H2, TL, TM  values
	Pl string //
	T  string
}

// JS Events by index
type Event struct {
	E1  Values // Events.0
	EX  Values // Events.1
	E2  Values // Events.2
	E1X Values // Events.3
	E12 Values // Events.4
	EX2 Values // Events.5
	H1  Values // Events.6 - Handicap 1 (value = .H1)
	H2  Values // Events.7 - Handicap 2 (value = .H2)
	TL  Values // Events.9 - Total More
	TM  Values // Events.8 - Total Less
}

type Description struct {
	EventsCount   string // => 176
	Finished      string // =>
	GameType      string // =>
	GameTypeId    string // => 1
	GameVid       string // =>
	Id            string // => 64221015
	LigaId        string // => 1120599
	MainGameId    string // => 64221015
	NameGame      string // =>
	Num           string // => 867
	Period        string // => 0
	PeriodName    string // =>
	SportId       string // => 1
	Top           string // => 1000
	Before        string // => 12077455
	Champ         string // => Футбол. Чемпионат Европы 2016
	ChampEng      string // => 2016 UEFA Euro
	ChampRus      string // => Чемпионат Европы 2016
	GameVidEng    string // =>
	GameVidRus    string // =>
	MainEvents    string // =>
	Opp1          string // => Австрия
	Opp1Eng       string // => Austria
	Opp1Id        string // => 11809
	Opp1Rus       string // => Австрия
	Opp2          string // => Венгрия
	Opp2Eng       string // => Hungary
	Opp2Id        string // => 12795
	Opp2Rus       string // => Венгрия
	SportName     string // => Футбол
	SportNameEng  string // => Football
	SportNameRus  string // => Футбол
	Start         string // => 1465923600
	SubGames      string // =>
	TvChannel     string // =>
	curPeriodAsia string // =>
}

/*
PART OF JSON-file

                    [Events] => Array
                        (
                            П1[0] => Array
                                (
                                    [B] =>
                                    [C] => 1.88
                                    [CV] =>
                                    [G] => 1
                                    [P] => 0
                                    [Pl] =>
                                    [T] => 1
                                )

                            Х[1] => Array
                                (
                                    [B] =>
                                    [C] => 3.65
                                    [CV] =>
                                    [G] => 1
                                    [P] => 0
                                    [Pl] =>
                                    [T] => 2
                                )

                            П2[2] => Array
                                (
                                    [B] =>
                                    [C] => 4.85
                                    [CV] =>
                                    [G] => 1
                                    [P] => 0
                                    [Pl] =>
                                    [T] => 3
                                )

                            1X[3] => Array
                                (
                                    [B] =>
                                    [C] => 1.22
                                    [CV] =>
                                    [G] => 8
                                    [P] => 0
                                    [Pl] =>
                                    [T] => 4
                                )

                            12[4] => Array
                                (
                                    [B] =>
                                    [C] => 1.32
                                    [CV] =>
                                    [G] => 8
                                    [P] => 0
                                    [Pl] =>
                                    [T] => 5
                                )

                           X2[5] => Array
                                (
                                    [B] =>
                                    [C] => 2.03
                                    [CV] =>
                                    [G] => 8
                                    [P] => 0
                                    [Pl] =>
                                    [T] => 6
                                )
                            1[6] => Array
                                (
                                    [B] =>
                                    [C] => 2.71
                                    [CV] =>
                                    [G] => 2
                                фора1->    [P] => -1
                                    [Pl] =>
                                    [T] => 7
                                )
                            2 [7] => Array
                                (
                                    [B] =>
                                    [C] => 1.56
                                    [CV] =>
                                    [G] => 2
                                фора2->    [P] => 1
                                    [Pl] =>
                                    [T] => 8
                                )

                            Б[8] => Array
                                (
                                    [B] =>
                                    [C] => 2.29
                                    [CV] =>
                                    [G] => 17
                            Тотал -       [P] => 2.5
                                    [Pl] =>
                                    [T] => 9
                                )

                            М [9] => Array
                                (
                                    [B] =>
                                    [C] => 1.73
                                    [CV] =>
                                    [G] => 17
                            Тотал -     [P] => 2.5
                                    [Pl] =>
                                    [T] => 10
                                )

                        )


                    [EventsCount] => 176
                    [Finished] =>
                    [GameType] =>
                    [GameTypeId] => 1
                    [GameVid] =>
                    [Id] => 64221015
                    [LigaId] => 1120599
                    [MainGameId] => 64221015
                    [NameGame] =>
                    [Num] => 867
                    [Period] => 0
                    [PeriodName] =>
                    [SportId] => 1
                    [Top] => 1000
                    [Before] => 12077455
                    [Champ] => Футбол. Чемпионат Европы 2016
                    [ChampEng] => 2016 UEFA Euro
                    [ChampRus] => Чемпионат Европы 2016
                    [GameVidEng] =>
                    [GameVidRus] =>
                    [MainEvents] =>
                    [Opp1] => Австрия
                    [Opp1Eng] => Austria
                    [Opp1Id] => 11809
                    [Opp1Rus] => Австрия
                    [Opp2] => Венгрия
                    [Opp2Eng] => Hungary
                    [Opp2Id] => 12795
                    [Opp2Rus] => Венгрия
                    [SportName] => Футбол
                    [SportNameEng] => Football
                    [SportNameRus] => Футбол
                    [Start] => 1465923600
                    [SubGames] =>
                    [TvChannel] =>
                    [curPeriodAsia] =>

*/
