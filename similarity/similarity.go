package similarity

import (
	"fmt"
	"math"
)

func HelloSimilarity() {
	fmt.Println("NLP")
}

func find2Grams(word string) []string {
	var result []string

	for i := 0; i < len(word)-1; i++ {
		gram := word[i:i+2]
		result = append(result, gram)
	}

	return result
}

func wordToVector(word string) map[string]int {
	var m = make(map[string]int)
	grams := find2Grams(word)
	for _,v := range grams {
		m[v]++
	}

	return m
}

func getVectorLength(v map[string]int) float64 {
	result := 0
	for _, val := range v {
		result +=  val*val
	}

	return math.Sqrt(float64(result))
}
func dotProduct(v1 map[string]int, v2 map[string]int) int {
	result := 0
	for key1, val1 := range v1 {
		for key2, val2 := range v2 {
			if key2 == key1 && val2 == val1 {
				result += val1*val2
			}
		}
	}

	return result
}

func calculateSimilarity(v1 map[string]int, v2 map[string]int) float64 {
	len1 := getVectorLength(v1)
	len2 := getVectorLength(v2)
	product := dotProduct(v1, v2)

	return float64(product) / (len1 * len2)
}

func MostSimilarTo(word string, dictionary []string) string{
	var max float64 = 0.0
	var result string;
	for _,d := range dictionary {
		current := calculateSimilarity(wordToVector(word), wordToVector(d))
		if current > max {
			result = d
		}
	}

	return result
}

/*func main() {
	fmt.Println(find2Grams("hello"))
	fmt.Println(find2Grams("hallo"))
	fmt.Println(wordToVector("hello"))
	fmt.Println(getVectorLength(wordToVector("hello")))
	fmt.Println(getVectorLength(wordToVector("hallo")))
	fmt.Println(dotProduct(wordToVector("helo"), wordToVector("phelo")))
	fmt.Println(dotProduct(wordToVector("hello"), wordToVector("hallo")))
	fmt.Println(calculateSimilarity(wordToVector("hello"), wordToVector("hallo")))
}*/