# Collection Types

## Arrays

The type `[n]T` is an array of `n` values of type `T`.

```go
package main

import "fmt"

func main() {
    var a [2]string         //array of size 2
    a[0] = "Hello"          //Zero index of "a" has "Hello"
    a[1] = "World"          //1st index of "a" has "World"
    fmt.Println(a[0], a[1])   // will print Hello World
    fmt.Println(a)           // will print [Hello World]

    primes := [6]int{2, 3, 5, 7, 11, 13}  // array of prime numbers of size 6 
    fmt.Println(primes)
}
```

array of string

```go
package main

import "fmt"

func main() {
    a  := [2]string{"hello", "world!"}
    fmt.Printf("%q", a)
}
```

Trying to access of set a value at an index that doesn't exist will get an error

```go
package main

func main() {
    var a [2]string
    a[3] = "Hello"
}
```

```bash
# command-line-arguments
./main.go:5: invalid array index 3 (out of bounds for 2-element array)
```

## Slices

Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data. Slices hold references to an underlying array, and if you assign one slice to another, **both refer to the same array**.

```go
package main

import "fmt"

func main() {
    mySlice := []int{1, 2, 3, 4, 5, 6} // slice literal
    fmt.Println(mySlice) // [1 2 3 4 5 6]

    fmt.Println(mySlice[1:4]) // [2 3 4] from index 1 to exclude index 4

    fmt.Println(mySlice[:3]) // [1 2 3] from beginning to exclude index 3

    fmt.Println(mySlice[4:]) // [5 6] from index 4 to the end
}
```

### Making slices

We can also use `make` to create an empty slice of a specific length and then populate each entry.

```go
package main

import "fmt"

func main() {
    cities := make([]string, 3)

    cities[0] = "Santa Monica"
    cities[1] = "Venice"
    cities[2] = "Los Angeles"

    fmt.Printf("%q", cities) // ["Santa Monica" "Venice" "Los Angeles"]
}
```

### Appending to a slice

Let's see the example

```go
cities := []string{}
cities[0] = "Santa Monica" // runtime error
```

A slice is seating on top of an array, in this case, the array is empty and the slice can’t set a value in the referred array. We need to use `append` function.

```go
package main

import "fmt"

func main() {
    cities := []string{}
    cities = append(cities, "Bangkok", "Chiang Mai", "Nan") // append one of more
    fmt.Println(cities) // [Bangkok]
}
```

Or append with `ellipsis` (`...` syntax), the compiler will accept the operation **since the types are matching.**

```go
package main

import "fmt"

func main() {
    musics := []string{"Metal", "Disco"}
    otherMusic := []string{"Funk", "Rock"}
    musics = append(musics, otherMusic...)

    fmt.Printf("%q", musics) // ["Metal" "Disco" "Funk" "Rock"]

    // Length of slice
    fmt.Println(len(cities)) // 4
}
```

### Nil slices

The zero value of a slice is `nil`. A `nil` slice has a length and capacity of 0.

```go
package main

import "fmt"

func main() {
    var z []int
    fmt.Println(z, len(z), cap(z))
    // [] 0 0
    if z == nil {
        fmt.Println("nil!")
    }
    // nil!
}
```

Resources about Slices

- [Slice Tricks](https://github.com/golang/go/wiki/SliceTricks)
- [Go by Example: Slices](https://gobyexample.com/slices)
- [Practical Go Lessons: Slices](https://www.practical-go-lessons.com/chap-21-slices)
- [Effective Go: Slices](https://golang.org/doc/effective_go#slices)
