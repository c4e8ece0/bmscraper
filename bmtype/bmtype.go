// Package bmtype contents types and interfaces for odds-scrapper
// of most famous bookmakers
package bmtype

//
type Odd struct {
	sport  string
	first  string
	second string
	result string
}

//
type Bookmaker interface {
	Sports() []string // Return list of sports
	Get() []Odd       // Array of fetched odds
}
