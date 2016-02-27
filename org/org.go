// Package org is the entry point for all of the described bookmakers.
// Each bookmaker must register themself by database/sql-init style.
// All interactions goes throught this package.
package org

// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// FIRST TOUCH, DOESN'T MATTER
// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

import (
	"github.com/c4e8ece0/bmscraper/types"
	"sync"
)

var list = make(map[string]types.Bookmaker)
var mu = new(sync.Mutex)

//
func Register(b types.Bookmaker) {
	mu.Lock()
	defer mu.Unlock()
	list[b.Name()] = b
}

//
func Names() []string {
	s := make([]string, len(list))
	mu.Lock()
	defer mu.Unlock()
	for n, _ := range list {
		s = append(s, n)
	}
	return s
}

//
func Stat() {

}
