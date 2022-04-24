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
	writer   = bufio.NewWriter(os.Stdout)
)

func main() {
	log.SetPrefix(os.Args[0] + ": ")
	log.SetFlags(0)
	flag.Parse()

	if flag.NArg() == 0 {
		vis(os.Stdin)
	} else {
		for _, fname := range flag.Args() {
			f, err := os.Open(fname)
			if err != nil {
				exitCode = 1
				log.Println(err)
				continue
			}
			vis(f)
			f.Close()
		}
	}
	writer.Flush()
	os.Exit(exitCode)
}

func vis(r io.Reader) {
	b := bufio.NewReader(r)
	for {
		r, _, err := b.ReadRune()
		if err == io.EOF {
			return
		}
		if err != nil {
			exitCode = 1
			log.Println(err)
			return
		}
		if unicode.IsGraphic(r) || unicode.IsSpace(r) {
			_, err = writer.WriteRune(r)
		} else {
			_, err = fmt.Fprintf(writer, "\\u%04x", r)
		}
		if err != nil {
			return
		}
	}
}
