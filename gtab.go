// Copyright 2022 Michael Kahlke.

package main

/*
Given a known chord name, display the guitar tab.
TODO:
- support for alternate tabs further up the fretboard
- determine chord name from tab
- support for alternate names taking (musical) keys into account
- support for "add" notes for close matches
*/

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Half Steps    Note    Interval
//     0           1     unison
//     1          (2)    minor 2nd
//     2           2     major 2nd
//     3          (3)    minor 3rd
//     4           3     major 3rd
//     5           4     perfect 4th
//     6          (5)    diminished 5th
//     7           5     perfect 5th
//     8          (6)    minor 6th
//     9           6     major 6th
//    10          (7)    minor 7th
//    11           7     major 7th
//    12           8     perfect octave
//    13          (9)    minor 9th
//    14           9     major 9th, octave + 2nd
//    15         (10)    minor 10th
//    16          10     major 10th, octave + 3rd
//    17          11     major 11th, octave + perfect 4th
//    18         (12)    diminished 12th
//    19          12     major 12th, octave + 5th
//    20         (13)    minor 13th
//    21          13     major 13th, octave + 6th
//    22         (14)    minor 14th
//    23          14     major 14th, octave + 7th
//    24          15     15th, octave + 8th
var chords = map[string][]int {
	// major
	"":		[]int{0,4,7},		// major = 1,3,5
	"sus2":		[]int{0,2,7},		// sustained 2nd = 1,2,5
	"sus":		[]int{0,5,7},		// sustained = 1,4,5
	"sus4":		[]int{0,5,7},		// sustained 4th = 1,4,5
	"6":		[]int{0,4,7,9},		// major 6th = 1,3,5,6
	"6sus2":	[]int{0,2,7,9},		// 6th sustained 2nd = 1,2,5,6
	"6sus":		[]int{0,5,7,9},		// 6th sustained = 1,4,5,6
	"6sus4":	[]int{0,5,7,9},		// 6th sustained 4th = 1,4,5,6
	"maj7":		[]int{0,4,7,11},	// major 7th = 1,3,5,7
	"maj7sus2":	[]int{0,2,7,11},	// major 7th sustained 2nd = 1,2,5,7
	"maj7sus":	[]int{0,5,7,11},	// major 7th sustained = 1,4,5,7
	"maj7sus4":	[]int{0,5,7,11},	// major 7th sustained 4th = 1,4,5,7
	"add9":		[]int{0,4,7,14},	// add 9th = 1,3,5,9
	"maj9":		[]int{0,4,7,11,14},	// major 9th = 1,3,5,7,9
	"6/9":		[]int{0,4,7,9,14},	// major 6/9th = 1,3,5,6,9
	"maj13":	[]int{0,4,7,11,14,21},	// major 13th = 1,3,5,7,9,13
	"maj7/6":	[]int{0,4,7,11,21},	// major 7/6th = 1,3,5,7,13
	"6/9#11":	[]int{0,4,7,9,14,18},	// major 6/9#11 = 1,3,5,6,9,#11
	"maj7#11":	[]int{0,4,7,11,18},	// major 7/#11 = 1,3,5,7,#11
	"maj11":	[]int{0,4,7,11,14,17},	// major 11th = 1,3,5,7,9,11
	"maj9/#11":	[]int{0,4,7,11,14,18},	// major 9/#11 = 1,3,5,7,9,#11
	"add9#11":	[]int{0,4,7,14,18},	// add 9th/#11 = 1,3,5,9,#11
	"add11":	[]int{0,4,7,17},	// add 11th = 1,3,5,11
	"add#11":	[]int{0,4,7,18},	// add #11th = 1,3,5,#11
	"maj7+":	[]int{0,4,8,11},	// major 7th augmented = 1,3,#5,7
	"maj9+":	[]int{0,4,8,11,14},	// major 9th augmented = 1,3,#5,7,9
	// minor
	"m":		[]int{0,3,7},		// minor = 1,b3,5
	"m7":		[]int{0,3,7,10},	// minor 7th = 1,b3,5,b7
	"m7/11":	[]int{0,3,7,10,17},	// minor 7/11th = 1,b3,5,b7,11
	"m9":		[]int{0,3,7,10,14},	// minor 7th = 1,b3,5,b7,9
	"m11":		[]int{0,3,7,10,14,17},	// minor 11th = 1,b3,5,b7,9,11
	"madd9":	[]int{0,3,7,14},	// minor add 9th = 1,b3,5,9
	"m7/13":	[]int{0,3,7,10,21},	// minor 7/13th = 1,b3,5,b7,13
	"m9/13":	[]int{0,3,7,10,14,21},	// minor 9/13th = 1,b3,5,b7,9,13
	"m13":		[]int{0,3,7,10,14,21},	// minor 13th = 1,b3,5,b7,9,13
	"m7/11/13":	[]int{0,3,7,10,17,21},	// minor 7/11/13th = 1,b3,5,b7,11,13
	"m7b5":		[]int{0,3,6,10},	// minor 7th flat 5th = 1,b3,b5,b7
	"m7b5/11":	[]int{0,3,6,10,17},	// minor 7th flat 5/11th = 1,b3,b5,b7,11
	"m7+":		[]int{0,3,8,10},	// minor 7th augmented = 1,b3,#5,b7
	"m7/11+":	[]int{0,3,8,10,17},	// minor 7/11th augmented = 1,b3,#5,b7,11
	"m6":		[]int{0,3,7,9},		// minor 6th = 1,b3,5,6
	"m6/9":		[]int{0,3,7,9,14},	// minor 6/9th = 1,b3,5,6,9
	"m6/7-":	[]int{0,3,7,9,11},	// minor 6/7- = 1,b3,5,6,7
	"m6/9.7-":	[]int{0,3,7,9,11,14},	// minor 6/9/7- = 1,b3,5,6,7,9
	"m6/11":	[]int{0,3,7,9,17},	// minor 6/11th = 1,b3,5,6,11
	"m6/9/11":	[]int{0,3,7,9,14,17},	// minor 6/9/11th = 1,b3,5,6,9,11
	"m6/9/#11":	[]int{0,3,7,9,14,18},	// minor 6/9/#11th = 1,b3,5,6,9,#11
	"m7-":		[]int{0,3,7,11},	// minor 7th - = 1,b3,5,7
	"minmaj7":	[]int{0,3,7,11},	// minor major 7th = 1,b3,5,7
	"m9-":		[]int{0,3,7,11,14},	// minor 9th - = 1,b3,5,7,9
	"minmaj9":	[]int{0,3,7,11,14},	// minor major 9th = 1,b3,5,7,9
	// dominant
	"7":		[]int{0,4,7,10},	// dominant 7th = 1,3,5,b7
	"7/6":		[]int{0,4,7,10,21},	// dominant 7/6th = 1,3,5,b7,13
	"9":		[]int{0,4,7,10,14},	// dominant 9th = 1,3,5,b7,9
	"13":		[]int{0,4,7,10,14,21},	// dominant 13th = 1,3,5,b7,9,13
	"7sus":		[]int{0,5,7,10},	// dominant 7th sustained = 1,4,5,b7
	"7/6sus":	[]int{0,5,7,10,21},	// dominant 7/6th sustained = 1,4,5,b7,13
	"11":		[]int{0,7,10,14,17},	// dominant 11th = 1,5,b7,9,11
	"13sus":	[]int{0,7,10,14,17,21},	// dominant 13th sustained = 1,5,b7,9,11,13
	"13#11":	[]int{0,4,7,10,14,18,21},// dominant 13/#11th = 1,3,5,b7,9,#11,13
	"9#11":		[]int{0,4,7,10,14,18},	// dominant 9/#11th = 1,3,5,b7,9,#11
	"9b5":		[]int{0,4,6,10,14},	// dominant 9b5 = 1,3,b5,b7,9
	"7+":		[]int{0,4,8,10},	// dominant 7th augmented = 1,3,#5,b7
	"7b5":		[]int{0,4,6,10},	// dominant 7b5 = 1,3,b5,b7
	"7#9":		[]int{0,4,7,10,15},	// dominant 7#9 = 1,3,5,b7,#9
	"7b9":		[]int{0,4,7,10,13},	// dominant 7b9 = 1,3,5,b7,b9
	"7#9+":		[]int{0,4,8,10,15},	// dominant 7#9 augmented = 1,3,#5,b7,#9
	"7b9+":		[]int{0,4,8,10,13},	// dominant 7b9 augmented = 1,3,#5,b7,b9
	"7#9b5":	[]int{0,4,6,10,15},	// dominant 7#9b5 = 1,3,b5,b7,#9
	"7b9b5":	[]int{0,4,6,10,13},	// dominant 7b9b5 = 1,3,b5,b7,b9
	"13b9":		[]int{0,4,7,10,13,21},	// dominant 13b9 = 1,3,5,b7,b9,13
	"13b9#11":	[]int{0,4,7,10,13,18,21},// dominant 13b9#11 = 1,3,5,b7,b9,#11,13
	"13#9":		[]int{0,4,7,10,15,21},	// dominant 13b9 = 1,3,5,b7,#9,13
	"9+":		[]int{0,4,8,10,14},	// dominant 9 augmented = 1,3,#5,b7,9
	"+":		[]int{0,4,8},		// augmented = 1,3,#5
	"+11b9":	[]int{0,4,8,10,13,17},	// augmented 11b9 = 1,3,#5,b7,b9,11
	// diminished
	"dim":		[]int{0,3,6},		// diminished triad = 1,b3,b5
	"dim6":		[]int{0,3,6,8},		// diminished 6th = 1,b3,b5,b6
	"dim7":		[]int{0,3,6,9},		// diminished 7th = 1,b3,b5,bb7(6)
	"dim7/7-":	[]int{0,3,6,9,11},	// diminished 7/7th - = 1,b3,b5,6,7
	"dim7/9-":	[]int{0,3,6,9,14},	// diminished 7/9th - = 1,b3,b5,6,9
	"dim9":		[]int{0,3,6,9,13},	// diminished 9th = 1,b3,b5,bb7(6),b9
	"dim11":	[]int{0,6,9,13,16},	// diminished 11th = 1,b5,bb7,b9,b11
	// future: need to handle chords like aug6 (it, fr, de)
}

