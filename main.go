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

	indata, err := getIn(in)
	if err != nil {
		log.Fatal(err)
	}

	outData, err := getOut(out)
	if err != nil {
		log.Fatal(err)
	}

	defer outData.Close()

	res, err := countWords(indata)
	if err != nil {
		log.Fatal(err)
	}

	w := sortWords(res, *reverse)

	// TODO: usare json.NewEncoder per la codifica a json, cosi puoi passare il io.Writer ottenuto da getOut
	// enc, err := json.NewEncoder(w)
	// enc.Encode(w)

	if *j {
		res, err := json.MarshalIndent(w, "", "  ")
		if err != nil {
			log.Fatal(err)
			return
		}
		outData.Write(res)
		return
	}
	enc := json.NewEncoder(outData)
	enc.Encode(w)
}

func getIn(in string) (io.Reader, error) {
	if in == "-" {
		return os.Stdin, nil
	}
	data, err := os.Open(in)
	if err != nil {
		return nil, fmt.Errorf("can't open input file: %s", err)
	}
	return data, nil
}

func getOut(outArg string) (io.WriteCloser, error) {
	if outArg != "" {
		file, err := os.Create(outArg)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
	return os.Stdout, nil
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
