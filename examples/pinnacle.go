package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/c4e8ece0/bmscraper/org/pinnacle"
)

func main() {
	file, e := os.Open("pinnacle.xml")
	if e != nil {
		log.Fatalln(e)
	}
	defer file.Close()

	tree, e := pinnacle.ParseXML(file)
	if e != nil {
		log.Fatalln(e)
	}

	j, e := json.MarshalIndent(tree, "", "\t")
	if e != nil {
		log.Fatalln(e)
	}

	ioutil.WriteFile("pinnacle.res.json", j, 0777)
	fmt.Print("Done!")
}
