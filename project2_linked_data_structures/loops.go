package main

import (
	"errors"
	"fmt"
)

type Cell struct {
	data string
	next *Cell
}

type LinkedList struct {
	sentinel *Cell
}

// Add a cell after me.
func (me *Cell) addAfter(after *Cell) {
	after.next = me.next
	me.next = after
}

func (me *Cell) deleteAfter() *Cell {
	if me.next == nil {
		panic("there is nothing to delete after")
	}
	cell_to_delete := me.next
	me.next = cell_to_delete.next
	cell_to_delete.next = nil
	return cell_to_delete
}

func (list *LinkedList) addRange(values []string) {
	// find last cell in list
	lastCell := list.sentinel
	for {
		if lastCell.next != nil {
			lastCell = lastCell.next
		} else {
			break
		}
	}
	// add the new cells
	for _, value := range values {
		new_cell := Cell{data: value, next: nil}
		lastCell.addAfter(&new_cell)
		lastCell = &new_cell
	}
	return
}

// Prints the list as a string with the separator
// betweeen values
func (list *LinkedList) toString(separator string) string {
	var result string
	cur_cell := list.sentinel
	if cur_cell.next == nil {
		return ""
	}
	cur_cell = cur_cell.next
	result = cur_cell.data
	for {
		if cur_cell.next == nil {
			break
		}
		cur_cell = cur_cell.next
		result = result + separator + cur_cell.data
	}
	return result
}

func (list *LinkedList) length() int {
	cur_cell := list.sentinel
	length := 0
	for {
		if cur_cell.next == nil {
			break
		}
		length += 1
		cur_cell = cur_cell.next
	}
	return length
}

func (list *LinkedList) isEmpty() bool {
	cur_cell := list.sentinel
	if cur_cell.next == nil {
		return true
	}
	return false
}

func (list *LinkedList) contains(target string) bool {
	cur_cell := list.sentinel
	for {
		if cur_cell.next == nil {
			break
		}
		cur_cell = cur_cell.next
		if cur_cell.data == target {
			return true
		}
	}
	return false
}

// hasLoop checks for loops in a singly linked list
// using the Tortoise-Hare algorithm
func (list *LinkedList) hasLoop() bool {
	fast := list.sentinel
	slow := list.sentinel
	for {
		if fast.next != nil && fast.next.next != nil {
			fast = fast.next.next
		} else {
			return false
		}
		if slow.next != nil {
			slow = slow.next
		}
		if slow == fast {
			return true
		}
	}
	return false
}

// Prints a maximum number of items in the linked list.
// Thus, avoiding infinite loops
func (list *LinkedList) toStringMax(separator string, max int) string {
	var result string
	var count int
	cur_cell := list.sentinel
	if cur_cell.next == nil {
		return ""
	}
	cur_cell = cur_cell.next
	result = cur_cell.data
	for {
		if cur_cell.next == nil || count == max-1 {
			break
		}
		cur_cell = cur_cell.next
		count += 1
		result = result + separator + cur_cell.data
	}
	return result
}

func (list *LinkedList) push(new_data string) {
	cur_cell := list.sentinel
	new_cell := Cell{data: new_data}
	cur_cell.addAfter(&new_cell)
}

func (list *LinkedList) pop() (string, error) {
	if list.isEmpty() {
		err := errors.New("the list is empty, cannot pop item")
		return "", err
	}
	cur_cell := list.sentinel
	popped := cur_cell.deleteAfter()
	return popped.data, nil
}

func makeLinkedList() LinkedList {
	sentinel := Cell{data: "sentinel", next: nil}
	new_list := LinkedList{sentinel: &sentinel}
	return new_list
}

func main() {
	// Make a list from an array of values.
	values := []string{
		"0", "1", "2", "3", "4", "5",
	}
	list := makeLinkedList()
	list.addRange(values)

	fmt.Println(list.toString(" "))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
	fmt.Println()

	// Make cell 5 point to cell 2.
	list.sentinel.next.next.next.next.next.next = list.sentinel.next.next

	fmt.Println(list.toStringMax(" ", 10))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
	fmt.Println()

	// Make cell 4 point to cell 2.
	list.sentinel.next.next.next.next.next = list.sentinel.next.next

	fmt.Println(list.toStringMax(" ", 10))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
}
