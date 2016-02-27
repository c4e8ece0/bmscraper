package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/c4e8ece0/bmscraper/org/base/betconstruct"
)

func main() {
	file, e := os.Open("1xbet.json")
	if e != nil {
		log.Fatalln(e)
	}
	defer file.Close()

	tree, e := betconstruct.ParseJSON(file)
	if e != nil {
		log.Fatalln(e)
	}

	j, e := json.MarshalIndent(tree, "", "\t")
	if e != nil {
		log.Fatalln(e)
	}

	ioutil.WriteFile("1xbet.res.json", j, 0777)
	fmt.Print("Done!")
}
