package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kr/logfmt"
)

type pair struct {
	K, V string
}

type data struct {
	Items []pair
}

func (c *data) HandleLogfmt(key, val []byte) error {
	c.Items = append(c.Items, pair{K: string(key), V: string(val)})
	return nil
}

func main() {
	keys := flag.String("k", "", "keys to extract for the log using logfmt. Keys are separated by comma")
	flag.Parse()

	if *keys == "" {
		return
	}

	ks := strings.Split(*keys, ",")
	km := make(map[string]bool)

	for _, k := range ks {
		km[k] = true
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var d data
		err := logfmt.Unmarshal([]byte(s.Text()), &d)
		if err != nil {
			panic(err)
		}

		hasItem := false
		for _, p := range d.Items {
			if km[p.K] {
				if hasItem {
					fmt.Printf(",")
				}

				fmt.Printf("%s", p.V)
				hasItem = true
			}
		}

		if hasItem {
			fmt.Printf("\n")
		}
	}

	if s.Err() != nil {
		panic(s.Err())
	}
}
