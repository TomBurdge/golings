// anonymous functions2
// Make me compile!

package main

import "fmt"

func main() {
	var sayBye func(name string)
	name := "Tom"

	sayBye = func(n string) {
		fmt.Printf("Bye %s", n)
	}
	sayBye(name)
}
