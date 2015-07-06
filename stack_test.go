package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := New()
	_, e := s.Pop()
	if e != StackEmpty {
		t.Fatal("Must return error if stack empty")
	}
	s.Push("Early but late.")
	s.Push("Late but early")

	if v := s.Size(); v != 2 {
		t.Fatal("Size must be two")
	}

	if v := s.Arr(); len(v) != s.Size() {
		t.Fatal("Arr and Size must be equals")
	}

	if v, _ := s.Pop(); v != "Late but early" {
		t.Fatal("Value must be 'Late but early'")
	}

	if v := s.Size(); v != 1 {
		t.Fatal("Size must be one")
	}

	if v, _ := s.Pop(); v != "Early but late." {
		t.Fatal("Value must be 'Early but late.'")
	}

	if v := s.Size(); v != 0 {
		t.Fatal("Size must be zero")
	}

	if v := s.Size(); (v == 0) != s.Empty() {
		t.Fatal("Empty and Size must be equals")
	}

}
