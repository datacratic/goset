// Copyright (c) 2014 Datacratic. All rights reserved.

package set

import (
	"strconv"
	"testing"
)

func testStringIntersect(t *testing.T, title string, aSet, bSet, exp String) {
	result := aSet.Intersect(bSet)

	t.Logf("%s: %s U %s = %s ?= %s", title, aSet, bSet, result, exp)

	for ID := range result {
		if !exp.Test(ID) {
			t.Errorf("FAIL(%s): extra ID %s", title, ID)
		}
	}

	for ID := range exp {
		if !result.Test(ID) {
			t.Errorf("FAIL(%s): missing ID %s", title, ID)
		}
	}
}

func stringIntersect(t *testing.T, title string, a, b, exp String) {
	testStringIntersect(t, title, a, b, exp)
	testStringIntersect(t, title, b, a, exp)
}

func TestStringIntersect(t *testing.T) {
	stringIntersect(t, "empty",
		NewString("1", "2", "3", "4"),
		NewString(),
		NewString())

	stringIntersect(t, "overlap",
		NewString("1", "2", "3", "4"),
		NewString("3", "4", "5", "6"),
		NewString("3", "4"))

	stringIntersect(t, "subset",
		NewString("1", "2", "3", "4"),
		NewString("1", "2"),
		NewString("1", "2"))

	stringIntersect(t, "total",
		NewString("1", "2", "3", "4"),
		NewString("1", "2", "3", "4"),
		NewString("1", "2", "3", "4"))
}

func benchStringIntersect(bench *testing.B, n, m int) {
	a, b := NewString(), NewString()

	for i := 0; i < n; i++ {
		a.Put(strconv.FormatInt(int64(i), 10))
	}

	for i := 0; i < n; i = i + 2 {
		b.Put(strconv.FormatInt(int64(i), 10))
		m--
	}
	for i := n; m > 0; m-- {
		b.Put(strconv.FormatInt(int64(i), 10))
		i++
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		a.Intersect(b)
	}
}

func BenchmarkStringIntersectSmallSmall(b *testing.B) {
	benchStringIntersect(b, 10, 10)
}

func BenchmarkStringIntersectBigSmall(b *testing.B) {
	benchStringIntersect(b, 10, 1000)
}

func BenchmarkStringIntersectBigBig(b *testing.B) {
	benchStringIntersect(b, 1000, 1000)
}
