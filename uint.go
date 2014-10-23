// Copyright (c) 2014 Datacratic. All rights reserved.

package set

import (
	"bytes"
	"sort"
	"strconv"
)

// Uint represents a set of uint64s. Must first be initialized using the NewUint
// function before being used.
type Uint map[uint64]struct{}

// NewUint creates a new uint64 set from the given list of uint64s.
func NewUint(values ...uint64) Uint {
	set := make(map[uint64]struct{})

	for _, value := range values {
		set[value] = struct{}{}
	}

	return set
}

// Test returns true if the given uint64 is in the set.
func (set Uint) Test(values ...uint64) bool {
	for _, value := range values {
		if _, ok := set[value]; !ok {
			return false
		}
	}
	return true
}

// Put adds the given uint64 to the set and is a noop if the uint64 is already in
// the set.
func (set Uint) Put(values ...uint64) {
	for _, value := range values {
		set[value] = struct{}{}
	}
}

// Del removes the given uint64 from the set and is a noop if the uint64 is not
// already in the set.
func (set Uint) Del(values ...uint64) {
	for _, value := range values {
		delete(set, value)
	}
}

// Copy returns a copy of the set.
func (set Uint) Copy() Uint {
	result := NewUint()
	for value := range set {
		result[value] = struct{}{}
	}
	return result
}

// Merge adds all the values of other into the set.
func (set Uint) Merge(other Uint) {
	for value := range other {
		set[value] = struct{}{}
	}
}

// Union returns a set consisting of all values present in both other and the
// current set.
func (set Uint) Union(other Uint) Uint {
	result := set.Copy()
	result.Merge(other)
	return result
}

// Uintersect returns a set consisting of the values present in both sets.
func (set Uint) Uintersect(other Uint) Uint {
	var big, small Uint
	if len(set) < len(other) {
		big = other
		small = set
	} else {
		big = set
		small = other
	}

	result := NewUint()
	for value := range small {
		if _, ok := big[value]; ok {
			result[value] = struct{}{}
		}
	}

	return result
}

// Difference returns a set consisting of all the values in this set that are
// not present in the other set.
func (set Uint) Difference(other Uint) Uint {
	result := set.Copy()

	for value := range other {
		delete(result, value)
	}

	return result
}

type uintArray []uint64

func (array uintArray) Len() int           { return len(array) }
func (array uintArray) Swap(i, j int)      { array[i], array[j] = array[j], array[i] }
func (array uintArray) Less(i, j int) bool { return array[i] < array[j] }

// Array returns a sorted array of all the values in the current set.
func (set Uint) Array() (list []uint64) {
	for value := range set {
		list = append(list, value)
	}
	sort.Sort(uintArray(list))
	return
}

// String returns a uint64 representation of the set suitable for debugging.
func (set Uint) String() string {
	if len(set) == 0 {
		return "{}"
	}

	buffer := new(bytes.Buffer)
	buffer.WriteString("{ ")

	for _, value := range set.Array() {
		buffer.WriteString(strconv.FormatUint(value, 10))
		buffer.WriteString(" ")
	}

	buffer.WriteString("}")

	return buffer.String()
}
