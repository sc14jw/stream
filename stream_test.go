package stream

import "testing"

const (
	errorReturned   = "The error %v was displayed when none were expected"
	incorrectSlice  = "The slice %v does not equal the expected slice %v"
	matchingSlice   = "The slice %v equals the slice %v when they should not be equal"
	noErrorReturned = "No Error was returned"
	notNil          = "Stream was not nil"
)

var testItems = []interface{}{"hello", "this is a test"}

func TestOf(t *testing.T) {
	strm, err := Of(testItems)

	if err != nil {
		t.Errorf(errorReturned, err)
	} else if !sliceEqual(strm.s, testItems) {
		t.Errorf(incorrectSlice, strm.s, testItems)
	}

	strm.s = strm.s[:0]

	if sliceEqual(strm.s, testItems) {
		t.Errorf(matchingSlice, strm.s, testItems)
	}

	strm, err = Of(nil)

	if err == nil {
		t.Error(noErrorReturned)
	} else if strm != nil {
		t.Error(notNil)
	}
}

func TestFilter(t *testing.T) {
	strm, _ := Of(testItems)

	strm.Filter(func(elem *interface{}, i int) (inc bool) {
		s := (*elem).(string)
		inc = s == "hello"
		return
	})

	if !sliceEqual(strm.s, []interface{}{"hello"}) {
		t.Errorf(incorrectSlice, []interface{}{"hello"}, strm.s)
	}
}

func TestToSlice(t *testing.T) {
	strm, _ := Of(testItems)

	if !sliceEqual(testItems, strm.s) {
		t.Errorf(incorrectSlice, strm.s, testItems)
	}

	strm.s = []interface{}{1, 2, 5, 6}

	if !sliceEqual([]interface{}{1, 2, 5, 6}, strm.s) {
		t.Errorf(incorrectSlice, strm.s, []interface{}{1, 2, 5, 6})
	}
}

func sliceEqual(exp []interface{}, act []interface{}) (res bool) {
	if len(exp) != len(act) {
		return
	}

	for i := range exp {
		if exp[i] != act[i] {
			return
		}
	}

	res = true
	return
}
