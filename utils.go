// Copyright 2022 Michael Kahlke.

package main

import (
	"strconv"
	"strings"
)

// Order is important! List sharps and flats of a note first
var notes = []string {"Eb","E","F#","F","Gb","G#","G","Ab","A#","A","Bb","B","C#","C","Db","D#","D"}

func chordComponents(chord_name string) (string, string) {
	root, suffix := "", BADSUFFIX
	for _, note := range notes {
		if strings.HasPrefix(chord_name, note) {
			root = note
			break
		}
	}
	suffix = chord_name[len(root):]
	_, exists := chords[suffix]
	if !exists {
		suffix = BADSUFFIX
	}
	return root, suffix
}

func genTab(atab []int) string {
	tab := ""
	for _, fret := range atab {
		if fret == -1 {
			tab += " x"
		} else {
			tab += " " + strconv.Itoa(fret)
		}
	}
	return tab
}

func playable(atab []int) bool {
	var non_open_fret_fingerings = make(map[int]int)
	count := 0
	for i, fret := range atab {
		if fret > 0 {
			pos, fingered := non_open_fret_fingerings[fret]
			if !fingered {
				non_open_fret_fingerings[atab[i]] = pos
				count++
			}
		}
	}
	return count <= FINGERS
}

// low E-string root names
var sharps_E = [INTERVALS]string{"E","F","F#","G","G#","A","A#","B", "C","C#","D","D#"}
var flats_E = [INTERVALS]string{"E","F","Gb","G","Ab","A","Bb","B","C","Db","D","Eb"}

func findRoot(root string) (int, int) {
	cursor := 0
	// Find chord root at base of fretboard
	for i := 0; i < STRINGS; i++ {
		// Start with fret "0" on this string
		for fret := 0; fret < frets[i]; fret++ {
			if root == sharps_E[(cursor+fret)%INTERVALS] || root == flats_E[(cursor+fret)%INTERVALS] {
				return i, fret
			}
		}
		cursor += frets[i]
	}
	return -1, -1
}

func findRootE(root string) (int, int) {
	for f, n := range sharps_E {
		if root == n {
			return 0, f
		}
	}
	for f, n := range flats_E {
		if root == n {
			return 0, f
		}
	}
	return -1, -1
}

// A-string based root names
var sharps_A = [INTERVALS]string{"A","A#","B","C","C#","D","D#","E","F","F#","G","G#"}
var flats_A = [INTERVALS]string{"A","Bb","B","C","Db","D","Eb","E","F","Gb","G","Ab"}

func findRootA(root string) (int, int) {
	for f, n := range sharps_A {
		if root == n {
			return 1, f
		}
	}
	for f, n := range flats_A {
		if root == n {
			return 1, f
		}
	}
	return -1, -1
}

// D-string based root names
var sharps_D = [INTERVALS]string{"D","D#","E","F","F#","G","G#","A","A#","B","C","C#"}
var flats_D = [INTERVALS]string{"D","Eb","E","F","Gb","G","Ab","A","Bb","B","C","Db"}

func findRootD(root string) (int, int) {
	for f, n := range sharps_D {
		if root == n {
			return 2, f
		}
	}
	for f, n := range flats_D {
		if root == n {
			return 2, f
		}
	}
	return -1, -1
}

// G-string based root names
var sharps_G = [INTERVALS]string{"G","G#","A","A#","B","C","C#","D","D#","E","F","F#"}
var flats_G = [INTERVALS]string{"G","Ab","A","Bb","B","C","Db","D","Eb","E","F","Gb"}

func findRootG(root string) (int, int) {
	for f, n := range sharps_G {
		if root == n {
			return 3, f
		}
	}
	for f, n := range flats_G {
		if root == n {
			return 3, f
		}
	}
	return -1, -1
}

