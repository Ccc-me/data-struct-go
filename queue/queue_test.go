package queue

import (
	"reflect"
	"testing"
)

type StructForTest struct {
	A int
	B string
}

func Test_Queue(t *testing.T) {

	testList := []*StructForTest{
		&StructForTest{A: 1, B: "SFDDASF"},
		&StructForTest{A: 2, B: "SFDDASF"},
		&StructForTest{A: 3, B: "SFDDASF"},
		&StructForTest{A: 4, B: "SFDDASF"},
		&StructForTest{A: 5, B: "SFDDASF"},
		&StructForTest{A: 6, B: "SFDDASF"},
	}

	l := New[*StructForTest]()
	for i := 0; i < len(testList); i++ {
		l.Push(testList[i])

		if !reflect.DeepEqual(l.Len(), i+1) {
			t.Errorf("got = %v, want %v", l.Len(), i+1)
		}
	}

	i := 0
	for l.Len() > 0 {
		item := l.Front()
		l.Pop()
		if !reflect.DeepEqual(item, testList[i]) {
			t.Errorf("got = %v, want %v", item, testList[i])
		}
		i++
	}

}
