package elementaryalignment_test

import (
	"impl-using-go/aoj/elementaryalignment"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInsertionSort(t *testing.T) {
	t.Parallel()

	a := []int{5, 2, 4, 6, 1, 3}
	elementaryalignment.InsertionSort(a)
	want := []int{1, 2, 3, 4, 5, 6}
	if diff := cmp.Diff(want, a); diff != "" {
		t.Error(diff)
	}
}

func TestBubbleSort(t *testing.T) {
	t.Parallel()

	a := []int{5, 3, 2, 4, 1}
	elementaryalignment.BubbleSort(a)
	want := []int{1, 2, 3, 4, 5}

	if diff := cmp.Diff(want, a); diff != "" {
		t.Error(diff)
	}
}

func TestSelectionSort(t *testing.T) {
	t.Parallel()

	a := []int{5, 6, 4, 2, 1, 3}
	elementaryalignment.SelectionSort(a)
	want := []int{1, 2, 3, 4, 5, 6}

	if diff := cmp.Diff(want, a); diff != "" {
		t.Errorf(diff)
	}
}

func TestShellSort(t *testing.T) {
	t.Parallel()

	arr := []int{5, 1, 4, 3, 2, 7, 9, 8, 6, 10}
	gs := []int{4, 1}
	elementaryalignment.ShellSort(arr, gs)

	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if diff := cmp.Diff(want, arr); diff != "" {
		t.Error(diff)
	}
}
