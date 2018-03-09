package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	nPtr := flag.Bool("n", false, "number output lines")
	bPtr := flag.Bool("b", false, "number non-blank output lines")
	EPtr := flag.Bool("E", false, "display end-of-line characters as $")
	sPtr := flag.Bool("s", false, "squeeze multiple adjacent blank lines")
	TPtr := flag.Bool("T", false, "display tabs as ^")
	flag.Parse()

	var lastChar rune
	var prefix, suffix string
	newLineCounter := 0

	for _, path := range flag.Args() {
		lineNum := 1
		// check if file exist
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("File does not exist.")
			os.Exit(-1)
		}
		// open file and make it an iterable obj
		file, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		reader := bufio.NewReader(file)

		for {
			prefix = ""
			suffix = ""
			if c, _, err := reader.ReadRune(); err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Println("bad char")
					os.Exit(-1)
				}
			} else {
				if *EPtr {
					if c == '\n' {
						prefix += "$"
					}
				}

				if *sPtr {
					if c == '\n' {
						newLineCounter++
						if newLineCounter > 1 {
							continue
						}
					} else {
						if newLineCounter > 1 {
							prefix += "\n"
							newLineCounter = 0
						} else if newLineCounter == 1 {
							newLineCounter = 0
						}
					}
				}

				if *nPtr {
					if lastChar == 0 || lastChar == '\n' {
						prefix += fmt.Sprintf("%d. ", lineNum)
						lineNum++
					}
				}

				if *bPtr {
					if lastChar == 0 || (c != '\n' && lastChar == '\n') {
						prefix += fmt.Sprintf("%d. ", lineNum)
						lineNum++
					}
				}

				if *TPtr {
					if c == '\t' {
						c = '^'
					}
				}
				fmt.Printf("%s%c%s", prefix, c, suffix)
				lastChar = c
			}
		}

	}
	if lastChar != '\n' {
		fmt.Println()
	}
}
