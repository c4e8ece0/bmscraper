package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/c4e8ece0/bmscraper/org/intertops"
)

func main() {
	file, e := os.Open("intertops.xml")
	if e != nil {
		log.Fatalln(e)
	}
	defer file.Close()

	tree, e := intertops.ParseXML(file)
	if e != nil {
		log.Fatalln(e)
	}

	j, e := json.MarshalIndent(tree, "", "\t")
	if e != nil {
		log.Fatalln(e)
	}

	ioutil.WriteFile("intertops.res.json", j, 0777)
	fmt.Print("Done!")
}
