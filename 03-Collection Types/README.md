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

## Range in For-Loops

The `range` form of the for loop iterates over a `slice` or a `map`. Being able to iterate over all the elements of a data structure is very useful and `range` simplifies the iteration.

syntax

```go
func main() {
    list := make([]int, 10) // create slice length of 10
    for i, v := range list { //loops over the length (0 - 9)
        fmt.Printf("%d = %d\n", i, v) // i = index, v = value
    }
}
```

more examples

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
    for i, v := range pow { //loops over the length of pow
        fmt.Printf("2**%d = %d\n", i, v)
    }
}
```

You can skip the index or value by assigning to `_`. If you only want the index, drop the “, value” entirely.

```go
package main

import "fmt"

func main() {
    pow := make([]int, 10) // [0 0 0 0 0 0 0 0 0 0]

    for i := range pow {
        pow[i] = 1 << uint(i) // pow = [1 2 4 8 16 32 64 128 256 512]
    }

    for _, value := range pow {
        fmt.Printf("%d\n", value) // print only value of pow
    }
}
```

## Break & Continue

You can stop the iteration by using `break` keyword.

```go
package main

import "fmt"

func main() {
    pow := make([]int, 10)

    for i := range pow {
        pow[i] = 1 << uint(i)

        if pow[i] >= 16 {
            break // stop when pow reaches 16
        }
    }

    fmt.Println(pow) // [1 2 4 8 16 0 0 0 0 0]
}
```

and `continue` will skip an iteration

```go
package main

import "fmt"

func main() {
    pow := make([]int, 10)

    for i := range pow {
        if i%2 == 0 {
            continue // skip
        }

        pow[i] = 1 << uint(i)
    }

    fmt.Println(pow) // [0 2 0 8 0 32 0 128 0 512]
}
```

Loop iterates over `key, value` pair

```go
package main

import "fmt"

func main() {
    cities := map[string]int{
        "Bangkok":    10539000,
        "Chiang Mai": 127240,
        "Nan":        20212,
    }

    for key, value := range cities {
        fmt.Printf("%s has %d populations\n", key, value)
    }

    // Nan has 20212 populations
    // Bangkok has 10539000 populations
    // Chiang Mai has 127240 populations
}
```

## Map

A map **maps** keys to values, like a **dictionary** or **hashes**

```go
package main

import "fmt"

func main() {
    rhcp := map[string]int{
        "Anthony Kiedis":  58,
        "Flea":            58,
        "AChad Smith":     59,
        "John Frusciante": 51,
    }

    fmt.Printf("%#v\n", rhcp)

    for key, value := range rhcp {
        fmt.Printf("%s is %d years old\n", key, value)
    }
}

// ## Result
// map[string]int{"AChad Smith":59, "Anthony Kiedis":58, "Flea":58, "John Frusciante":51}

// Anthony Kiedis is 58 years old
// Flea is 58 years old
// AChad Smith is 59 years old
// John Frusciante is 51 years old
```
