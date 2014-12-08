// Copyright (c) 2014 Datacratic. All rights reserved.

package set

import (
	"testing"
)

func testUuintIntersect(t *testing.T, title string, aSet, bSet, exp Uint) {
	result := aSet.Intersect(bSet)

	t.Logf("%s: %s U %s = %s ?= %s", title, aSet, bSet, result, exp)

	for ID := range result {
		if !exp.Test(ID) {
			t.Errorf("FAIL(%s): extra ID %d", title, ID)
		}
	}

	for ID := range exp {
		if !result.Test(ID) {
			t.Errorf("FAIL(%s): missing ID %d", title, ID)
		}
	}
}

func uintIntersect(t *testing.T, title string, a, b, exp Uint) {
	testUuintIntersect(t, title, a, b, exp)
	testUuintIntersect(t, title, b, a, exp)
}

func TestUuintIntersect(t *testing.T) {
	uintIntersect(t, "empty",
		NewUint(uint64(1), uint64(2), uint64(3), uint64(4)),
		NewUint(),
		NewUint())

	uintIntersect(t, "overlap",
		NewUint(uint64(1), uint64(2), uint64(3), uint64(4)),
		NewUint(uint64(3), uint64(4), uint64(5), uint64(6)),
		NewUint(uint64(3), uint64(4)))

	uintIntersect(t, "subset",
		NewUint(uint64(1), uint64(2), uint64(3), uint64(4)),
		NewUint(uint64(1), uint64(2)),
		NewUint(uint64(1), uint64(2)))

	uintIntersect(t, "total",
		NewUint(uint64(1), uint64(2), uint64(3), uint64(4)),
		NewUint(uint64(1), uint64(2), uint64(3), uint64(4)),
		NewUint(uint64(1), uint64(2), uint64(3), uint64(4)))
}

func benchUuintIntersect(bench *testing.B, n, m int) {
	a, b := NewUint(), NewUint()

	for i := 0; i < n; i++ {
		a.Put(uint64(i))
	}

	for i := 0; i < n; i = i + 2 {
		b.Put(uint64(i))
		m--
	}
	for i := n; m > 0; m-- {
		b.Put(uint64(i))
		i++
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		a.Intersect(b)
	}
}

func BenchmarkUuintIntersectSmallSmall(b *testing.B) {
	benchUuintIntersect(b, 10, 10)
}

func BenchmarkUuintIntersectBigSmall(b *testing.B) {
	benchUuintIntersect(b, 10, 1000)
}

func BenchmarkUuintIntersectBigBig(b *testing.B) {
	benchUuintIntersect(b, 1000, 1000)
}
