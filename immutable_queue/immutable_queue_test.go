package immutable_queue

import "testing"

func TestIsEmpty(t *testing.T) {
	iq := New([]interface{}{})
	assert(t, true, iq.IsEmpty(), "iq.IsEmpty() is expected to be %t, but got=%t")

	iq = New([]interface{}{1, 2, 3, 4, 5})
	assert(t, false, iq.IsEmpty(), "iq.IsEmpty() is expected to be %t, but got=%t")
}

func TestHead(t *testing.T) {
	iq := New([]interface{}{})
	assert(t, nil, iq.Head(), "iq.Head() is expected to be %#v, but got=%#v")

	iq = New([]interface{}{1, 2, 3, 4, 5})
	assert(t, 1, iq.Head(), "iq.Head() is expected to be %#v, but got=%#v")
}

func TestTail(t *testing.T) {
	iq := New([]interface{}{})
	assert(t, nil, iq.Tail(), "iq.Tail() is expected to be %#v, but got=%#v")

	iq = New([]interface{}{1, 2, 3, 4, 5})
	assert(t, 5, iq.Tail(), "iq.Tail() is expected to be %#v, but got=%#v")
}

func TestEnQueue(t *testing.T) {
	iq := New([]interface{}{})
	assert(t, nil, iq.Head(), "iq.Head() is expected to be %#v, but got=%#v")

	iq2 := iq.EnQueue(1)
	assert(t, nil, iq.Head(), "iq.Head() is expected to be %#v, but got=%#v")
	assert(t, 1, iq2.Head(), "iq2.Head() is expected to be %#v, but got=%#v")
}

func TestDeQueue(t *testing.T) {
	iq := New([]interface{}{1, 2, 3})
	assert(t, 1, iq.Head(), "iq.Head() is expected to be %#v, but got=%#v")

	iq2 := iq.DeQueue()
	assert(t, 1, iq.Head(), "iq.Head() is expected to be %#v, but got=%#v")
	assert(t, 2, iq2.Head(), "iq2.Head() is expected to be %#v, but got=%#v")
}

func assert(t *testing.T, expected interface{}, actual interface{}, error string) {
	if actual != expected {
		t.Errorf(error, expected, actual)
	}
}
