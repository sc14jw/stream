package stream

import "testing"

const (
	errorReturned   = "The error %v was displayed when none were expected"
	incorrectSlice  = "The slice %v does not equal the expected slice %v"
	incorrectValue  = "The value %v does not equal the expected value %v"
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

	if !sliceEqual(testItems, strm.ToSlice()) {
		t.Errorf(incorrectSlice, strm.ToSlice(), testItems)
	}

	strm.s = []interface{}{1, 2, 5, 6}

	if !sliceEqual([]interface{}{1, 2, 5, 6}, strm.s) {
		t.Errorf(incorrectSlice, strm.s, []interface{}{1, 2, 5, 6})
	}
}

func TestTransform(t *testing.T) {
	strm, _ := Of(testItems)

	strm.Transform(func(elem interface{}, i int) (newElem interface{}) {
		if elem.(string) == "hello" {
			return interface{}("hello changed")
		}

		return elem
	})

	if !sliceEqual([]interface{}{"hello changed", "this is a test"}, strm.s) {
		t.Errorf(incorrectSlice, []interface{}{"hello changed", "this is a test"}, strm.s)
	} else if sliceEqual(testItems, strm.s) {
		t.Errorf(matchingSlice, testItems, strm.s)
	}

	strm, _ = Of([]interface{}{1, 2, 3, 4})

	strm.Transform(func(elem interface{}, i int) (newElem interface{}) {
		if elem.(int) == 1 {
			return 2
		}

		return elem
	})

	if !sliceEqual([]interface{}{2, 2, 3, 4}, strm.s) {
		t.Errorf(incorrectSlice, []interface{}{2, 2, 3, 4}, strm.s)
	} else if sliceEqual(testItems, strm.s) {
		t.Errorf(matchingSlice, testItems, strm.s)
	}
}

func TestToMap(t *testing.T) {
	testMap := make(map[interface{}]interface{})
	testMap[interface{}("hello")] = interface{}("hello value")
	testMap[interface{}("this is a test")] = interface{}("this is a test value")

	s, err := Of(testItems)
	m := s.ToMap(func(elem interface{}, i int) (k interface{}, v interface{}) {
		if str, ok := elem.(string); ok {
			k = elem
			v = interface{}(str + " value")
			return
		}

		k = elem
		v = interface{}(4)
		return
	})

	if err != nil {
		t.Errorf(errorReturned, err)
	} else if !mapEqual(testMap, m) {
		t.Errorf(incorrectSlice, m, testMap)
	}
}

func TestFlatten(t *testing.T) {
	s, _ := Of(testItems)

	res := s.Flatten(func(acc interface{}, elem interface{}, i int) (res interface{}) {
		str := acc.(string)
		str += (", " + elem.(string))
		res = interface{}(str)
		return
	})

	if res != "hello, this is a test" {
		t.Errorf(incorrectValue, res, "hello, this is a test")
	}
}

func TestSort(t *testing.T) {
	s, _ := Of(testItems)

	sortedStream := s.Sort(func(first interface{}, second interface{}) (res bool) {
		firstStr := first.(string)
		if firstStr[0] != 'h' {
			res = true
		}

		return
	})

	if sortedStream.s[0] != testItems[1] {
		t.Errorf(incorrectValue, sortedStream.s[0], testItems[1])
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

func mapEqual(exp map[interface{}]interface{}, act map[interface{}]interface{}) (res bool) {
	if len(exp) != len(act) {
		return
	}

	for k, v := range exp {
		if v != act[k] {
			return
		}
	}

	res = true
	return
}
