// Copyright 2022 Michael Kahlke.

package main

const BADSUFFIX string = "nope"
const FINGERS int = 4		// # fingers on one hand, not counting the thumb
const INTERVALS int = 12	// # intervals in an octave
const STRINGS int = 6		// # strings on a guitar

var frets = []int{5,5,5,4,5,5}	// # halfsteps (frets) between strings (including imaginary 7th string)
