// Copyright 2022 Michael Kahlke.

package main

// Given a known chord name, display the guitar tab.

import (
	"flag"
	"fmt"
	"log"
	"sort"
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
		for key, halfsteps := range chords {
			fmt.Printf("Chord suffix: %s, intervals:", key)
			for _, halfstep := range halfsteps {
				fmt.Printf(" %d", halfstep)
			}
			fmt.Println()
		}
	}

	root, suffix := chordComponents(chord)
	if root == "" { log.Fatal("Invalid chord name specified") }
	if suffix == BADSUFFIX { log.Fatal("Unsupported chord suffix specified") }
	// Don't need error checking at this point
	match(root, suffix)
}

func match(root string, suffix string) {
	halfsteps := chords[suffix]
	var tabs = make(map[string][]int)

	// Fingered chord
	root_string, root_fret := findRoot(root)
	tab := findTab(root_string, root_fret, halfsteps, false)
	if playable(tab) { tabs[genTab(tab)] = tab }

	// Bar chord
	tab = findTab(root_string, root_fret, halfsteps, true)
	if playable(tab) { tabs[genTab(tab)] = tab }

	// Root on low E string
	root_string, root_fret = findRootOnString(root, E)
	tab = findTab(root_string, root_fret, halfsteps, true)
	if playable(tab) { tabs[genTab(tab)] = tab }

	// Root on A string
	root_string, root_fret = findRootOnString(root, A)
	tab = findTab(root_string, root_fret, halfsteps, true)
	if playable(tab) { tabs[genTab(tab)] = tab }

	// Root on D string
	root_string, root_fret = findRootOnString(root, D)
	tab = findTab(root_string, root_fret, halfsteps, true)
	if playable(tab) { tabs[genTab(tab)] = tab }

	// Root on G string
	root_string, root_fret = findRootOnString(root, G)
	tab = findTab(root_string, root_fret, halfsteps, true)
	if playable(tab) { tabs[genTab(tab)] = tab }

	// Sort and display tabs
	keys := make([]string, 0, len(tabs))
	for t := range tabs {
		keys = append(keys, t)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s\n", k)
	}
}

func findTab(root_string int, root_fret int, halfsteps []int, forefinger_on_root bool) []int {
	var atab []int

	// X out any strings that precede the root string if applicable
	for i := 0; i < root_string; i++ { atab = append(atab, -1) }

	// Starting at root string and fret, look for notes on successive strings that match chord tones
	current_offset, prev_halfstep := 0, -1
	for i := root_string; i < len(frets); i++ {
		matched := false
		foundmatch:
		for j := 0; j < frets[i]; j++ {
			for k := range halfsteps {
				// TODO: 9th/11th/13th chords exceed an octave, need a fix here
				if (current_offset+j) % INTERVALS == halfsteps[k] && prev_halfstep != halfsteps[k] {
					// This usually results in generating a tab for bar chords
					if forefinger_on_root {
						atab = append(atab, root_fret+j)
						current_offset += frets[i]
					} else {
						// Root note isn't necessarily on the zeroeth fret of the root string
						if i == root_string {
							current_offset += frets[i] - root_fret
							atab = append(atab, root_fret+j)
						} else {
							current_offset += frets[i]
							atab = append(atab, j)
						}
					}
					matched = true
					prev_halfstep = halfsteps[k]
					break foundmatch
				}
			}
		}
		if !matched { atab = append(atab, -1) }
	}
	return atab
}

