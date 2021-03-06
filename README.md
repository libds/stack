# stack
A simple LIFO queue with 100% test coverage.


## Usage

```go
package main

import (
    "github.com/libds/stack"
    "fmt"
)

func main() {
    s := stack.New()
    s.Push("d1")
    s.Push("d2")

    v, err := s.Pop()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(v) // Output: d2

    if !s.Empty()  {
        v = s.Pop()
        fmt.Println(v) // Output: d1
    }
}
```


## API Docs

http://godoc.org/github.com/libds/stack



## LIcense


The MIT License (MIT)

Copyright (c) 2015 libds

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

