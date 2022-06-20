package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	helper "src/main.go/src/helpers"
	"strings"
)

const resultLimit int = 25
const thresholdMultiplier float32 = 0.5

var works []helper.ShakespeareWork

func showResult(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("term") {
		w.Write([]byte("No query received."))
		fmt.Println("No query received.")
		return
	}
	frag := r.URL.Query()["term"][0]
	regFrag, err := regexp.Compile("[^a-zA-Z ]+")
	if err != nil {
		log.Fatal(err)
	}

	results := []helper.ResultsFromLebenshtein{}

	for _, work := range works {
		parsedFrag := strings.ToLower(regFrag.ReplaceAllLiteralString(frag, ""))
		threshold := int(float32(len(parsedFrag)) * thresholdMultiplier)

		if len(parsedFrag) == 0 {
			w.Write([]byte("No query received."))
			fmt.Println("No query received.")
			return
		}

		dist := getLevDistance(work.Title, parsedFrag)
		if dist <= threshold {
			results = append(results, helper.Copy(work, dist))
		}
	}
	helper.Sort(results)

	if len(results) > 0 {
		if len(results) > resultLimit {
			results = results[:resultLimit]
		}
		for _, res := range results {
			fmt.Printf("Results are: %v\n", res.Title)
			w.Write([]byte(res.Title + "\n"))
		}
	} else {
		w.Write([]byte("No results found based on your query: " + frag))
		fmt.Println("No results found based on your query: ", frag)
	}
}

func getLevDistance(title string, frag string) int {

	if len(title) < len(frag) {
		return helper.LevenshteinDistance(strings.ToLower(title), frag)
	}
	return helper.LevenshteinDistance(strings.ToLower(title)[:len(frag)], frag)
}

func loadList() {
	content, err := ioutil.ReadFile("files/shakespeare_works.json")
	if err != nil {
		log.Fatal(err)
	}

	err2 := json.Unmarshal(content, &works)
	if err2 != nil {
		fmt.Println("Error unmarshalling JSON data: ", err2.Error())
	}
}

func main() {
	loadList()

	http.HandleFunc("/autocomplete", showResult)

	if err3 := http.ListenAndServe(":9000", nil); err3 != nil {
		log.Fatal(err3)
	}
}
