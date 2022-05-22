// Copyright 2022 Michael Kahlke.

package gtab

import (
	"strconv"
	"strings"
)

// Order is important! List sharps and flats of a note first
var notes = []string {"Eb","E","F#","F","Gb","G#","G","Ab","A#","A","Bb","B","C#","C","Db","D#","D"}

func ChordComponents(chord_name string) (string, string) {
	root, suffix := "", BADSUFFIX
	for _, note := range notes {
		if strings.HasPrefix(chord_name, note) {
			root = note
			break
		}
	}
	suffix = chord_name[len(root):]
	_, exists := Chords[suffix]
	if !exists {
		suffix = BADSUFFIX
	}
	return root, suffix
}

func GenTab(atab []int) string {
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

func findRoot(root string) (int, int) {
	cursor := 0
	// Find chord root at base of fretboard
	for i := 0; i < STRINGS; i++ {
		// Try halfsteps on this string leading up to the next string
		for fret := 0; fret < frets[i]; fret++ {
			if root == sharps[(cursor+fret)%INTERVALS] || root == flats[(cursor+fret)%INTERVALS] {
				return i, fret
			}
		}
		// Advance cursor to the next string and keep going
		cursor += frets[i]
	}
	return -1, -1
}

func findRootOnString(root string, s String) (int, int) {
	cursor := 0
	for i := 0; i < int(s); i++ {
		cursor += frets[i]
	}
	for i := 0; i < INTERVALS; i++ {
		if root == sharps[(cursor+i)%INTERVALS] || root == flats[(cursor+i)%INTERVALS] {
			return int(s), i
		}
	}
	return -1, -1
}

