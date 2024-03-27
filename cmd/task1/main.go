package main

import (
	"flag"
	"log"
	"strings"

	"github.com/kljensen/snowball"
)

var inputText string

var stopWords = "a an the and that but or as if when that because while " +
	"where after so though since until whether before although nor not like " +
	"once unless now except of on in " +
	"i i'm i'll i'd i've you you're you'll you'd you've he he's he'll he'd " +
	"she she's she'll she'd it it's we we're we'll we'd we've " +
	"they they're they'll they'd they've me him her us them my your its our their " +
	"am is are can have had could may might shall should will wount must "

func isStopWord(word string) bool {
	return strings.Contains(stopWords, word)
}

func normalizeText(inputText string) string {
	var result string
	words := strings.Fields(inputText)
	for _, word := range words {
		stemmedWord, _ := snowball.Stem(word, "english", true)
		if isStopWord(stemmedWord) || strings.Contains(result, stemmedWord) {
			continue
		}

		result += stemmedWord + " "
	}
	return result
}

func main() {
	flag.StringVar(&inputText,
		"s",
		"",
		"enter your string for analyzer",
	)
	flag.Parse()

	log.Printf("%s", normalizeText(inputText))
}
