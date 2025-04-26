// anonymous functions1
// Make me compile!

package main

import "fmt"

func main() {
	pr := func(name string) {
		fmt.Printf("Hello %s", name)
	}
	pr("Tomo")
}
