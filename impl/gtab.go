// Copyright 2022 Michael Kahlke.

package gtab

// Given a known chord name, display the guitar tab.

import (
	"sort"
)

func Match(root string, suffix string) []string {
	halfsteps := Chords[suffix]
	var tabs = make(map[string][]int)

	// Fingered chord
	root_string, root_fret := findRoot(root)
	tab := findTab(root_string, root_fret, halfsteps, false)
	if playable(tab) { tabs[GenTab(tab)] = tab }

	// Bar chord
	tab = findTab(root_string, root_fret, halfsteps, true)
	if playable(tab) { tabs[GenTab(tab)] = tab }

	// Root on low E string
	root_string, root_fret = findRootOnString(root, E)
	tab = findTab(root_string, root_fret, halfsteps, true)
	if playable(tab) { tabs[GenTab(tab)] = tab }

	// Root on A string
	root_string, root_fret = findRootOnString(root, A)
	tab = findTab(root_string, root_fret, halfsteps, true)
	if playable(tab) { tabs[GenTab(tab)] = tab }

	// Root on D string
	root_string, root_fret = findRootOnString(root, D)
	tab = findTab(root_string, root_fret, halfsteps, true)
	if playable(tab) { tabs[GenTab(tab)] = tab }

	// Root on G string
	root_string, root_fret = findRootOnString(root, G)
	tab = findTab(root_string, root_fret, halfsteps, true)
	if playable(tab) { tabs[GenTab(tab)] = tab }

	// Sort and display tabs
	keys := make([]string, 0, len(tabs))
	for t := range tabs {
		keys = append(keys, t)
	}
	sort.Strings(keys)
	return keys
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

