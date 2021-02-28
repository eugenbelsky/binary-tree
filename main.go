package main

import "fmt"

type Node struct {
	right  *Node
	left   *Node
	parent *Node
	value  int
}

type Tree struct {
	root *Node
}

func main() {
	var tree Tree
	// tree.root = &Node{value: 2, left: nil, right: nil}

	input := []int{2, 4, 1, 0, 10, 123, 10, 20, -4, -2, 100000}
	for _, value := range input {
		tree.Insert(value)
	}

	fmt.Println(tree.FindMaxNode())
	del, _ := tree.Find(123)
	fmt.Println(del)
	tree.Delete(del)
	newv := tree.FindMaxNode()
	fmt.Println(newv)
	// fmt.Println(del.parent.value)
	// fmt.Println(del)
	// fmt.Println(del.right.value)

}
func (t *Tree) Insert(data int) {
	if t.root == nil {
		t.root = &Node{value: data}
	} else {
		t.root.Insert(data)
	}
}
func (n *Node) Insert(data int) {
	if data <= n.value {
		if n.left != nil {
			n.left.Insert(data)
		} else {
			n.left = &Node{value: data, parent: n}

		}
	}
	if data > n.value {
		if n.right != nil {
			n.right.Insert(data)
		} else {
			n.right = &Node{value: data, parent: n}
		}
	}
}

func (t *Tree) Find(data int) (*Node, bool) {
	return t.root.Find(data)
}
func (n *Node) Find(data int) (*Node, bool) {

	if n == nil {
		return nil, false
	}

	switch {
	case data == n.value:
		return n, true

	case data < n.value:
		return (n.left.Find(data))

	default:
		return (n.right.Find(data))

	}
}

func (t *Tree) FindMaxNode() *Node {
	return t.root.FindMaxNode()
}

func (n *Node) FindMaxNode() *Node {
	if n.right == nil {
		return n
	} else {
		return n.right.FindMaxNode()
	}
}

func (t *Tree) FindMinNode() *Node {
	return t.root.FindMinNode()
}
func (n *Node) FindMinNode() *Node {
	if n.left == nil {
		return n

	} else {
		return n.left.FindMinNode()
	}
}
func (t *Tree) Delete(node *Node) {

	node.Delete() //fails when trying to delete the root

}
func (n *Node) Delete() {
	switch {
	case (n.right == nil && n.left == nil): // if leaf -> just remove

		n.ReplaceNode(nil)

	case (n.right == nil || n.left == nil): // if has only one branch
		if n.right != nil {
			n.ReplaceNode(n.right)

		} else {
			n.ReplaceNode(n.left)

		}

	case (n.right != nil && n.left != nil): // if has both branches

		repl := n.left.FindMaxNode()
		n.ReplaceNode(repl)

	}

}

func (n *Node) ReplaceNode(newNode *Node) {
	if n.left != nil { // if has left leaf
		n.left.parent = newNode

	}
	if n.right != nil {
		n.right.parent = newNode

	}

	if n.parent != nil {
		if n.parent.left == n {
			n.parent.left = newNode

		} else {

			n.parent.right = newNode
		}
	}

	// n = nil // remove pointer to parent on old node to trigger garbage collection
}
