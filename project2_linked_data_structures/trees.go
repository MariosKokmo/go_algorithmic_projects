package main

import (
	"errors"
	"fmt"
)

// ======================================================== Queue ===============================================================
type Cell struct {
	data *Node
	next *Cell
	prev *Cell
}

type DoublyLinkedList struct {
	topSentinel    *Cell
	bottomSentinel *Cell
}

// Add a cell after me.
func (me *Cell) addAfter(after *Cell) {
	after.next = me.next
	after.prev = me
	me.next.prev = after
	me.next = after
}

// Adds a cell before current cell
func (me *Cell) addBefore(before *Cell) {
	// adding before is like adding after to the previous cell
	if me.prev != nil {
		me.prev.addAfter(before)
	} else {
		me.addAfter(before)
	}
}

// Deletes current cell
func (me *Cell) delete() *Node {
	me.prev.next = me.next
	me.next.prev = me.prev
	return me.data
}

func (list *DoublyLinkedList) length() int {
	cur_cell := list.topSentinel
	length := 0
	for {
		if cur_cell.next == list.bottomSentinel {
			break
		}
		length += 1
		cur_cell = cur_cell.next
	}
	return length
}

func (list *DoublyLinkedList) isEmpty() bool {
	cur_cell := list.topSentinel
	if cur_cell.next == list.bottomSentinel {
		return true
	}
	return false
}

func (list *DoublyLinkedList) contains(target *Node) bool {
	cur_cell := list.topSentinel
	for {
		if cur_cell.next == list.bottomSentinel {
			break
		}
		cur_cell = cur_cell.next
		if cur_cell.data == target {
			return true
		}
	}
	return false
}

func (list *DoublyLinkedList) push(new_data *Node) {
	cur_cell := list.topSentinel
	new_cell := Cell{data: new_data}
	cur_cell.addAfter(&new_cell)
}

func (list *DoublyLinkedList) pop() (*Node, error) {
	if list.isEmpty() {
		err := errors.New("the list is empty, cannot pop item")
		return nil, err
	}
	cur_cell := list.topSentinel
	popped := cur_cell.delete()
	return popped, nil
}

// Add an item to the top of the queue.
func (queue *DoublyLinkedList) enqueue(value *Node) {
	queue.push(value)
}

// Remove an item from the bottom of the queue.
func (queue *DoublyLinkedList) dequeue() (*Node, error) {
	if !queue.isEmpty() {
		data := queue.bottomSentinel.prev.delete()
		return data, nil
	} else {
		return nil, errors.New("queue is empty")
	}
}

// Add an item at the bottom of the deque.
func (deque *DoublyLinkedList) pushBottom(value *Node) {
	new_cell := Cell{data: value}
	deque.bottomSentinel.addBefore(&new_cell)
}

// Add an item at the top of the deque.
func (deque *DoublyLinkedList) pushTop(value *Node) {
	deque.push(value)
}

// Remove an item from the top of the deque.
func (deque *DoublyLinkedList) popTop() (*Node, error) {
	return deque.pop()
}

// Delete from the end of the queue
func (deque *DoublyLinkedList) popBottom() (*Node, error) {
	return deque.dequeue()
}

// Makes a queue for the BFS
func makeDoublyLinkedList() DoublyLinkedList {
	topSentinel := Cell{data: nil, next: nil}
	bottomSentinel := Cell{data: nil, next: nil}
	topSentinel.next = &bottomSentinel
	bottomSentinel.prev = &topSentinel
	new_list := DoublyLinkedList{topSentinel: &topSentinel, bottomSentinel: &bottomSentinel}
	return new_list
}

// ============================ Trees ===============================================================================================
type Node struct {
	data  string
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func buildTree() *Node {
	aNode := Node{"A", nil, nil}
	bNode := Node{"B", nil, nil}
	cNode := Node{"C", nil, nil}
	dNode := Node{"D", nil, nil}
	eNode := Node{"E", nil, nil}
	fNode := Node{"F", nil, nil}
	gNode := Node{"G", nil, nil}
	hNode := Node{"H", nil, nil}
	iNode := Node{"I", nil, nil}
	jNode := Node{"J", nil, nil}
	aNode.left = &bNode
	aNode.right = &cNode
	bNode.left = &dNode
	bNode.right = &eNode
	eNode.left = &gNode
	cNode.right = &fNode
	fNode.left = &hNode
	hNode.left = &iNode
	hNode.right = &jNode

	return &aNode
}

// displays the tree given the starting node
func (node *Node) displayIndented(indent string, depth int) string {
	result := ""
	for d := 0; d <= depth; d++ {
		result += indent
	}
	result += node.data
	result += ":\n"
	if (node.right != nil) && (node.left != nil) {
		result = result + node.right.displayIndented(indent, depth+1) + node.left.displayIndented(indent, depth+1)
	} else if node.right != nil {
		result = result + node.right.displayIndented(indent, depth+1)
	} else if node.left != nil {
		result = result + node.left.displayIndented(indent, depth+1)
	} else {
		result = result
	}
	return result
}

func (node *Node) preorder() string {
	result := ""
	result += node.data + " "
	if (node.right != nil) && (node.left != nil) {
		result = result + node.left.preorder() + node.right.preorder()
	} else if node.right != nil {
		result = result + node.right.preorder()
	} else if node.left != nil {
		result = result + node.left.preorder()
	} else {
		result = result
	}
	return result
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

func (node *Node) postorder() string {
	result := ""
	if (node.right != nil) && (node.left != nil) {
		result += node.left.postorder() + node.right.postorder() + node.data + " "
	} else if node.right != nil {
		result += node.right.postorder() + node.data + " "
	} else if node.left != nil {
		result += node.left.postorder() + node.data + " "
	} else {
		result += node.data + " "
	}
	return result
}

func (node *Node) breadthFirst() string {
	queue := makeDoublyLinkedList()
	queue.enqueue(node)
	result := ""
	for {
		if queue.isEmpty() {
			break
		}
		nd, _ := queue.dequeue()
		result += nd.data + " "
		if nd.left != nil {
			queue.enqueue(nd.left)
		}
		if nd.right != nil {
			queue.enqueue(nd.right)
		}

	}
	return result
}

func main() {
	// Build a tree.
	aNode := buildTree()

	// Display with indentation.
	fmt.Println(aNode.displayIndented("  ", 0))

	// Preorder
	fmt.Println("Preorder:     ", aNode.preorder())

	// Inorder
	fmt.Println("Inorder:     ", aNode.inorder())

	// Postorder
	fmt.Println("Postorder:     ", aNode.postorder())

	// BFS
	fmt.Println("Breadth first:", aNode.breadthFirst())
}
