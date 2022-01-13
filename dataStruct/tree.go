package main

import "fmt"

type Tree struct {
	data  int
	left  *Tree
	right *Tree
}

func NewTree(data int) *Tree {
	tree := new(Tree)
	tree.data = data

	return tree
}

func (t *Tree) Insert(data int) {
	if t.data < data {
		if t.right == nil {
			t.right = &Tree{
				data: data,
			}
		} else {
			t.right.Insert(data)
		}
	} else {
		if t.left == nil {
			t.left = &Tree{
				data: data,
			}
		} else {
			t.left.Insert(data)
		}
	}
}

func (t *Tree) Print() {
	if t == nil {
		return
	}
	t.right.Print()
	fmt.Println(t.data)
	t.left.Print()

}

func main() {
	t := NewTree(8)
	for data := range []int{5,4,7,8,2,0} {
		t.Insert(data)
	}

	t.Print()
}
