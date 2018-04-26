package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kr/logfmt"
)

type data struct {
	Items map[string]string
}

func (c *data) HandleLogfmt(key, val []byte) error {
	if c.Items == nil {
		c.Items = make(map[string]string)
	}

	c.Items[string(key)] = string(val)
	return nil
}

func main() {
	keys := flag.String("k", "", "keys to extract for the log using logfmt. Keys are separated by comma")
	flag.Parse()

	if *keys == "" {
		return
	}

	ks := strings.Split(*keys, ",")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var d data
		err := logfmt.Unmarshal([]byte(s.Text()), &d)
		if err != nil {
			panic(err)
		}

		hasItem := false
		var values []string
		for _, k := range ks {
			v, ok := d.Items[k]
			if ok {
				hasItem = true
			}

			values = append(values, v)
		}

		if hasItem {
			for i, v := range values {
				if i > 0 {
					fmt.Printf(",")
				}

				fmt.Printf(v)
			}
			fmt.Printf("\n")
		}
	}

	if s.Err() != nil {
		panic(s.Err())
	}
}
