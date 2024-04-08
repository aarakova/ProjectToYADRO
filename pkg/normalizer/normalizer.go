package normalizer

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/kljensen/snowball"
	log "github.com/sirupsen/logrus"
)

type Data struct {
	ID       string // num
	Url      string // img
	Keywords string // transcript
}

const TextLanguage = "english"

func NewFormatter(pathToFile string) (*Formatter, error) {
	format := Formatter{stopWordsMap: make(map[string]bool)}

	var err error
	var file *os.File

	if file, err = os.OpenFile(filepath.Clean(pathToFile), os.O_RDONLY, os.ModePerm); err != nil {
		return nil, err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Errorf("cannot close file, error: %s", err.Error())
		}
	}()

	if err = format.loadStopWords(file); err != nil {
		return nil, err
	}

	return &format, nil
}

type Formatter struct {
	stopWordsMap map[string]bool
}

func (f *Formatter) loadStopWords(file *os.File) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f.stopWordsMap[scanner.Text()] = true
	}

	return scanner.Err()
}

func (f *Formatter) isStopWord(word string) bool {
	return f.stopWordsMap[word]
}

func (f *Formatter) NormalizeText(inputText string) []string {
	var result []string
	seenWords := make(map[string]bool)
	words := strings.FieldsFunc(inputText, func(c rune) bool {
		return unicode.IsSpace(c) || unicode.IsPunct(c)
	})
	for _, word := range words {
		stemmedWord, _ := snowball.Stem(word, TextLanguage, true)
		if f.isStopWord(stemmedWord) || seenWords[stemmedWord] {
			continue
		}

		result = append(result, stemmedWord)
		seenWords[stemmedWord] = true
	}
	return result
}
