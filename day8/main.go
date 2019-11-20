package day8

import (
	"aoc2015/helpers"
	"fmt"
	"regexp"
)

var regexHex = regexp.MustCompile(`(?:[^\\]|^)\\\\x`)
var regexQuote = regexp.MustCompile(`(?:[^\\]|^)\\\\"`)
var regexSlash = regexp.MustCompile(`(?:[^\\]|^)\\\\\\`)

func GetStringCodeCharCount(str string) int {
	return len(str)
}

func GetStringMemoryCharacterCount(str string) int {
	// Drop the string delimiter quotes
	str = str[1:len(str) - 1]
	quoteCount := len(helpers.StringContainsSubstringCount(str, `\"`))
	slashCount := len(helpers.StringContainsSubstringCount(str, `\\`))
	hexCount := len(helpers.StringContainsSubstringCount(str, `\x`))
	rhex := regexHex.FindAllStringSubmatch(str, -1)
	if len(rhex) > 0 && len(rhex[0]) > 0 {
		hexCount -= len(rhex[0])
	}
	rquote := regexQuote.FindAllStringSubmatch(str, -1)
	if len(rquote) > 0 && len(rquote[0]) > 0 {
		quoteCount -= len(rquote[0])
	}
	rslash := regexSlash.FindAllStringSubmatch(str, -1)
	if len(rslash) > 0 && len(rslash[0]) > 0 {
		slashCount -= len(rslash[0])
	}
	return len(str) - quoteCount - slashCount - (hexCount * 3)
}

func GetEncodedLen(str string) int {
	quoteCount := helpers.StringContainsCharCount(str, '"')
	slashCount := helpers.StringContainsCharCount(str, '\\')
	return len(str) + quoteCount + slashCount + 2
}

func DoSilver()  {
	santasList := helpers.LoadInputAsStringSlice(8, "\n")
	charCount := 0
	memCount := 0
	for _, str := range santasList {
		charCount += GetStringCodeCharCount(str)
		memCount += GetStringMemoryCharacterCount(str)
	}
	fmt.Println(charCount, memCount)
	fmt.Println("Solution:", charCount - memCount)
}

func DoGold() {
	santasList := helpers.LoadInputAsStringSlice(8, "\n")
	encodeCount := 0
	codeCount := 0
	for _, str := range santasList {
		encodeCount += GetEncodedLen(str)
		codeCount += GetStringCodeCharCount(str)
	}
	fmt.Println(encodeCount, codeCount)
	fmt.Println("Solution: ", encodeCount - codeCount)
}