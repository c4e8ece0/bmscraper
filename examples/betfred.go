package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/c4e8ece0/bmscraper/org/betfred"
)

func main() {
	main_templates()
	main_odds()
	fmt.Print("Done!")
}

func main_odds() {
	file, e := os.Open("betfred.xml") // template = "Football-Premiership"
	if e != nil {
		log.Fatalln(e)
	}
	defer file.Close()

	tree, e := betfred.ParseXML(file)
	if e != nil {
		log.Fatalln(e)
	}

	j, e := json.MarshalIndent(tree, "", "\t")
	if e != nil {
		log.Fatalln(e)
	}

	ioutil.WriteFile("betfred.res.json", j, 0777)
}

func main_templates() {
	file, e := os.Open("betfred-templates.xml")
	if e != nil {
		log.Fatalln(e)
	}
	defer file.Close()

	tree, e := betfred.ParseTemplateXML(file)
	if e != nil {
		log.Fatalln(e)
	}

	j, e := json.MarshalIndent(tree, "", "\t")
	if e != nil {
		log.Fatalln(e)
	}

	ioutil.WriteFile("betfred-templates.res.json", j, 0777)

}
