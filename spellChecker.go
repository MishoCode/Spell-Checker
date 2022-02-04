package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"spell_checker/similarity"
	"log"
	"strconv"
	//"io/ioutil"
)

/*func sayHi() {
	fmt.Println("Hello wolrd!")
}*/

func copyContent(source string, destination *os.File) {
    /*bytesRead, err := ioutil.ReadFile(source)

    if err != nil {
        log.Fatal(err)
    }

    err = ioutil.WriteFile(destination, bytesRead, 0644)

    if err != nil {
        log.Fatal(err)
    }*/

	destination.WriteString("= = = Original input = = = \n")
	f, err := os.Open(source)
	if err != nil {
		fmt.Printf("error opening file: %v\n",err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, e := Readln(r)
	for e == nil {
		destination.WriteString(s + "\n")
		s, e = Readln(r)
	}
}

type spellChecker struct {
	dictionary []string
	stopwords []string
	targetFiles []string
	outputFile string
	//o *os.File
}

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func Readln(r *bufio.Reader) (string, error) {
	var (isPrefix bool = true
		 err error = nil
		 line, ln []byte
		)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln),err
  }

func (sc spellChecker) analyzeFile(file string) {

	out, err := os.OpenFile(sc.outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY | os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	
	copyContent(file, out)
	out.WriteString("\n= = = Metadata = = = \n")

	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("error opening file: %v\n",err)
		os.Exit(1)
	}

	out.WriteString("= = = Findings = = =\n")
	r := bufio.NewReader(f)
	s, e := Readln(r)
	var i uint32 = 1
	for e == nil {
		//fmt.Println(s , i) //process line i instead of print
		sc.processLine(s, i, out)
		i++
		s,e = Readln(r)
	}
}

func (sc spellChecker) processLine(line string, lineNumber uint32, out *os.File) {
	words := strings.Split(line, " ")
	trim(words)
	//fmt.Println(lineNumber, words[0])
	sc.validateWords(words, lineNumber, out)
	//split the line into words - done
	//check whether the word is valid(i.e. it is in the dictionary and it is not a stopwprd)
	//if yes, continue
	//otherwise find most similar words in the dictioanry and build output(i.e. write to the output file)
	//idea: create structure Output that will represent the result of line processing
}

func trim(words []string) {
	for i, w := range words {
		words[i] = strings.Trim(w, ",?!. ")
	}
}

func (sc spellChecker)validateWords(words []string, lineNumber uint32, out *os.File) {

	for _, w := range words {
		if !contains(sc.stopwords, w) && !contains(sc.dictionary, w) {
			mostSimilarWord := similarity.MostSimilarTo(w, sc.dictionary)
			sc.buildOutput(lineNumber, w, mostSimilarWord, out)
		}
	}
}

func (sc spellChecker) buildOutput(lineNumber uint32, wrongWord string, mostSimilarWord string, out *os.File) {
	//_,err := out.WriteString("= = = Findings = = =\n")
	outputLine :=  "Line #" + strconv.FormatUint(uint64(lineNumber), 10) + " "  + wrongWord + 
	" - Possible suggestion is: " + mostSimilarWord + "\n"
	out.WriteString(outputLine)
	/*if err != nil {
        log.Fatal(err)
    }*/
}

func contains(s []string, word string) bool {
    for _, a := range s {
        if a == word {
            return true
        }
    }
    return false
}
