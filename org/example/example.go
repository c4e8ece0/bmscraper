// Package example implements interface types.Bookmaker
// and shows the base posibilities for working with bookmakers
package example

import (
	"github.com/c4e8ece0/bmscraper/org"
	"github.com/c4e8ece0/bmscraper/types"
)

func init() {
	org.Register(new(Unit))
}

// Each bookmaker must has its own internal "Unit" struct with
// implementation of types.Bookmaker interface.
// For the case when bookmaker has several variants of data export
// we may switch between them with replacing one word in code.
type Unit struct {
	HTTP // or XML or JSON, i.e. simplest variant of the existing
}

// HTTP-version of scraper
type HTTP struct {
	types.Bookmaker
}

// XML-version of scraper
type XML struct {
	types.Bookmaker
}

// JSON-version of scraper
type JSON struct {
	types.Bookmaker
}
