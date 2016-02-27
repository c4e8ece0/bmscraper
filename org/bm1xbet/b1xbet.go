// Package bm1xBet implements types.Bookmaker interface
package bm1xbet

import (
	"github.com/c4e8ece0/bmscraper/org/base/betconstruct"
	// "github.com/c4e8ece0/bmscraper/types/sport"
)

// var feed = map[int]string{
// 	sport.FOOTBALL: "http://...",
// 	sport.HOCKEY:   "http://...",
// }

type Unit struct {
	betconstruct.Unit
}

func (u Unit) Name() {
	return "1xbet"
}
