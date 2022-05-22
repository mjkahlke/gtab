// Copyright 2022 Michael Kahlke.

package gtab

import (
	"strconv"
	"testing"
)

func TestChordComponentsBogusRoot(t *testing.T) {
	root, _ := ChordComponents("Xmaj7")
	if root == "" { return }
	t.Fatalf("Bogus chord name Xmaj7 was not caught")
}

func TestChordComponentsUnsupportedSuffix(t *testing.T) {
	_, suffix := ChordComponents("Aaug6")
	if suffix == BADSUFFIX { return }
	t.Fatalf("Unsupported chord suffix aug6 was not caught")
}

func TestGenTabCmaj7(t *testing.T) {
	cmaj7 := []int{-1,3,2,0,0,0}
	if GenTab(cmaj7) == " x 3 2 0 0 0" { return }
	t.Fatalf("Error generating proper tab for Cmaj7 chord")
}

func TestPlayableYes(t *testing.T) {
	ok := []int{-1,3,2,0,1,0}
	if playable(ok) { return }
	t.Fatalf("I do not mean to pry, but you don't by any chance happen to have six fingers on your right hand?")
}

func TestPlayableNo(t *testing.T) {
	extra_finger := []int{-1,5,4,3,2,1}
	if !playable(extra_finger) { return }
	t.Fatalf("Hello, my name is Inigo Montoya. You killed my father. Prepare to die!")
}

func TestFindRoot(t *testing.T) {
	str, fret := findRoot("C")
	if str != 1 { t.Fatalf("Incorrect string index "+strconv.Itoa(str)+" returned instead of 1 for C") }
	if fret != 3 { t.Fatalf("Incorrect fret index "+strconv.Itoa(fret)+" returned instead of 3 for C") }
}

func TestFindRootE(t *testing.T) {
	str, fret := findRootOnString("C", E)
	if str != 0 { t.Fatalf("Incorrect string index "+strconv.Itoa(str)+" returned instead of 0 for C") }
	if fret != 8 { t.Fatalf("Incorrect fret index "+strconv.Itoa(fret)+" returned instead of 8 for C") }
}

func TestFindRootA(t *testing.T) {
	str, fret := findRootOnString("C", A)
	if str != 1 { t.Fatalf("Incorrect string index "+strconv.Itoa(str)+" returned instead of 1 for C") }
	if fret != 3 { t.Fatalf("Incorrect fret index "+strconv.Itoa(fret)+" returned instead of 3 for C") }
}

func TestFindRootD(t *testing.T) {
	str, fret := findRootOnString("C", D)
	if str != 2 { t.Fatalf("Incorrect string index "+strconv.Itoa(str)+" returned instead of 2 for C") }
	if fret != 10 { t.Fatalf("Incorrect fret index "+strconv.Itoa(fret)+" returned instead of 10 for C") }
}

func TestFindRootG(t *testing.T) {
	str, fret := findRootOnString("C", G)
	if str != 3 { t.Fatalf("Incorrect string index "+strconv.Itoa(str)+" returned instead of 3 for C") }
	if fret != 5 { t.Fatalf("Incorrect fret index "+strconv.Itoa(fret)+" returned instead of 5 for C") }
}

