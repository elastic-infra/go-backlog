// Generated by: gen
// TypeWriter: slice
// Directive: +gen on *Category

package gobacklog

import (
	"errors"
	"math/rand"
)

// Sort implementation is a modification of http://golang.org/pkg/sort/#Sort
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at http://golang.org/LICENSE.

// CategorySlice is a slice of type *Category. Use it where you would use []*Category.
type CategorySlice []*Category

// Where returns a new CategorySlice whose elements return true for func. See: http://clipperhouse.github.io/gen/#Where
func (rcv CategorySlice) Where(fn func(*Category) bool) (result CategorySlice) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Count gives the number elements of CategorySlice that return true for the passed func. See: http://clipperhouse.github.io/gen/#Count
func (rcv CategorySlice) Count(fn func(*Category) bool) (result int) {
	for _, v := range rcv {
		if fn(v) {
			result++
		}
	}
	return
}

// SortBy returns a new ordered CategorySlice, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv CategorySlice) SortBy(less func(*Category, *Category) bool) CategorySlice {
	result := make(CategorySlice, len(rcv))
	copy(result, rcv)
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := len(result)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSortCategorySlice(result, less, 0, n, maxDepth)
	return result
}

// GroupByString groups elements into a map keyed by string. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv CategorySlice) GroupByString(fn func(*Category) string) map[string]CategorySlice {
	result := make(map[string]CategorySlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// GroupByInt groups elements into a map keyed by int. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv CategorySlice) GroupByInt(fn func(*Category) int) map[int]CategorySlice {
	result := make(map[int]CategorySlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// GroupByBool groups elements into a map keyed by bool. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv CategorySlice) GroupByBool(fn func(*Category) bool) map[bool]CategorySlice {
	result := make(map[bool]CategorySlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// First returns the first element that returns true for the passed func. Returns error if no elements return true. See: http://clipperhouse.github.io/gen/#First
func (rcv CategorySlice) First(fn func(*Category) bool) (result *Category, err error) {
	for _, v := range rcv {
		if fn(v) {
			result = v
			return
		}
	}
	err = errors.New("no CategorySlice elements return true for passed func")
	return
}

// MaxBy returns an element of CategorySlice containing the maximum value, when compared to other elements using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#MaxBy
func (rcv CategorySlice) MaxBy(less func(*Category, *Category) bool) (result *Category, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine the MaxBy of an empty slice")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if rcv[i] != rcv[m] && !less(rcv[i], rcv[m]) {
			m = i
		}
	}
	result = rcv[m]
	return
}

// MinBy returns an element of CategorySlice containing the minimum value, when compared to other elements using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#MinBy
func (rcv CategorySlice) MinBy(less func(*Category, *Category) bool) (result *Category, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine the Min of an empty slice")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if less(rcv[i], rcv[m]) {
			m = i
		}
	}
	result = rcv[m]
	return
}

// Distinct returns a new CategorySlice whose elements are unique. See: http://clipperhouse.github.io/gen/#Distinct
func (rcv CategorySlice) Distinct() (result CategorySlice) {
	appended := make(map[*Category]bool)
	for _, v := range rcv {
		if !appended[v] {
			result = append(result, v)
			appended[v] = true
		}
	}
	return result
}

// DistinctBy returns a new CategorySlice whose elements are unique, where equality is defined by a passed func. See: http://clipperhouse.github.io/gen/#DistinctBy
func (rcv CategorySlice) DistinctBy(equal func(*Category, *Category) bool) (result CategorySlice) {
Outer:
	for _, v := range rcv {
		for _, r := range result {
			if equal(v, r) {
				continue Outer
			}
		}
		result = append(result, v)
	}
	return result
}

// Shuffle returns a shuffled copy of CategorySlice, using a version of the Fisher-Yates shuffle. See: http://clipperhouse.github.io/gen/#Shuffle
func (rcv CategorySlice) Shuffle() CategorySlice {
	numItems := len(rcv)
	result := make(CategorySlice, numItems)
	copy(result, rcv)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result[r], result[i] = result[i], result[r]
	}
	return result
}

// Sort implementation based on http://golang.org/pkg/sort/#Sort, see top of this file

func swapCategorySlice(rcv CategorySlice, a, b int) {
	rcv[a], rcv[b] = rcv[b], rcv[a]
}

// Insertion sort
func insertionSortCategorySlice(rcv CategorySlice, less func(*Category, *Category) bool, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(rcv[j], rcv[j-1]); j-- {
			swapCategorySlice(rcv, j, j-1)
		}
	}
}

