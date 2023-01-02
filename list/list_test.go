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
	l := New[*StructForTest]()

	got := &StructForTest{A: 1, B: "SFDDASF"}
	want := l.PushBack(got).Value
	if !reflect.DeepEqual(got, want) {
		t.Errorf("New() = %v, want %v", got, want)
	}
	if !reflect.DeepEqual(l.len, 1) {
		t.Errorf("New() = %v, want %v", l.len, 1)
	}
	got1 := &StructForTest{A: 1, B: "SFDDASF"}
	want1 := l.PushBack(got1).Value
	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("New() = %v, want %v", got1, want1)
	}

	got2 := got
	want2 := l.Front().Value
	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("New() = %v, want %v", got2, want2)
	}

	got3 := got
	want3 := l.Remove(l.Front())
	if !reflect.DeepEqual(got3, want3) {
		t.Errorf("New() = %v, want %v", got3, want3)
	}
	if !reflect.DeepEqual(l.len, 1) {
		t.Errorf("New() = %v, want %v", l.len, 1)
	}
}
