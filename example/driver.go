package main

import (
	"fmt"

	"github.com/nobelium/SplayTree"
)

type Var struct {
	Key int
	Val string
}

func (x Var) Less(than splaytree.Item) bool {
	return x.Key < than.(Var).Key
}

func main() {
	fmt.Println("SplayTree Example")

	tree := splaytree.NewSplayTree()
	var1 := Var{Key: 1, Val: "a"}
	var2 := Var{Key: 2, Val: "b"}
	var3 := Var{Key: 3, Val: "c"}
	var4 := Var{Key: 4, Val: "d"}

	tree.Insert(var1, true)
	tree.Insert(var2, true)
	tree.Insert(var3, true)
	tree.Insert(var4, true)

	temp := Var{Key: 2, Val: "Doesn't matter"}
	res := tree.Find(temp)
	fmt.Println("Found splaynode: " + string(temp.Key) + " val: " + res.(Var).Val)

}
