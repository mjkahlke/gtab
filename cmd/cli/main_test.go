// Copyright 2022 Michael Kahlke.

package main

import (
	"os"
	"os/exec"
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

