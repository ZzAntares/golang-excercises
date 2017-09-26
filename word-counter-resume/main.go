package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

type Pair struct {
	Token     string
	Frequency int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Frequency < p[j].Frequency }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func PruneToken(token string) []string {
	var symbols string = "?–-_’—.,;:“”()/\\'\"+=~"

	splitter := func(r rune) bool {
		for _, s := range symbols {
			if r == s {
				return true
			}
		}

		return false
	}

	return strings.FieldsFunc(token, splitter)
}

func RankValueMap(freq map[string]int) []Pair {
	var pairs PairList = make(PairList, len(freq))

	var i int
	for k, v := range freq {
		pairs[i] = Pair{Token: k, Frequency: v}
		i++
	}

	sort.Sort(sort.Reverse(pairs))

	return pairs
}

func WordCount(rdr, prepos io.Reader) map[string]int {
	bytes, err := ioutil.ReadAll(prepos)
	if err != nil {
		log.Fatalln(err)
	}
	var contents string = string(bytes)

	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)

	var dict map[string]int = make(map[string]int)

	for scanner.Scan() {
		var words []string = PruneToken(scanner.Text())

		for _, w := range words {
			var word string = strings.ToLower(w)
			if !strings.Contains(contents, word) {
				dict[word]++
			}
		}
	}

	return dict
}

func WordCountFiles(textFile, preposFile string) map[string]int {
	f, err := os.Open(textFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	prepos, err := os.Open(preposFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer prepos.Close()

	return WordCount(f, prepos)
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Usage: word-count <FILE> <PREPOS FILE>")
	}

	var wordCounts map[string]int = WordCountFiles(os.Args[1], os.Args[2])

	var rank PairList = RankValueMap(wordCounts)
	var longest, popular Pair
	var i int

	var pareto int = len(rank) / 80

	for _, pair := range rank {
		if i < pareto {
			fmt.Printf("%s\t\t%d\n", pair.Token, pair.Frequency)
		}

		if len(longest.Token) < len(pair.Token) {
			longest = pair
		}

		if popular.Frequency < pair.Frequency {
			popular = pair
		}

		i++
	}

	fmt.Println("Word whale:", wordCounts["whale"])
	fmt.Println("Longest word:", longest)
	fmt.Println("Popular word:", popular)
}
