// Package stream provides a fluent interface for interacting with a given slice including filtering, transformations and ordering.
package stream

import (
	"github.com/bradfitz/slice"
	"github.com/sc14jw/optional"
)

// NilSliceError contains the error message returned should a slice be attempted to be created from a Nil slice.
const NilSliceError = "The given slice was nil resulting in a stream being unable to be created"

// Stream provides a fluent interface for interacting with a given slice. Note: a stream will create a new slice of the slice passed to it rather than modifying the original slice.
type Stream struct {
	s []interface{}
}

// Filter allows for elements within the stream to be filtered based on a given function. The function is applied to each element within the stream being passed the current element and index.
// The function should return true if the passed in element should be included as part of the filter otherwise false. Returns the Stream object reference.
func (s *Stream) Filter(f func(*interface{}, int) bool) (strm *Stream) {
	strm = s
	newS := make([]interface{}, 0)
	for i, elem := range strm.s {
		if f(&elem, i) {
			newS = append(newS, elem)
		}
	}
	strm.s = newS
	return
}

// Transform run a passed in function over all elements within the given Stream's slice of elements. The function is applied to each element within the stream being passed the current element and the index.
// This function should return the altered element as it's return value. Returns the stream reference.
func (s *Stream) Transform(f func(interface{}, int) interface{}) (strm *Stream) {
	strm = s
	newS := make([]interface{}, 0)
	for i, elem := range strm.s {
		newS = append(newS, f(elem, i))
	}
	strm.s = newS
	return
}

// Sort sorts a Stream using the https://godoc.org/github.com/bradfitz/slice sorting package with a given sorting function. The given function must take two interface representing the elements to be compared against each other returning true if the first element is less than
// the second, else false. This function will return the stream you have been working using sorted in accordance with the passed in function.
func (s *Stream) Sort(sortFunc func(i interface{}, j interface{}) bool) (strm *Stream) {
	strm = s

	convFunc := func(i, j int) (res bool) {
		res = sortFunc(s.s[i], s.s[j])
		return
	}

	slice.Sort(strm.s, convFunc)
	return
}

// Flatten use a passed in function to flatten a given stream into a singular object through running a function against all elements within the stream.
// This passed in function should take in an accumilator showing the current state of the flatten stream, the next element in the stream and the current index of the stream.
func (s *Stream) Flatten(f func(interface{}, interface{}, int) interface{}) (accum interface{}) {
	accum = s.s[0]
	for i, elem := range s.s[1:] {
		accum = f(accum, elem, i)
	}

	return
}

// ToMap allows for a given stream's contents to be converted to a map using a given function. This function should take an element and an index and
// return two parameters, the key for the element first, and the value representing this element second.
func (s *Stream) ToMap(f func(interface{}, int) (interface{}, interface{})) (m map[interface{}]interface{}) {
	m = make(map[interface{}]interface{})
	for i, elem := range s.s {
		key, value := f(elem, i)
		m[key] = value
	}
	return
}

// ToSlice returns the given stream as a slice of interfaces.
func (s *Stream) ToSlice() (elems []interface{}) {
	elems = s.s
	return
}

// Of returns a new Stream from a given slice of interfaces. An error with the message NilSliceError will be returned should the given slice be nil.
func Of(list []interface{}) (strm *Stream, err error) {
	opt, err := optional.NotNilWithMessage(list, NilSliceError)
	if !opt.WasInitialized() {
		return
	}

	s := make([]interface{}, len(list))
	copy(s, list)
	strm = &Stream{s: s}
	return
}
