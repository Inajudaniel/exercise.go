// package main

// import (
// 	"fmt"
// )

// func main() {
// 	for i := 1; i <= 20; i += 2 {
// 		fmt.Println(i)
// 	}
// }

package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 20; i++ {

		if i%2 == 0 {
			fmt.Println(i, "even number")
		} else {

			fmt.Println(i, "odd number")
		}

	}

}
