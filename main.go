package main

import "fmt"

type Node struct {
	right *Node
	left  *Node
	value int
}

type Tree struct {
	root *Node
}

func main() {
	var tree Tree
	// tree.root = &Node{value: 2, left: nil, right: nil}

	input := []int{200000000, 4000000, -1000, -4000, 2, 4, 1, 0, 10, 123, 10, 20, -4, -2, 100000}
	for _, value := range input {
		tree.Insert(value)
	}

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
			n.left = &Node{value: data}

		}
	}
	if data > n.value {
		if n.right != nil {
			n.right.Insert(data)
		} else {
			n.right = &Node{value: data}
		}
	}
}

func (t *Tree) FindNode(data int) (*Node, bool) {
	if t.root == nil {
		return nil, false
	}
	return t.root.FindNode(data)
}
func (n *Node) FindNode(data int) (*Node, bool) {

	if n == nil {
		return nil, false
	}

	switch {
	case data == n.value:
		return n, true

	case data < n.value:
		return (n.left.FindNode(data))

	default:
		return (n.right.FindNode(data))

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

func (t *Tree) FindParentNode(child *Node) *Node {
	if child == nil {
		return nil
	}
	if t == nil {
		return nil
	}

	return t.root.FindParentNode(child)
}
func (n *Node) FindParentNode(child *Node) *Node {

	if n == nil {
		return nil
	}

	if n.left == child {
		return n
	} else if n.right == child {
		return n
	} else {
		l := n.left.FindParentNode(child)
		if l != nil {
			return l
		}
		r := n.right.FindParentNode(child)

		if r != nil {
			return r
		}

		return nil

	}

}

func PrintPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.value)
		PrintPreOrder(n.left)
		PrintPreOrder(n.right)
	}
}

// func (t *Tree) Delete(node *Node) bool {		//buggy
// 	if node == t.root { // if delete root
// 		parent := &Node{right: t.root}
// 		st := node.Delete(parent)
// 		fmt.Println(1)
// 		switch {
// 		case t.root.left == nil && t.root.right == nil:
// 			return false
// 		default:
// 			status := node.Delete(parent)
// 			if status == true {
// 				fmt.Println(node)
// 				t.root = node
// 				return st
// 			}
// 		}
// 	}

// 	parent := t.FindParent(node) // if basic case
// 	return node.Delete(parent)

// }
// func (n *Node) Delete(parent *Node) bool {

// 	if n == nil {
// 		return false
// 	}
// 	switch {
// 	case (n.right == nil && n.left == nil): // if leaf -> just remove
// 		fmt.Println(n)
// 		n.ReplaceNode(parent, nil)

// 		return true

// 	case (n.right == nil || n.left == nil): // if has only one branch -> replace with the branch leaft
// 		if n.right != nil {

// 			n.ReplaceNode(parent, n.right)
// 			return true

// 		}

// 		n.ReplaceNode(parent, n.left)
// 		return true

// 	case (n.right != nil && n.left != nil): // if has both branches -> find max on left side and replace
// 		replace := n.left.FindMax()
// 		replaceParent := n.left.FindParent(replace)
// 		n.value = replace.value

// 		replace.Delete(replaceParent)
// 		return true
// 	}

// 	return false

// }

// func (n *Node) ReplaceNode(parent, replace *Node) {

// 	if parent != nil { //modify parent`s pointer
// 		if parent.left == n {
// 			parent.left = replace
// 			return

// 		}
// 		parent.right = replace
// 		return
// 	}

// 	// n = nil // remove pointer to parent on old node to trigger garbage collection
// }
