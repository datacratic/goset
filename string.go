// Copyright (c) 2014 Datacratic. All rights reserved.

package set

import (
	"bytes"
	"sort"
)

// String represents a set of strings. Must first be initialized using the
// NewString function before being used.
type String map[string]struct{}

// NewString creates a new string set from the given list of strings.
func NewString(values ...string) String {
	set := make(map[string]struct{})

	for _, value := range values {
		set[value] = struct{}{}
	}

	return set
}

// Test returns true if the given string is in the set.
func (set String) Test(values ...string) bool {
	for _, value := range values {
		if _, ok := set[value]; !ok {
			return false
		}
	}
	return true
}

// Put adds the given string to the set and is a noop if the string is already
// in the set.
func (set String) Put(values ...string) {
	for _, value := range values {
		set[value] = struct{}{}
	}
}

// Del removes the given string from the set and is a noop if the string is not
// already in the set.
func (set String) Del(values ...string) {
	for _, value := range values {
		delete(set, value)
	}
}

// Copy returns a copy of the set.
func (set String) Copy() String {
	result := NewString()
	for value := range set {
		result[value] = struct{}{}
	}
	return result
}

// Merge adds all the values of other into the set.
func (set String) Merge(other String) {
	for value := range other {
		set[value] = struct{}{}
	}
}

// Union returns a set consisting of all values present in both other and the
// current set.
func (set String) Union(other String) String {
	result := set.Copy()
	result.Merge(other)
	return result
}

// Intersect returns a set consisting of the values present in both sets.
func (set String) Intersect(other String) String {
	var big, small String
	if len(set) < len(other) {
		big = other
		small = set
	} else {
		big = set
		small = other
	}

	result := NewString()
	for value := range small {
		if _, ok := big[value]; ok {
			result[value] = struct{}{}
		}
	}

	return result
}

// Difference returns a set consisting of all the values in this set that are
// not present in the other set.
func (set String) Difference(other String) String {
	result := set.Copy()

	for value := range other {
		delete(result, value)
	}

	return result
}

// Array returns a sorted array of all the values in the current set.
func (set String) Array() (list []string) {
	for value := range set {
		list = append(list, value)
	}
	sort.Strings(list)
	return
}

// String returns a string representation of the set suitable for debugging.
func (set String) String() string {
	if len(set) == 0 {
		return "{}"
	}

	buffer := new(bytes.Buffer)
	buffer.WriteString("{ ")

	for _, value := range set.Array() {
		buffer.WriteString(value)
		buffer.WriteString(" ")
	}

	buffer.WriteString("}")

	return buffer.String()
}
