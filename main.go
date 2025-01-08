package main

import "fmt"

func main() {
	// Counter() We dont need another Counter!
	
	fmt.Println("HELLO WORLD")

	Counter()
  }

func Counter() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d -> Hello from git\n", i) // 2nd time
	}
}