// siftDown implements the heap property on rcv[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDownCategorySlice(rcv CategorySlice, less func(*Category, *Category) bool, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(rcv[first+child], rcv[first+child+1]) {
			child++
		}
		if !less(rcv[first+root], rcv[first+child]) {
			return
		}
		swapCategorySlice(rcv, first+root, first+child)
		root = child
	}
}

func heapSortCategorySlice(rcv CategorySlice, less func(*Category, *Category) bool, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDownCategorySlice(rcv, less, i, hi, first)
	}

	// Pop elements, largest first, into end of rcv.
	for i := hi - 1; i >= 0; i-- {
		swapCategorySlice(rcv, first, first+i)
		siftDownCategorySlice(rcv, less, lo, i, first)
	}
}

// Quicksort, following Bentley and McIlroy,
// Engineering a Sort Function, SP&E November 1993.

// medianOfThree moves the median of the three values rcv[a], rcv[b], rcv[c] into rcv[a].
func medianOfThreeCategorySlice(rcv CategorySlice, less func(*Category, *Category) bool, a, b, c int) {
	m0 := b
	m1 := a
	m2 := c
	// bubble sort on 3 elements
	if less(rcv[m1], rcv[m0]) {
		swapCategorySlice(rcv, m1, m0)
	}
	if less(rcv[m2], rcv[m1]) {
		swapCategorySlice(rcv, m2, m1)
	}
	if less(rcv[m1], rcv[m0]) {
		swapCategorySlice(rcv, m1, m0)
	}
	// now rcv[m0] <= rcv[m1] <= rcv[m2]
}

func swapRangeCategorySlice(rcv CategorySlice, a, b, n int) {
	for i := 0; i < n; i++ {
		swapCategorySlice(rcv, a+i, b+i)
	}
}

func doPivotCategorySlice(rcv CategorySlice, less func(*Category, *Category) bool, lo, hi int) (midlo, midhi int) {
	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's Ninther, median of three medians of three.
		s := (hi - lo) / 8
		medianOfThreeCategorySlice(rcv, less, lo, lo+s, lo+2*s)
		medianOfThreeCategorySlice(rcv, less, m, m-s, m+s)
		medianOfThreeCategorySlice(rcv, less, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThreeCategorySlice(rcv, less, lo, m, hi-1)

	// Invariants are:
	//	rcv[lo] = pivot (set up by ChoosePivot)
	//	rcv[lo <= i < a] = pivot
	//	rcv[a <= i < b] < pivot
	//	rcv[b <= i < c] is unexamined
	//	rcv[c <= i < d] > pivot
	//	rcv[d <= i < hi] = pivot
	//
	// Once b meets c, can swap the "= pivot" sections
	// into the middle of the slice.
	pivot := lo
	a, b, c, d := lo+1, lo+1, hi, hi
	for {
		for b < c {
			if less(rcv[b], rcv[pivot]) { // rcv[b] < pivot
				b++
			} else if !less(rcv[pivot], rcv[b]) { // rcv[b] = pivot
				swapCategorySlice(rcv, a, b)
				a++
				b++
			} else {
				break
			}
		}
		for b < c {
			if less(rcv[pivot], rcv[c-1]) { // rcv[c-1] > pivot
				c--
			} else if !less(rcv[c-1], rcv[pivot]) { // rcv[c-1] = pivot
				swapCategorySlice(rcv, c-1, d-1)
				c--
				d--
			} else {
				break
			}
		}
		if b >= c {
			break
		}
		// rcv[b] > pivot; rcv[c-1] < pivot
		swapCategorySlice(rcv, b, c-1)
		b++
		c--
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := min(b-a, a-lo)
	swapRangeCategorySlice(rcv, lo, b-n, n)

	n = min(hi-d, d-c)
	swapRangeCategorySlice(rcv, c, hi-n, n)

	return lo + b - a, hi - (d - c)
}

func quickSortCategorySlice(rcv CategorySlice, less func(*Category, *Category) bool, a, b, maxDepth int) {
	for b-a > 7 {
		if maxDepth == 0 {
			heapSortCategorySlice(rcv, less, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivotCategorySlice(rcv, less, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSortCategorySlice(rcv, less, a, mlo, maxDepth)
			a = mhi // i.e., quickSortCategorySlice(rcv, mhi, b)
		} else {
			quickSortCategorySlice(rcv, less, mhi, b, maxDepth)
			b = mlo // i.e., quickSortCategorySlice(rcv, a, mlo)
		}
	}
	if b-a > 1 {
		insertionSortCategorySlice(rcv, less, a, b)
	}
}