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
	scanner.Split(bufio.ScanWords)
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
	words = transformText(words)
	words = addSpaces(words)
	words = formatText(words)
	words = removeSpaces(words)
	words = replaceAWithAn(words)
	words = formatText(words)
	words = removeExtraSpaces(words)

	outputFile, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("Failed to create output file:", err)
		return
	}
	defer outputFile.Close()
	writer := bufio.NewWriter(outputFile)
	for _, word := range words {
		_, err := writer.WriteString(word + " ")
		if err != nil {
			fmt.Println("Failed to write word:", err)
			return
		}
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("Failed to flush writer:", err)
		return
	}
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
			words[i-1] = strings.Title(words[i-1])
			words[i] = ""
		}
		if word == "(cap)" && i > 999 {
			words[i-1] = strings.Title(words[i-1])
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
		if strings.ContainsAny(s[i+1], "AEIOUHaeiouh") {
			s[i] = strings.Replace(s[i], "A", "An", -1)
		}
	}
	return s
}

func transformText(input []string) []string {
	re := regexp.MustCompile(`(\d+)\)`)
	for idx, word := range input {
		switch word {
		case "(low,":
			if idx != len(input) && re.MatchString(input[idx+1]) {
				matches := re.FindStringSubmatch(input[idx+1])
				numWords, _ := strconv.Atoi(matches[1])
				for i := idx; i >= idx-numWords && i != -1; i-- {
					input[i] = strings.ToLower(input[i])
				}
				input = append(input[:idx], input[idx+1:]...)

			}
		case "(up,":
			if idx != len(input) && re.MatchString(input[idx+1]) {
				matches := re.FindStringSubmatch(input[idx+1])
				numWords, _ := strconv.Atoi(matches[1])
				for i := idx; i >= idx-numWords && i != -1; i-- {
					input[i] = strings.ToUpper(input[i])
				}
				input = append(input[:idx], input[idx+1:]...)

			}
		case "(cap,":
			if idx != len(input) && re.MatchString(input[idx+1]) {
				matches := re.FindStringSubmatch(input[idx+1])
				numWords, _ := strconv.Atoi(matches[1])
				for i := idx; i >= idx-numWords && i != -1; i-- {
					input[i] = strings.Title(input[i])
				}
				input = append(input[:idx], input[idx+1:]...)
			}

		}
	}
	temp := strings.Join(input, " ")
	newStr := re.ReplaceAllString(temp, "")
	output := strings.Fields(newStr)
	return output
}
