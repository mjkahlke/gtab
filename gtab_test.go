// Copyright 2022 Michael Kahlke.

package gtab

import (
	"os"
	"os/exec"
	"reflect"
	"testing"
)

func TestChordRootBogusRoot(t *testing.T) {
	if os.Getenv("EXIT_TESTER") == "1" {
		name := "Xmaj7"
		chord_root(name)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestChordRootBogusRoot")
	cmd.Env = append(os.Environ(), "EXIT_TESTER=1")
	err := cmd.Run()
	if err == nil {
		return
	}
	//if e, ok := err.(*exec.ExitError); ok && !e.Success() {
	t.Fatalf("Process ran with error %v, exit status 0", err)
}

func TestMatchNonexistentRoot(t *testing.T) {
	if os.Getenv("EXIT_TESTER") == "1" {
		match("E#", "7b5")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestMatchNonexistentRoot")
	cmd.Env = append(os.Environ(), "EXIT_TESTER=1")
	err := cmd.Run()
	if err == nil {
		return
	}
	//if e, ok := err.(*exec.ExitError); ok && !e.Success() {
	t.Fatalf("Process ran with error %v, exit status 0", err)
}

func TestUnsupportedSuffix(t *testing.T) {
	if os.Getenv("EXIT_TESTER") == "1" {
		match("A", "aug6")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestUnsupportedSuffix")
	cmd.Env = append(os.Environ(), "EXIT_TESTER=1")
	err := cmd.Run()
	if err == nil {
		return
	}
	//if e, ok := err.(*exec.ExitError); ok && !e.Success() {
	t.Fatalf("Process ran with error %v, exit status 0", err)
}

func TestCmajorTabFingered(t *testing.T) {
	tab := find_tab(1, 3, []int{0,4,7}, false)
	answer := []int{-1,3,2,0,1,0}
	if reflect.DeepEqual(tab, answer) { return }
	t.Fatalf("C major (fingered) was not correctly produced")
}

func TestCmajorTabBarred(t *testing.T) {
	tab := find_tab(1, 3, []int{0,4,7}, true)
	answer := []int{-1,3,5,5,5,3}
	if reflect.DeepEqual(tab, answer) { return }
	t.Fatalf("C major (barred) was not correctly produced")
}

func TestPlayableYes(t *testing.T) {
	ok := []int{-1,3,2,0,1,0}
	if playable(ok) { return }
	t.Fatalf("C major is certainly playable")
}

func TestPlayableNo(t *testing.T) {
	extra_finger := []int{-1,5,4,3,2,1}
	if !playable(extra_finger) { return }
	t.Fatalf("Hello, my name is Inigo Montoya. Prepare to die!")
}

func TestGenTabCmaj7(t *testing.T) {
	cmaj7 := []int{-1,3,2,0,0,0}
	if gen_tab(cmaj7) == " x 3 2 0 0 0" { return }
	t.Fatalf("Error generating proper tab for Cmajy chord")
}

