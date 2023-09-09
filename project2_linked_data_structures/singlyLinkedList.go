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
	// smallListTest()

	// Make a list from an array of values.
	greekLetters := []string{
		"α", "β", "γ", "δ", "ε",
	}
	list := makeLinkedList()
	list.addRange(greekLetters)
	fmt.Println(list.toString(" "))
	fmt.Println()

	//Demonstrate a stack.
	stack := makeLinkedList()
	stack.push("Apple")
	stack.push("Banana")
	stack.push("Coconut")
	stack.push("Date")
	for !stack.isEmpty() {
		data, _ := stack.pop()
		fmt.Printf("Popped: %-7s   Remaining %d: %s\n",
			data,
			stack.length(),
			stack.toString(" "))
	}
}
