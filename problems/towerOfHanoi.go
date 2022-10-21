package main

import (
	"fmt"

	"github.com/souvikhaldar/go-ds/stacks"
)

// Graphical represtaation of 2 disks TOH
func TwoDisksTOH() {
	src := stacks.NewStacks()
	src.Push(2)
	src.Push(1)
	aux := stacks.NewStacks()
	dst := stacks.NewStacks()

	fmt.Println("Intial state...")
	fmt.Println("-------------")
	fmt.Println("Source stack")
	src.Print()
	fmt.Println("Destination stack")
	dst.Print()
	fmt.Println("Aux stack")
	aux.Print()
	fmt.Println("-------------")

	fmt.Println("first step...")
	fmt.Println("-------------")
	el, _ := src.Pop()
	aux.Push(el)
	fmt.Println("Source stack")
	src.Print()
	fmt.Println("Destination stack")
	dst.Print()
	fmt.Println("Aux stack")
	aux.Print()
	fmt.Println("-------------")

	fmt.Println("second step...")
	fmt.Println("-------------")
	el, _ = src.Pop()
	dst.Push(el)
	fmt.Println("Source stack")
	src.Print()
	fmt.Println("Destination stack")
	dst.Print()
	fmt.Println("Aux stack")
	aux.Print()
	fmt.Println("-------------")

	fmt.Println("third step...")
	fmt.Println("-------------")
	el, _ = aux.Pop()
	dst.Push(el)
	fmt.Println("Source stack")
	src.Print()
	fmt.Println("Destination stack")
	dst.Print()
	fmt.Println("Aux stack")
	aux.Print()
	fmt.Println("-------------")

}
func main() {
	fmt.Println("Visual representation of Tower of Hanoi for 2 disks..")
	TwoDisksTOH()
}
