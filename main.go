package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type WordCount struct {
	Word  string
	Count int
}

func main() {
	reverse := flag.Bool("reverse", false, "reverse the input")
	j := flag.Bool("json", false, "print output as JSON")

	flag.Parse()

	in := flag.Arg(0)
	out := flag.Arg(1)

	data, err := getIn(in)
	if err != nil {
		log.Fatal(err)
	}

	res, err := countWords(data)
	if err != nil {
		log.Fatal(err)
	}

	w := sortWords(res, *reverse)

	if *j {
		res, err := json.MarshalIndent(w, "", "  ")
		if err != nil {
			log.Fatal(err)
			return
		}
		outResult(out, string(res))
		return
	}

	outResult(out, fmt.Sprintf("%v", w))
}

func getIn(in string) (data io.Reader, err error) {
	if in == "-" {
		return os.Stdin, nil
	} else {
		data, err := os.Open(in)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}

func countWords(r io.Reader) (map[string]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	count := 0
	m := make(map[string]int)

	for scanner.Scan() {
		m[scanner.Text()] = count
		count++
	}

	return m, scanner.Err()
}

func sortWords(m map[string]int, reverse bool) []WordCount {
	words := make([]WordCount, 0, len(m))
	for w, c := range m {
		world := WordCount{Word: w, Count: c}
		words = append(words, world)
	}
	if reverse {
		sort.Slice(words, func(i, j int) bool { return words[i].Count > words[j].Count })
	} else {
		sort.Slice(words, func(i, j int) bool { return words[i].Count < words[j].Count })
	}

	return words
}

func outResult(outArg string, content string) {
	if outArg != "" {
		file, err := os.Create(outArg)
		if err != nil {
			log.Fatalf("can't create file: %v", err)
		}
		defer file.Close()

		fmt.Fprint(file, content)
		return
	}

	fmt.Println(content)
}
