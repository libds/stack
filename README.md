# stack
LIFO queue


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
