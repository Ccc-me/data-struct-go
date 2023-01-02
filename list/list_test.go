package list

import (
	"reflect"
	"testing"
)

type StructForTest struct {
	A int
	B string
}

func Test_List(t *testing.T) {

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
		e := l.PushBack(testList[i])

		got := testList[i]
		want := e.Value
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got = %v, want %v", got, want)
		}
		if !reflect.DeepEqual(l.len, i+1) {
			t.Errorf("got = %v, want %v", l.len, i+1)
		}
	}

	i := 0
	for f := l.Front(); f != nil; f = f.Next() {
		if !reflect.DeepEqual(f.Value, testList[i]) {
			t.Errorf("got = %v, want %v", f.Value, testList[i])
		}
		i++
	}

	for i := 0; i < len(testList); i++ {
		f := l.Front()
		if !reflect.DeepEqual(f.Value, testList[i]) {
			t.Errorf("got = %v, want %v", f.Value, testList[i])
		}

		v := l.Remove(f)
		if !reflect.DeepEqual(v, testList[i]) {
			t.Errorf("got = %v, want %v", v, testList[i])
		}

		if !reflect.DeepEqual(l.len, len(testList)-(i+1)) {
			t.Errorf("got = %v, want %v", l.len, len(testList)-(i+1))
		}
	}
}
