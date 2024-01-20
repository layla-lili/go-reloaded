package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	// "unicode"

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
	// words = replaceaWithan(words)
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

//  func addSpaces(str []string) []string {
// 	 re := regexp.MustCompile(`\W([.,!,?,:;])(\s*)(\S*)`)
// 	temp := strings.Join(str, " ")
// 	 s := re.ReplaceAllString(temp, "$1$3")
// 	output := strings.Split(s, " ")
// 	return output
//  }
func addSpaces(str []string) []string {
	re := regexp.MustCompile(`(\w)\s*([.,!?:;])\s*(\S*)`)
	temp := strings.Join(str, " ")
	s := re.ReplaceAllString(temp, "$1$2$3")
	output := strings.Fields(s)
	return output
}


//  func addSpaces(s []string) []string {
// 	stringin := strings.Join(s, " ")
// 	re := regexp.MustCompile(`(?s)\s*([,.!?:;]+)\s*`)
// 	result := re.ReplaceAllStringFunc(stringin, func(s string) string {
//         // Replace single punctuation marks surrounded by whitespace
//         if len(s) > 1 && (s[0] == ' ' || s[len(s)-1] == ' ') {
//             return strings.TrimSpace(s) + " "
//         }

//         return s
//     })
// 	return strings.Split(result, " ")

//  }






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











func replaceAWithAn(input []string) []string {
	stringStr := strings.Join(input, " ")

	re := regexp.MustCompile(`(A|a)\s+\b[aeiouh]\w+\b`)
	result := re.ReplaceAllStringFunc(stringStr, func(s string) string {
		word := regexp.MustCompile(`\b\w+\b`).FindString(s)
		if isVowel(word[0]) {
			if word[0] == 'A' {
				return "An" + s[1:]
			} else {
				return "an" + s[1:]
			}
		}
		return s
	})

	return strings.Split(result, " ")
	}

func isVowel (c byte) bool{
	switch c {
	case 'a','e','i','o','u','h','A','E','I','O','U','H':
		return true
	}
	return false
}
// func replaceaWithan(s []string) []string {
// 	for i := 0; i < len(s)-1; i++ {
// 		if
// 		strings.HasPrefix(strings.ToLower(s[1+i]), "a") || 
// 		strings.HasPrefix(strings.ToLower(s[1+i]), "e") || 
// 		strings.HasPrefix(strings.ToLower(s[1+i]), "i") ||
// 	    strings.HasPrefix(strings.ToLower(s[1+i]), "o") ||
// 		strings.HasPrefix(strings.ToLower(s[1+i]), "h") ||
// 		strings.HasPrefix(strings.ToLower(s[1+i]), "u") {
//         s[i] = newFunction(s[i], i)
// 		}
// 	}
// 	return s
// }
// func newFunction(s string, i int) string {
// 	return strings.Replace(s, "a", " an ", -1)
// }

// func replaceAWithAn(s []string) []string {
// 	for i := 0; i < len(s)-1; i++ {
// 		if
// 	    strings.HasPrefix(strings.ToUpper(s[1+i]), "A") || 
// 		strings.HasPrefix(strings.ToUpper(s[1+i]), "E") ||
// 		strings.HasPrefix(strings.ToUpper(s[1+i]), "I") || 
// 		strings.HasPrefix(strings.ToUpper(s[1+i]), "O") ||
// 		strings.HasPrefix(strings.ToUpper(s[1+i]), "H") ||
// 		strings.HasPrefix(strings.ToUpper(s[1+i]), "U") {
//         s[i] = newFunctions(s[i], i)
// 		}
// 	}
// 	return s
// }
// func newFunctions(s string, i int) string {
// 	return strings.Replace(s, "A", "An", -1)
// }
   
//  func replaceAWithAn(s []string) []string {
	// 	for i := 0; i < len(s)-1; i++ {
	// 		// re := regexp.MustCompile(`(\w*)`)
	// 	if strings.ContainsAny(s[1+i], "A E I O U H a e i o u h") {
	// 		s[i] = strings.Replace(s[i],"A", "An", -1)
	// 		s[i] = strings.Replace(s[i],"a", "an", -1)

	// 	}
	// }
	// return s
// 	for i := 0; i < len(s)-1; i++ {
// 		if s[i] == "A" || s[i] == "a" {
// 			nextWordFirstRune := []rune(s[i+1])[0]
// 			if unicode.IsUpper(nextWordFirstRune) {
// 				s[i] = "An"
// 			} else {
// 				s[i] = "an"
// 			}
// 		}
// 	}
// 	return s
// }

















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
