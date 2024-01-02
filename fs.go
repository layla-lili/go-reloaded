package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var words []string
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Print("file not found")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	// Scan all words from the file.
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Print("file not found")
		return
	}
	words = Bin(words)
	words = Hex(words)
	words = Up(words)
	words = Low(words)
	words = Cap(words)
	words = addSpaces(words)
	words = formatText(words)
	words = removeSpaces(words)
	words = replaceAWithAn(words)
	words = formatText(words)
	words = removeExtraSpaces(words)
	fmt.Println(words)
}

func Hex(s []string) []string {
	for i, word := range s {
		if word == "(hex)" {
			if decimal, err := strconv.ParseInt(s[i-1], 16, 64); err == nil {
				s[i-1] = strconv.Itoa(int(decimal))
			} else {
				fmt.Println("cannot convert a non hex number")
				os.Exit(0)
			}
			s[i] = ""
		}
	}
	return s
}

func Bin(s []string) []string {
	for i, word := range s {
		if word == "(bin)" {
			if decimal, err := strconv.ParseInt(s[i-1], 2, 64); err == nil {
				s[i-1] = strconv.Itoa(int(decimal))
			} else {
				fmt.Println("cannot convert a non bin number")
				os.Exit(0)
			}
			s[i] = ""
		}
	}
	return s
}

func Up(str []string) []string {
	var sim []string
	for _, word := range str {
		if word == "(up)" && len(sim) > 0 {
			sim[len(sim)-1] = strings.ToUpper(sim[len(sim)-1])
		} else {
			sim = append(sim, word)
		}
	}
	return sim
}

func Low(str []string) []string {
	var sim []string
	for _, word := range str {
		if word == "(low)" && len(sim) > 0 {
			sim[len(sim)-1] = strings.ToLower(sim[len(sim)-1])
		} else {
			sim = append(sim, word)
		}
	}
	return sim
}

func Cap(s []string) []string {
	str := strings.Join(s, " ")
	words := strings.Fields(str)
	for i, word := range words {
		if word == "(cap)" && i > 0 {
			words[i-1] = strings.ToTitle(words[i-1])
			words[i] = ""
		}
		if word == "(cap)" && i > 999 {
			words[i-1] = strings.ToTitle(words[i-1])
			words[i] = ""
		}
	}
	return strings.Fields(strings.Join(words, " "))
}

func addSpaces(str []string) []string {
	re := regexp.MustCompile(`\W([.,!,?,:;])(\s*)(\S*)`)
	temp := strings.Join(str, " ")
	s := re.ReplaceAllString(temp, "$1$3")
	output := strings.Split(s, " ")
	return output
}

func formatText(inputs []string) []string {
	re := regexp.MustCompile(`\w([...,!?])(\s)(\S*)`)
	temp := strings.Join(inputs, " ")
	s := re.ReplaceAllString(temp, "$1$3")
	output := strings.Split(s, " ")
	return output
}

func removeSpaces(strs []string) []string {
	re := regexp.MustCompile(`'\s*(\w+)\s*'`)
	temp := strings.Join(strs, " ")
	newStr := re.ReplaceAllString(temp, "'${1}'")

	output := strings.Split(newStr, " ")
	return output
}

func removeExtraSpaces(input []string) []string {
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	temp := strings.Join(input, " ")
	newStr := re.ReplaceAllString(temp, "'${1}'")
	output := strings.Split(newStr, " ")
	return output
}

func replaceAWithAn(s []string) []string {
	for i := 0; i < len(s)-1; i++ {
		if strings.ContainsAny(s[i+1], "a e i o u h A E I O U H") {
			s[i] = strings.Replace(s[i], "A", "An", -1)
		}
	}
	return s
}
