# Data Types

## Common Types

| syntax   | Type                   |
| -------- | ---------------------- |
| `bool`   | true or false          |
| `string` | an array of characters |

## Numeric types

| syntax       | Type                                                                                |
| ------------ | ----------------------------------------------------------------------------------- |
| `uint`       | either 32 or 64 bits.                                                               |
| `int`        | same size as uint.                                                                  |
| `uintptr`    | an unsigned integer large enough to store the uninterpreted bits of a pointer value |
| `uint8`      | the set of all unsigned 8-bit integers (0 to 255)                                   |
| `uint16`     | the set of all unsigned 16-bit integers (0 to 65535)                                |
| `uint32`     | the set of all unsigned 32-bit integers (0 to 4294967295)                           |
| `uint64`     | the set of all unsigned 64-bit integers (0 to 18446744073709551615)                 |
| `int8`       | the set of all signed 8-bit integers (-128 to 127)                                  |
| `int16`      | the set of all signed 16-bit integers (-32768 to 32767)                             |
| `int32`      | the set of all signed 32-bit integers (-2147483648 to 2147483647)                   |
| `int64`      | the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807) |
| `float32`    | the set of all IEEE-754 32-bit floating-point numbers                               |
| `float64`    | the set of all IEEE-754 64-bit floating-point numbers                               |
| `complex64`  | the set of all complex numbers with `float32` real and imaginary parts              |
| `complex128` | the set of all complex numbers with `float64` real and imaginary parts              |
| `byte`       | alias for `uint8`                                                                   |
| `rune`       | alias for `int32` (represents a Unicode code point)                                 |

```go
package main

import (
  "fmt"
  "math/cmplx"
)

var (
  goIsFun bool       = true
  maxInt  uint64     = 1<<64 - 1
  complex complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
  const f = "%T(%v)\n" // print type(%T) and value(%v)

  fmt.Printf(f, goIsFun, goIsFun) // bool(true)
  fmt.Printf(f, maxInt, maxInt)   // uint64(18446744073709551615)
  fmt.Printf(f, complex, complex) // complex128((2+3i))
}
```

## Converting Values

```go
func main() {
  i := 42
  j := float64(i)
  k := uint(j)

  fmt.Printf("i %T : %v\n", i, i) // i int : 42
  fmt.Printf("j %T : %v\n", j, j) // j float64 : 42
  fmt.Printf("k %T : %v\n", k, k) // k uint : 42
}
```

## Type Assertion

A type assertion takes a value and tries to create another version in the specified explicit type.

```go
package main

import (
  "fmt"
  "time"
)

func timeMap(y interface{}) {
  z, ok := y.(map[string]interface{}) //asserting y as a map of interfaces

  if ok {
    z["updated_at"] = time.Now() //z now has the type map[string]interface 
  }
}

func main() {
  foo := map[string]interface{}{
    "Matt": 42,
  }

  timeMap(foo)
  fmt.Println(foo)
}
```

## Struct

A struct is a **collection of fields/properties**. You can define new types as structs or interfaces. You can list just a subset of fields by using the `"Name:" syntax`

```go
package main

import (
  "fmt"
  "time"
)

type Bootcamp struct {
  // Latitude of the event
  Lat float64
  // Longitude of the event
  Lon float64
  // Date of the event
  Date time.Time
}

func main() {
  fmt.Println(Bootcamp{
    Lat:  34.012836,
    Lon:  -118.495338,
    Date: time.Now(),
  })
}
```

Declaration of struct literals

```go
package main

import "fmt"

type Point struct {
  X, Y int
}

var (
  p = Point{1, 2}  // has type Point
  q = &Point{1, 2} // has type *Point
  r = Point{X: 1}  // Y:0 is implicit
  s = Point{}      // X:0 and Y:0
)

func main() {
  fmt.Println(p, q, r, s)
}
```

and we can access fields using the dot notation

```go
package main

import (
  "fmt"
  "time"
)

type Bootcamp struct {
  Lat, Lon float64
  Date     time.Time
}

func main() {
  event := Bootcamp{
    Lat: 34.012836,
    Lon: -118.495338,
  }

  event.Date = time.Now()
  fmt.Printf("Event on %s, location (%f, %f)",
    event.Date, event.Lat, event.Lon)
}
```
