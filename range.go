package main

import "fmt"

func main() {
  ids := []int{32,43,89,600}

  // Loop through ids
  for i, id := range ids {
	fmt.Printf("%d - ID: %d\n", i, id)
  }
}
