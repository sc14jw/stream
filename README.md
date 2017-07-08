# stream 
[![Build Status](https://travis-ci.org/sc14jw/stream.svg?branch=master)](https://travis-ci.org/sc14jw/stream) [![Coverage Status](https://coveralls.io/repos/github/sc14jw/stream/badge.svg?branch=master)](https://coveralls.io/github/sc14jw/stream?branch=master)[![Go Report Card](https://goreportcard.com/badge/github.com/sc14jw/stream)](https://goreportcard.com/report/github.com/sc14jw/stream)

A small Go libary for supplying a fluent interface for interacting with Go slices. A stream can be created with the following
line:
```go
testData := []interface{}{"test", "test2"} // slices passed into the stream must be of type []interface{}
strm := stream.Of(testData) // we now have a stream. Note: operations using this stream will not effect the original list
```
Using this stream it is possible to complete filter, flatten or transform operations shown below:
```go
testData := []interface{}{"test", "test2"}
newSlice := stream.Of(testData).
                   Filter(
                                func(elem interface{}, i int)(res bool){
                                  res = elem.(string) == "test"
                                }
                         ).
                   Transform(
                                func(elem interface{}, i int)(new interface{}){
                                  new = interface{}(elem.(string)[0])
                                }
                             ).
                   ToSlice()

// newSlice = []interface{}{"t"}
```
Through a stream it is also possible to map a slice of interface{} types to a map with a key of interface{} and value interface{}
as shown:
``` go
map := stream.Of(testData).
              ToMap(
                      func(elem interface{}, i int)(k interface{}, v interface{}){
                        k = elem
                        v = interface{}(elem.(string) += " test")
                      }
                    )
                    
fmt.Println(map[interface{}("test")]) // test test
fmt.Println(map[interface{}("test2"]) // test2 test
```
