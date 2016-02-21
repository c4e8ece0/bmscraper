// Package types contents types and interfaces for odds-scraper
// from most famous bookmakers
package types

// TODO: make default struct with basic implementation of types.Bookmaker.
// TODO: first of all - betconstruct

import (
	"io"
)

//
type Odd struct {
	coutry     string
	sport      string
	league     string
	tournament string
	first      string
	second     string
	result     string
}

//
type Bookmaker interface {
	Name() string         // Return name of bookmaker
	Register()            // Register bookmaker themself in org package
	New(chan chan string) // New object with chan to pool
	Sports() []string     // Return list of Country+League+Sports+Tournaments
	Get(io.Reader) []Odd  // Read io.Reader and return array of fetched odds
}
