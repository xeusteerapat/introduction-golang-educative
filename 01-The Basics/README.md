# Basics of Go

## Variables

Using `var` keyword declares a list variables

```go
var (
  name  string
  age   int
  location  string
)
```

variables can also declared one-by-one

```go
var name     string
var age      int
var location string
```

or initialize multiple variables at once

```go
var (
  name, location, age = "Prince Oberyn", "Dorne", 32
)
```

Using `:=` syntax for **implicit type** declaration, then you don't need to specify type of variable

```go
package main
import "fmt"

func main() {
  name, location := "Teerapat Prommarak", "Bangkok"
  age := 35
  fmt.Printf("%s age %d from %s ", name, age, location)
  // Go know that name and location are string and age is int
}
```

A variable can contain any type, including functions

```go
func main() {
  myFunc := func() {
    fmt.Printf("From my func")
  }

  myFunc()
}
```

Constants using `const` keyword and cannot using `:=` syntax

```go
const Pi = 3.14
const (
  StatusOK                   = 200
  StatusCreated              = 201
  StatusAccepted             = 202
  StatusNonAuthoritativeInfo = 203
  StatusNoContent            = 204
  StatusResetContent         = 205
  StatusPartialContent       = 206
)
```

## Printing variables

Using `fmt` package as we can see the example as above

```go
package main
import "fmt"

func main() {
  name := "Caprica-Six"
  aka := fmt.Sprintf("Number %d", 6)
  fmt.Printf("%s is also known as %s",
    name, aka) //printing variables within the string
}
```

read more about [`fmt` package documentation](https://golang.org/pkg/fmt/#hdr-Printing)

## Packages and Import

### Packages

Every Go program is made up of packages. Programs start running in package `main`.

```go
package main

import "fmt"

func main() {
  fmt.Printf("Hello, World!\n")
}
```

We need to define a main `package` and a `main()` function which will be the entry point to our software.

### Import statement

individual

```go
import "fmt"
import "math/rand"
```

or grouped

```go
import (
  "fmt"
  "math/rand"
)
```

Usually, nonstandard lib packages are namespaced using a web url.

```go
import "github.com/mattetti/goRailsYourself/crypto"
```

## Functions

When declaring functions, the type comes after the variable name in the inputs. The return type(s) are then specified after the function name and inputs.

```go
package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println((add(3, 5)))
}
```

We also can return multiple values

```go
func location(city string) (string, string) {
  var region string
  var continent string  

  switch city {
    case "Los Angeles", "LA", "Santa Monica":
      region, continent = "California", "North America"
    case "New York", "NYC":
      region, continent = "New York", "North America"
    default:
      region, continent = "Unknown", "Unknown"
  }

  return region, continent
}

func main() {
  region, continent := location("Santa Monica")
  fmt.Printf("Matt lives in %s, %s", region, continent)
  // region = California, continent = North America
}
```

Or we can defined the "result parameters" like this

```go
func location(city string) (region, continent string) {
  ...
  return // returning region and continent
}
```

But we recommended the first one because we can obviously see the return variables.

## Pointers

Note that by default Go **passes arguments by value** (copying the arguments). We need to pass **pointers** if we want to **pass the arguments by reference.**

Use `&` symbol in front of the value **to get the pointer** and use the `*` for **dereference a pointer**

```go
package main

import "fmt"

func main() {
  var yourAge *int // dereference

  var myAge int = 35
  fmt.Printf("&i=%p i=%v\n", &myAge, myAge) // 0xc0000ae008

  yourAge = &myAge // get the pointer of myAge variable

  fmt.Println(yourAge) // 0xc0000ae008 this will print the same memory address with my age
}
```
