// Copyright (c) 2014 Datacratic. All rights reserved.

package set

import (
	"bytes"
	"sort"
	"strconv"
)

// Int represents a set of int64s. Must first be initialized using the NewInt
// function before being used.
type Int map[int64]struct{}

// NewInt creates a new int64 set from the given list of int64s.
func NewInt(values ...int64) Int {
	set := make(map[int64]struct{})

	for _, value := range values {
		set[value] = struct{}{}
	}

	return set
}

// Test returns true if the given int64 is in the set.
func (set Int) Test(values ...int64) bool {
	for _, value := range values {
		if _, ok := set[value]; !ok {
			return false
		}
	}
	return true
}

// Put adds the given int64 to the set and is a noop if the int64 is already in
// the set.
func (set Int) Put(values ...int64) {
	for _, value := range values {
		set[value] = struct{}{}
	}
}

// Del removes the given int64 from the set and is a noop if the int64 is not
// already in the set.
func (set Int) Del(values ...int64) {
	for _, value := range values {
		delete(set, value)
	}
}

// Copy returns a copy of the set.
func (set Int) Copy() Int {
	result := NewInt()
	for value := range set {
		result[value] = struct{}{}
	}
	return result
}

// Merge adds all the values of other into the set.
func (set Int) Merge(other Int) {
	for value := range other {
		set[value] = struct{}{}
	}
}

// Union returns a set consisting of all values present in both other and the
// current set.
func (set Int) Union(other Int) Int {
	result := set.Copy()
	result.Merge(other)
	return result
}

// Intersect returns a set consisting of the values present in both sets.
func (set Int) Intersect(other Int) Int {
	var big, small Int
	if len(set) < len(other) {
		big = other
		small = set
	} else {
		big = set
		small = other
	}

	result := NewInt()
	for value := range small {
		if _, ok := big[value]; ok {
			result[value] = struct{}{}
		}
	}

	return result
}

// Difference returns a set consisting of all the values in this set that are
// not present in the other set.
func (set Int) Difference(other Int) Int {
	result := set.Copy()

	for value := range other {
		delete(result, value)
	}

	return result
}

type intArray []int64

func (array intArray) Len() int           { return len(array) }
func (array intArray) Swap(i, j int)      { array[i], array[j] = array[j], array[i] }
func (array intArray) Less(i, j int) bool { return array[i] < array[j] }

// Array returns a sorted array of all the values in the current set.
func (set Int) Array() (list []int64) {
	for value := range set {
		list = append(list, value)
	}
	sort.Sort(intArray(list))
	return
}

// String returns a int64 representation of the set suitable for debugging.
func (set Int) String() string {
	if len(set) == 0 {
		return "{}"
	}

	buffer := new(bytes.Buffer)
	buffer.WriteString("{ ")

	for _, value := range set.Array() {
		buffer.WriteString(strconv.FormatInt(value, 10))
		buffer.WriteString(" ")
	}

	buffer.WriteString("}")

	return buffer.String()
}