var debug bool

func main() {
	// Set log prefix, disable time, source file and line number
	log.SetPrefix("gtab: ")
	log.SetFlags(0)

	var chord string
	flag.StringVar(&chord, "chord", "C", "set to chord name, b=flat, #=sharp")
	flag.BoolVar(&debug, "debug", false, "enable debug statements")
	flag.Parse()

	if debug {
		for key, halfsteps := range chords {
			fmt.Printf("Chord suffix: %s, intervals:", key)
			for _, halfstep := range halfsteps {
				fmt.Printf(" %d", halfstep)
			}
			fmt.Println()
		}

		flag.Parse()
	}

	root, suffix := chord_root(chord)
	match(root, suffix)
}

func chord_root(name string) (string, string) {
	var notes = [7]string{"C","D","E","F","G","A","B"}
	base := string(name[0])
	found := false
	for i := range notes {
		if base == notes[i] {
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Chord root must be one of: C,D,E,F,G,A,B with b/# as appropriate\n")
		os.Exit(0)
	}
	if len(name) > 1 {
		var root string = name[0:2]
		if len(root) > 1 && root[1] == '#' || root[1] == 'b' {
			base = root
		}
	}
	if debug { fmt.Printf("Chord name %s, base %s, name[%d:] = %s\n", name, base, len(base), name[len(base):]) }
	return base, name[len(base):]
}

const INTERVALS int = 12

var frets = []int{5,5,5,4,5,5}	// # halfsteps (frets) between strings (including imaginary 7th string)

func match(root string, suffix string) {
	var num_strings = len(frets)
	// Find chord root on fretboard
	root_string := -1
	root_fret := -1
	// low E-string based root names
	var sharp_root_names = [INTERVALS]string{"E","F","F#","G","G#","A","A#","B", "C","C#","D","D#"}
	var flat_root_names = [INTERVALS]string{"E","F","Gb","G","Ab","A","Bb","B","C","Db","D","Eb"}
	cursor := 0
	findroot:
	for gstring := 0; gstring < num_strings; gstring++ {
		// Start with fret "0" on this gstring
		for fret := 0; fret < frets[gstring]; fret++ {
			if root == sharp_root_names[(cursor+fret)%INTERVALS] || root == flat_root_names[(cursor+fret)%INTERVALS] {
				root_string = gstring
				root_fret = fret
				break findroot
			}
		}
		cursor += frets[gstring]
	}
	if root_string == -1 {
		fmt.Printf("Cannot determine root string and fret for bass note of chord name %s%s\n", root, suffix)
		os.Exit(0)
	}

	// Find matching suffix
	halfsteps, exists := chords[suffix]
	if !exists {
		fmt.Printf("Cannot determine tab for chord %s%s\n", root, suffix)
		os.Exit(0)
	}

	if debug {
		var root_strings = []string{"E","A","D","G","B","E"}
		fmt.Printf("Root string: %s, root fret: %d for %s%s, tab progression found: ", root_strings[root_string], root_fret, root, suffix)
		fmt.Printf("halfsteps:")
		for i := 0; i < len(halfsteps); i++ {
			fmt.Printf(" %d", halfsteps[i])
		}
		fmt.Println()
	}

	tab := find_tab(root_string, root_fret, halfsteps, false)
	if playable(tab) { fmt.Printf("Finger %s%s: %s\n", root, suffix, gen_tab(tab)) }
	tab = find_tab(root_string, root_fret, halfsteps, true)
	if playable(tab) { fmt.Printf("Barred %s%s: %s\n", root, suffix, gen_tab(tab)) }

}

const FINGERS int = 4	// on one hand, not counting thumbs

func playable(atab []int) bool {
	var non_open_fret_fingerings = make(map[int]int)
	x_transition := false
	count := 0
	for i := range atab {
		if atab[i] > 0 {
			fret, fingered := non_open_fret_fingerings[atab[i]]
			if !fingered {
				non_open_fret_fingerings[atab[i]] = fret
				count++
			}
		} else {
			if atab[i] > 0 && atab[max(0, i-1)] < 0 {
				x_transition = true
			}
		}
	}
	return !x_transition && count <= FINGERS
}

func gen_tab(atab []int) string {
	tab := ""
	for i := range atab {
		if atab[i] == -1 {
			tab += " x"
		} else {
			tab += " " + strconv.Itoa(atab[i])
		}
	}
	return tab
}

func max(a int, b int) int {
	if (a > b ) {
		return a
	} else {
		return b
	}
}

func find_tab(root_string int, root_fret int, halfsteps []int, bar_chord bool) []int {
	var atab []int

	// X out any strings that precede the root string if applicable
	for i := 0; i < root_string; i++ { atab = append(atab, -1) }

	// Starting at root string and fret, look for notes on successive strings that match chord tones
	current_offset := 0
	prev_halfstep := -1
	for i := root_string; i < len(frets); i++ {
		matched := false
		init_j := 0
		foundmatch:
		for j := init_j; j < frets[i]; j++ {
			for k := range halfsteps {
				if (current_offset+j) % INTERVALS == halfsteps[k] && prev_halfstep != halfsteps[k] {
					if bar_chord {
						atab = append(atab, root_fret+j)
						current_offset += frets[i]
					} else {
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
