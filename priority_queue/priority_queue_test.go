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

func check(t []*StructForTest) bool {

	for i := 0; i < len(t); i++ {
		i21 := 2*i + 1
		i22 := 2*i + 2

		if i22 >= len(t) {
			break
		} else if i22 < len(t) {
			if !less(t[i], t[i21]) || !less(t[i], t[i22]) {
				return false
			}
		} else if i21 < len(t) {
			return less(t[i], t[i21])
		}
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

	_ = New[*StructForTest](testList, func(a, b *StructForTest) bool {
		return a.A < b.A
	})

	if !reflect.DeepEqual(check(testList), true) {
		t.Errorf("got = %v, want %v", check(testList), true)
	}
}
