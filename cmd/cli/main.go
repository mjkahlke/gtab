// Copyright 2022 Michael Kahlke.

package main

// Given a known chord name, display the guitar tab.

import (
	"flag"
	"fmt"
	"log"
	gtab "gtab/impl"
)

func main() {
	// Set log prefix, disable time, source file and line number
	log.SetPrefix("gtab: ")
	log.SetFlags(0)

	var chord string
	var list bool
	flag.StringVar(&chord, "chord", "C", "set to chord name, b=flat, #=sharp")
	flag.BoolVar(&list, "list", false, "list supported chord formulas")
	flag.Parse()

	if list {
		for key, halfsteps := range gtab.Chords {
			fmt.Printf("Chord suffix: %s, intervals:", key)
			for _, halfstep := range halfsteps {
				fmt.Printf(" %d", halfstep)
			}
			fmt.Println()
		}
	}

	root, suffix := gtab.ChordComponents(chord)
	if root == "" { log.Fatal("Invalid chord name specified") }
	if suffix == gtab.BADSUFFIX { log.Fatal("Unsupported chord suffix specified") }
	// Don't need error checking at this point
	tabs := gtab.Match(root, suffix)
	for _, t := range tabs {
		fmt.Printf("%s\n", t)
	}
}

