// Copyright 2022 Michael Kahlke.

package main

import (
	"os"
	"os/exec"
	"reflect"
	"testing"
)

func TestChordRootBogusRoot(t *testing.T) {
	if os.Getenv("EXIT_TESTER") == "1" {
		os.Args = []string{"gtab", "-chord", "Xmaj7"}
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestChordRootBogusRoot")
	cmd.Env = append(os.Environ(), "EXIT_TESTER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && e.Success() {
		t.Fatalf("Xmaj7 does not have a valid chord root")
	}
}

func TestMatchNonexistentRoot(t *testing.T) {
	if os.Getenv("EXIT_TESTER") == "1" {
		os.Args = []string{"gtab", "-chord", "E#7b5"}
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestMatchNonexistentRoot")
	cmd.Env = append(os.Environ(), "EXIT_TESTER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && e.Success() {
		t.Fatalf("E#7b5 has a non-existent root")
	}
}

func TestUnsupportedSuffix(t *testing.T) {
	if os.Getenv("EXIT_TESTER") == "1" {
		os.Args = []string{"gtab", "-chord", "Aaug6"}
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestUnsupportedSuffix")
	cmd.Env = append(os.Environ(), "EXIT_TESTER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && e.Success() {
		t.Fatalf("Aaug6 has an unsupported chord suffix")
	}
}

func TestCmajorTabFingered(t *testing.T) {
	tab := findTab(1, 3, []int{0,4,7}, false)
	answer := []int{-1,3,2,0,1,0}
	if reflect.DeepEqual(tab, answer) { return }
	t.Fatalf("C major (fingered) was not correctly produced")
}

func TestCmajorTabBarred(t *testing.T) {
	tab := findTab(1, 3, []int{0,4,7}, true)
	answer := []int{-1,3,5,5,5,3}
	if reflect.DeepEqual(tab, answer) { return }
	t.Fatalf("C major (barred) was not correctly produced")
}

func TestCmajorTabRootE(t *testing.T) {
	tab := findTab(0, 8, []int{0,4,7}, true)
	answer := []int{8,10,10,9,8,8}
	if reflect.DeepEqual(tab, answer) { return }
	t.Fatalf("C major (root E) was not correctly produced")
}

func TestAflatMajor7RootA(t *testing.T) {
	tab := findTab(1, 11, []int{0,4,7,11}, true)
	answer := []int{-1,11,13,12,13,11}
	if reflect.DeepEqual(tab, answer) { return }
	t.Fatalf("Abmaj7 (root A) was not correctly produced")
}

