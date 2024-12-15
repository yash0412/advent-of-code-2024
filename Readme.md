# Advent of Code 2024 in Golang

## Reference

[Advent of Code 2024](https://adventofcode.com/2024)

Each day's solution is kept inside it's own folder and package

To get the solution for a particular day, import the package into main method in main.go and call the Solve() or Solve2() method.

## Example

```go
package main

import (
	"adventofcode/dayone"
	"adventofcode/daytwo"
)

func main() {
	dayone.Solve()
	daytwo.Solve()
	daytwo.Solve2()
}

```
