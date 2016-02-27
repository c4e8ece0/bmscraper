package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/c4e8ece0/bmscraper/org/smarkets"
)

func main() {
	file, e := os.Open("smarkets.xml")
	if e != nil {
		log.Fatalln(e)
	}
	defer file.Close()

	tree, e := smarkets.ParseXML(file)
	if e != nil {
		log.Fatalln(e)
	}

	j, e := json.MarshalIndent(tree, "", "\t")
	if e != nil {
		log.Fatalln(e)
	}

	ioutil.WriteFile("smarkets.res.json", j, 0777)
	fmt.Print("Done!")
}
