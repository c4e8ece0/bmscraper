// Package bwin implements interface bmtype.Bookmaker
package bwin

import (
	"../../bmtype"
)

type Unit struct {
	HTTP // or XML or JSON, i.e. simplest variant of the existing
}

// HTTP-version of scraper
type HTTP struct {
	bmtype.Bookmaker
}

// XML-version of scraper
type XML struct {
	bmtype.Bookmaker
}

// JSON-version of scraper
type JSON struct {
	bmtype.Bookmaker
}
