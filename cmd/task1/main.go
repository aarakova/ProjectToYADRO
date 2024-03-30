package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"github.com/kljensen/snowball"
)

const TextLanguage = "english"

func NewFormatter(pathToFile string) (*Formatter, error) {
	format := Formatter{stopWordsMap: make(map[string]bool)}

	var err error

	if format.file, err = os.OpenFile(filepath.Clean(pathToFile), os.O_RDONLY, os.ModePerm); err != nil {
		return nil, err
	}

	if err = format.loadStopWords(); err != nil {
		return nil, err
	}

	return &format, nil
}

type Formatter struct {
	stopWordsMap map[string]bool

	file *os.File
}

func (f *Formatter) loadStopWords() error {
	scanner := bufio.NewScanner(f.file)
	for scanner.Scan() {
		f.stopWordsMap[scanner.Text()] = true
	}

	return scanner.Err()
}

func (f *Formatter) isStopWord(word string) bool {
	return f.stopWordsMap[word]
}

func (f *Formatter) normalizeText(inputText string) []string {
	var result []string
	seenWords := make(map[string]bool)
	words := strings.FieldsFunc(inputText, func(c rune) bool {
		return unicode.IsSpace(c)
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

func NewNormalizerConfig() *NormalizerConfig {
	var config NormalizerConfig
	flag.StringVar(&config.inputText,
		"src",
		"",
		"enter your string for analyzer",
	)
	flag.StringVar(&config.stopWordsFilePath,
		"stop",
		"stopWords.txt",
		"file containing stop words",
	)
	flag.Parse()

	return &config
}

type NormalizerConfig struct {
	inputText         string
	stopWordsFilePath string
}

func main() {
	conf := NewNormalizerConfig()
	formatter, err := NewFormatter(conf.stopWordsFilePath)
	if err != nil {
		log.Fatalf("cannot create fromatter: %s", err)
	}
	timer := time.Now()

	log.Printf("Нормальизованный текст: \"%s\"", strings.Join(formatter.normalizeText(conf.inputText), " "))

	log.Printf("Время выполнения программы: %fs", time.Since(timer).Seconds())
}
