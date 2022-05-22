// Copyright 2022 Michael Kahlke.

package gtab

import (
	"reflect"
	"testing"
)

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

