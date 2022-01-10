# Go by example

## Values

Go has various value types including strings, integers, floats, booleans, etc.
Here are a few basic examples.

  ```go
  package main

  import "fmt"

  func main() {
    fmt.Println("go" + "lang")

    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
  }
  ```

- Strings, which can be added together with `+`.

  ```go
  fmt.Println("go" + "lang")
  ```

- Integers and floats.

  ```go
  fmt.Println("1+1 =", 1+1)
  ```

- Booleans, with boolean operators as you’d expect.

  ```go
  fmt.Println(true && false)
  fmt.Println(true || false)
  fmt.Println(!true)
  ```

- To run the program.

  ```sh
  $ go run values.go
  golang
  1+1 = 2
  7.0/3.0 = 2.3333333333333335
  false
  true
  false
  ```

- [Go by Example][0]
- [Go by Example: Values][1]

[0]: <https://gobyexample.com> "Go by Example"
[1]: <https://gobyexample.com/values> "Go by Example: Values"
