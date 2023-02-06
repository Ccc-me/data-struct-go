package priority_queue

import (
	"reflect"
	"testing"
)

type StructForTest struct {
	A int
	B string
}

func less(a, b *StructForTest) bool {
	return a.A < b.A
}

func check(t *PriorityQueue[*StructForTest]) bool {

	a := -1
	for i := 0; i < t.len(); i++ {
		if i != 0 && a > t.Pop().A {
			return false
		}
		a = t.Pop().A
	}
	return true
}

func Test_PriorityQueue(t *testing.T) {

	testList := []*StructForTest{
		&StructForTest{A: 0, B: "SFDDASF"},
		&StructForTest{A: 0, B: "SFDDASF"},
		&StructForTest{A: 0, B: "SFDDASF"},
		&StructForTest{A: 0, B: "SFDDASF"},
		&StructForTest{A: 0, B: "SFDDASF"},
		&StructForTest{A: 0, B: "SFDDASF"},
	}

	newTestList := New[*StructForTest](testList, func(a, b *StructForTest) bool {
		return a.A < b.A
	})

	if !reflect.DeepEqual(check(newTestList), true) {
		t.Errorf("got = %v, want %v", check(newTestList), true)
	}

	testList = []*StructForTest{
		&StructForTest{A: 6, B: "SFDDASF"},
		&StructForTest{A: 5, B: "SFDDASF"},
		&StructForTest{A: 4, B: "SFDDASF"},
		&StructForTest{A: 3, B: "SFDDASF"},
		&StructForTest{A: 2, B: "SFDDASF"},
		&StructForTest{A: 1, B: "SFDDASF"},
	}

	newTestList = New[*StructForTest](testList, func(a, b *StructForTest) bool {
		return a.A < b.A
	})

	if !reflect.DeepEqual(check(newTestList), true) {
		t.Errorf("got = %v, want %v", check(newTestList), true)
	}
}
