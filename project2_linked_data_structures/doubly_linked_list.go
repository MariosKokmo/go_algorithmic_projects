package main

import (
	"errors"
	"fmt"
)

type Cell struct {
	data string
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
func (me *Cell) delete() string {
	me.prev.next = me.next
	me.next.prev = me.prev
	return me.data
}

// Adds a range of values at the end
func (list *DoublyLinkedList) addRange(values []string) {
	// it can add before the bottomSentinel directly
	for _, value := range values {
		new_cell := Cell{data: value}
		list.bottomSentinel.addBefore(&new_cell)
	}
	return
}

// Prints the list as a string with the separator
// betweeen values
func (list *DoublyLinkedList) toString(separator string) string {
	var result string
	cur_cell := list.topSentinel
	if cur_cell.next == list.bottomSentinel {
		return ""
	}
	cur_cell = cur_cell.next
	result = cur_cell.data
	for {
		if cur_cell.next == list.bottomSentinel {
			break
		}
		cur_cell = cur_cell.next
		result = result + separator + cur_cell.data
	}
	return result
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

func (list *DoublyLinkedList) contains(target string) bool {
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

func (list *DoublyLinkedList) push(new_data string) {
	cur_cell := list.topSentinel
	new_cell := Cell{data: new_data}
	cur_cell.addAfter(&new_cell)
}

func (list *DoublyLinkedList) pop() (string, error) {
	if list.isEmpty() {
		err := errors.New("the list is empty, cannot pop item")
		return "", err
	}
	cur_cell := list.topSentinel
	popped := cur_cell.delete()
	return popped, nil
}

// Add an item to the top of the queue.
func (queue *DoublyLinkedList) enqueue(value string) {
	queue.push(value)
}

// Remove an item from the bottom of the queue.
func (queue *DoublyLinkedList) dequeue() (string, error) {
	if !queue.isEmpty() {
		data := queue.bottomSentinel.prev.delete()
		return data, nil
	} else {
		return "", errors.New("queue is empty")
	}
}

// Add an item at the bottom of the deque.
func (deque *DoublyLinkedList) pushBottom(value string) {
	new_cell := Cell{data: value}
	deque.bottomSentinel.addBefore(&new_cell)
}

// Add an item at the top of the deque.
func (deque *DoublyLinkedList) pushTop(value string) {
	deque.push(value)
}

// Remove an item from the top of the deque.
func (deque *DoublyLinkedList) popTop() (string, error) {
	return deque.pop()
}

// Delete from the end of the queue
func (deque *DoublyLinkedList) popBottom() (string, error) {
	return deque.dequeue()
}

func makeDoublyLinkedList() DoublyLinkedList {
	topSentinel := Cell{data: "topSentinel", next: nil}
	bottomSentinel := Cell{data: "bottomSentinel", next: nil}
	topSentinel.next = &bottomSentinel
	bottomSentinel.prev = &topSentinel
	new_list := DoublyLinkedList{topSentinel: &topSentinel, bottomSentinel: &bottomSentinel}
	return new_list
}

func test_DoublyLinkedList() {
	// Make a list from a slice of values.
	list := makeDoublyLinkedList()
	animals := []string{
		"Ant",
		"Bat",
		"Cat",
		"Dog",
		"Elk",
		"Fox",
	}
	list.addRange(animals)
	fmt.Println(list.toString(" "))
}

func main() {
	// Test queue functions.
	fmt.Printf("*** Queue Functions ***\n")
	queue := makeDoublyLinkedList()
	queue.enqueue("Agate")
	queue.enqueue("Beryl")
	data, _ := queue.dequeue()
	fmt.Printf("%s ", data)
	queue.enqueue("Citrine")
	data, _ = queue.dequeue()
	fmt.Printf("%s ", data)
	data, _ = queue.dequeue()
	fmt.Printf("%s ", data)
	queue.enqueue("Diamond")
	queue.enqueue("Emerald")
	for !queue.isEmpty() {
		data, _ = queue.dequeue()
		fmt.Printf("%s ", data)
	}
	fmt.Printf("\n\n")

	// Test deque functions. Names starting
	// with F have a fast pass.
	fmt.Printf("*** Deque Functions ***\n")
	deque := makeDoublyLinkedList()
	deque.pushTop("Ann")
	deque.pushTop("Ben")
	data, _ = deque.popBottom()
	fmt.Printf("%s ", data)
	deque.pushBottom("F-Cat")
	data, _ = deque.popBottom()
	fmt.Printf("%s ", data)
	data, _ = deque.popBottom()
	fmt.Printf("%s ", data)
	deque.pushBottom("F-Dan")
	deque.pushTop("Eva")
	for !deque.isEmpty() {
		data, _ = deque.popBottom()
		fmt.Printf("%s ", data)
	}
	fmt.Printf("\n")
}
