package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

var (
	exitCode = 0
)

func main() {
	log.SetPrefix(os.Args[0] + ": ")
	log.SetFlags(0)
	flag.Parse()
	defer func() { os.Exit(exitCode) }()

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	if flag.NArg() == 0 {
		vis(w, os.Stdin)
	} else {
		for _, fname := range flag.Args() {
			f, err := os.Open(fname)
			if err != nil {
				log.Print(err)
				exitCode = 1
				continue
			}
			vis(w, f)
			f.Close()
		}
	}
}

func vis(w io.Writer, r io.Reader) {
	b := bufio.NewReader(r)
	for {
		r, _, err := b.ReadRune()
		if err != nil {
			return
		}

		if unicode.IsGraphic(r) || unicode.IsSpace(r) {
			fmt.Fprintf(w, "%c", r)
		} else {
			fmt.Fprintf(w, "\\u%04x", r)
		}
	}
}
