package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/c4e8ece0/bmscraper/org/williamhill"
)

func main() {
	file, e := os.Open("williamhill.xml") // 556104d7a4b2160d85f962f71d821a0c.xml
	if e != nil {
		log.Fatalln(e)
	}
	defer file.Close()

	tree, e := williamhill.ParseXML(file)
	if e != nil {
		log.Fatalln(e)
	}

	j, e := json.MarshalIndent(tree, "", "\t")
	if e != nil {
		log.Fatalln(e)
	}

	ioutil.WriteFile("williamhill.res.json", j, 0777)

	fmt.Print("Done!")
}
