gobst
===
The package `sethrachwitz/gobst` implements a binary search tree in go.  Current functionality allows for `Insert` and `Find` to be called on the tree.  Future functionality will include `Delete`, and print functions such as `PrePrint`, `OrderPrint`, and `PostPrint` (pre-order, in-order and post-order).

This package can be downloaded by running `go get github.com/sethrachwitz/gobst`.

## Methods

`Insert` takes the value being inserted into the tree as a parameter, and returns true if successfully inserted, or false if unsuccessful (value is already in the tree.)

`Find` takes the value being searched for in the tree as a parameter, and returns true or false depending on if the value was found in the tree, and an integer representing the number of levels traversed before finding / not finding the value.

The number of levels traversed starts at 1 (the root node), and increases by one every time a child is visited.  If `Find` found the value to be the root node, the number of levels traversed would be 1.  If the value was a direct child of the root node, the number of levels traversed would be 2, etc.

Example:

```go
package main

import (
	"github.com/sethrachwitz/gobst"
	"fmt"
)

func main() {
	tree := bst.NewBST()
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(5)

	node := 5

	found, levels := tree.Find(node)

	fmt.Print("Node ", node, " was ")

	if !found {
		fmt.Print("not ")
	}

	fmt.Print("found after searching ", levels)

	if levels == 1 {
		fmt.Println(" level.")
	} else {
		fmt.Println(" levels.")
	}
}
```

Output:
`Node 5 was found after searching 3 levels.`


## License

MIT licensed. See the LICENSE file for details.