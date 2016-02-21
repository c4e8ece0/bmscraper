// Package betconstruct contains implementations of type.Unit{JSON}
// Betconstruct is popular developer of betting-sites and data-provider.

package betconstruct

import (
	"io"

	"github.com/c4e8ece0/bmscraper/types"
)

type Unit struct {
	types.Bookmaker
}

func (u *Unit) Get(r io.Reader) {
	// TODO: parse json
}

type Values struct {
	B  string
	C  string
	CV string
	G  string // group?
	P  string // H1, H2, TL, TM  values
	Pl string //
	T  string
}

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
