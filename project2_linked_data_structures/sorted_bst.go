package main

import "fmt"

type Node struct {
	data  string
	left  *Node
	right *Node
}

func (node *Node) insertValue(value string) {
	cur_node := node
	new_node := Node{}
	for {
		// if the root is empty string
		if cur_node.data == "" {
			cur_node.data = value
			break
		}
		// look left
		if value < cur_node.data {
			if cur_node.left == nil {
				new_node.data = value
				cur_node.left = &new_node
				break
			}
			cur_node = cur_node.left
			continue
		}
		if value > cur_node.data {
			if cur_node.right == nil {
				new_node.data = value
				cur_node.right = &new_node
				break
			}
			cur_node = cur_node.right
			continue
		}
	}

}

func (node *Node) findValue(target string) *Node {
	cur_node := node
	for {
		// if we ever reach a nil then
		// we haven;t found the target
		if cur_node.data == target {
			return cur_node
		}
		if target < cur_node.data {
			if cur_node.left == nil {
				return nil
			}
			cur_node = cur_node.left
		}
		if target > cur_node.data {
			if cur_node.right == nil {
				return nil
			}
			cur_node = cur_node.right
		}
	}
}

func (node *Node) inorder() string {
	result := ""
	if (node.right != nil) && (node.left != nil) {
		result += node.left.inorder() + node.data + " " + node.right.inorder()
	} else if node.right != nil {
		result += node.data + " " + node.right.inorder()
	} else if node.left != nil {
		result += node.left.inorder() + node.data + " "
	} else {
		result += node.data + " "
	}
	return result
}

func main() {
	// Make a root node to act as sentinel.
	root := Node{"", nil, nil}

	// Add some values.
	root.insertValue("I")
	root.insertValue("G")
	root.insertValue("C")
	root.insertValue("E")
	root.insertValue("B")
	root.insertValue("K")
	root.insertValue("S")
	root.insertValue("Q")
	root.insertValue("M")

	// Add F.
	root.insertValue("F")

	// Display the values in sorted order.
	fmt.Printf("Sorted values: %s\n", root.inorder())

	// Let the user search for values.
	for {
		// Get the target value.
		target := ""
		fmt.Printf("String: ")
		fmt.Scanln(&target)
		if len(target) == 0 {
			break
		}

		// Find the value's node.
		node := root.findValue(target)
		if node == nil {
			fmt.Printf("%s not found\n", target)
		} else {
			fmt.Printf("Found value %s\n", target)
		}
	}
}
