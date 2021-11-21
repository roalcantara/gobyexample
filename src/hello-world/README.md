# Go by example

## Hello World

- To print the classic `“hello world”` message

  ```go
  package main

  import "fmt"

  func main() {
    fmt.Println("hello world")
  }
  ```

- To run the program, put the code in hello-world.go and use go run.

  ```sh
  $ go run hello-world.go
  hello world
  ```

- Sometimes we’ll want to build our programs into binaries. We can do this using go build.

  ```sh
  $ go build hello-world.go
  $ ls
  hello-world    hello-world.go
  ```

- We can then execute the built binary directly.

  ```sh
  $ ./hello-world
  hello world
  ```

- [Go by Example][0]
- [Go by Example: Hello World][1]

[0]: <https://gobyexample.com> "Go by Example"
[1]: <https://gobyexample.com/hello-world> "Go by Example: Hello World"
