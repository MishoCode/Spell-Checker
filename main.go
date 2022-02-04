package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	//"spell_checker/similarity"
)

func scanDictionary(path string) ([]string, error) {
	file, err := os.Open(path);

	if err != nil {
		return nil, err;
	}

	scanner := bufio.NewScanner(file);
	scanner.Split(bufio.ScanLines);

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, nil

}

func execute(commandParts []string) {
	path := commandParts[1]
	ind := strings.Index(path, "=")
	path = path[ind+1:len(path)-1]

	target := commandParts[2:]

	fmt.Println("You chose to use", path, "as a dictionary")
	fmt.Println("You chose to analyze the following files", target)
	/*dictionary, error := scanDictionary(path);
	if error != nil {
		//create spellChecker with dictionary
		// spellChecker.analyze(given files)
	}*/

}

func initSpellChecker(sc *spellChecker, dictionaryFile string, targetFile string) *spellChecker {
	dictionary, err := scanDictionary(dictionaryFile)
	stopwords, err2 := scanDictionary("stopwords.txt")
	if err == nil && err2 == nil{
		sc.dictionary = dictionary
		sc.stopwords = stopwords
		sc.outputFile = "out.txt"
		sc.targetFiles = []string{"text2.txt"}
	}

	return sc
}


func main() {
	/*scanner := bufio.NewScanner(os.Stdin)
	//fmt.Print("Choose dictionary and source files: ")
	for {
		fmt.Print(">>> ")
        scanner.Scan()
        text := scanner.Text()
        if text != "exit" {
			tokens := strings.Split(text, " ")
			execute(tokens)
        } else if (text == "exit") {
            break
        }

    }*/

	/*sayHi()
	similarity.HelloSimilarity()
	//d := []string{"dict.txt"}
	sc := spellChecker{[]string{"dict.txt"}, []string{"stopwords.txt"}, []string{"stopwords.txt"}, "out.txt"}
	sc.analyzeFile("text.txt")*/

	//r := similarity.MostSimilarTo("hello", []string{"abc, hjk", "123456", "hallo", "yahhhoooo"})
	//fmt.Println(r)

	//copyContent("text.txt", "out.txt")
	//sc.buildOutput(100, "wrongWord", "similar")

	sc := new(spellChecker)
	s := initSpellChecker(sc, "dict.txt", "target.txt")
	s.analyzeFile("text2.txt")
}