# go-ds
A well-tested and reliable golang lib of all the common data structures.

# Install
`go get github.com/souvikhaldar/go-ds/bst`

# Instructions
* Create a binary search tree, insert items, then print in sorted order (i.e In-Order traversal)  

```
package main

import "github.com/souvikhaldar/go-ds/bst"

func main() {
	t := bst.NewIterativeBst()
	t.Insert(5)
	t.Insert(9)
	t.Insert(1)
	t.Insert(1)
	t.Insert(3)
	t.InOrder()

}
```

