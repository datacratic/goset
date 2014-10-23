// Copyright (c) 2014 Datacratic. All rights reserved.

package set

import (
	"testing"
)

func testIntIntersect(t *testing.T, title string, aSet, bSet, exp Int) {
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

func intIntersect(t *testing.T, title string, a, b, exp Int) {
	testIntIntersect(t, title, a, b, exp)
	testIntIntersect(t, title, b, a, exp)
}

func TestIntIntersect(t *testing.T) {
	intIntersect(t, "empty",
		NewInt(int64(1), int64(2), int64(3), int64(4)),
		NewInt(),
		NewInt())

	intIntersect(t, "overlap",
		NewInt(int64(1), int64(2), int64(3), int64(4)),
		NewInt(int64(3), int64(4), int64(5), int64(6)),
		NewInt(int64(3), int64(4)))

	intIntersect(t, "subset",
		NewInt(int64(1), int64(2), int64(3), int64(4)),
		NewInt(int64(1), int64(2)),
		NewInt(int64(1), int64(2)))

	intIntersect(t, "total",
		NewInt(int64(1), int64(2), int64(3), int64(4)),
		NewInt(int64(1), int64(2), int64(3), int64(4)),
		NewInt(int64(1), int64(2), int64(3), int64(4)))
}

func benchIntIntersect(bench *testing.B, n, m int) {
	a, b := NewInt(), NewInt()

	for i := 0; i < n; i++ {
		a.Put(int64(i))
	}

	for i := 0; i < n; i = i + 2 {
		b.Put(int64(i))
		m--
	}
	for i := n; m > 0; m-- {
		b.Put(int64(i))
		i++
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		a.Intersect(b)
	}
}

func BenchmarkIntIntersectSmallSmall(b *testing.B) {
	benchIntIntersect(b, 10, 10)
}

func BenchmarkIntIntersectBigSmall(b *testing.B) {
	benchIntIntersect(b, 10, 1000)
}

func BenchmarkIntIntersectBigBig(b *testing.B) {
	benchIntIntersect(b, 1000, 1000)
}
