package main

import "fmt"

var employeeNames [100]string

type Employee struct {
	name  string
	phone string
}

type ChainingHashTable struct {
	numBuckets int
	buckets    [][]*Employee
}

// Initialize a ChainingHashTable and return a pointer to it.
func NewChainingHashTable(numBuckets int) *ChainingHashTable {
	buck := make([][]*Employee, numBuckets)
	new_hash_table := ChainingHashTable{numBuckets: numBuckets, buckets: buck}
	return &new_hash_table
}

// Display the hash table's contents.
func (hashTable *ChainingHashTable) dump() {
	for index, bucket := range hashTable.buckets {
		fmt.Println("Bucket ", index, ":")
		for _, employee := range bucket {
			fmt.Println("\t", employee.name, ": ", employee.phone)
		}
	}
}

// Find the bucket and Employee holding this key.
// Return the bucket number and Employee number in the bucket.
// If the key is not present, return the bucket number and -1.
func (hashTable *ChainingHashTable) find(name string) (int, int) {
	hash_value := hash(name) % hashTable.numBuckets
	for index, empl := range hashTable.buckets[hash_value] {
		if empl.name == name {
			return hash_value, index
		}
	}
	return hash_value, -1
}

// Add an item to the hash table.
func (hashTable *ChainingHashTable) set(name string, phone string) {
	bucket, index := hashTable.find(name)
	if index >= 0 {
		hashTable.buckets[bucket][index].phone = phone
	} else {
		newEmployee := Employee{name: name, phone: phone}
		hashTable.buckets[bucket] = append(hashTable.buckets[bucket], &newEmployee)
	}
}

// Return an item from the hash table.
func (hashTable *ChainingHashTable) get(name string) string {
	bucket, index := hashTable.find(name)
	if index >= 0 {
		return hashTable.buckets[bucket][index].phone
	} else {
		return ""
	}
}

// Return true if the person is in the hash table.
func (hashTable *ChainingHashTable) contains(name string) bool {
	_, index := hashTable.find(name)
	if index == -1 {
		return false
	}
	return true
}

// Delete this key's entry.
func (hashTable *ChainingHashTable) delete(name string) {
	bucket, index := hashTable.find(name)
	if index >= 0 {
		hashTable.buckets[bucket] = hashTable.buckets[bucket][:index]
	}
}

// djb2 hash function. See http://www.cse.yorku.ca/~oz/hash.html.
func hash(value string) int {
	hash := 5381
	for _, ch := range value {
		hash = ((hash << 5) + hash) + int(ch)
	}

	// Make sure the result is non-negative.
	if hash < 0 {
		hash = -hash
	}
	return hash
}

func main() {
	// Make some names.
	employees := []Employee{
		Employee{"Ann Archer", "202-555-0101"},
		Employee{"Bob Baker", "202-555-0102"},
		Employee{"Cindy Cant", "202-555-0103"},
		Employee{"Dan Deever", "202-555-0104"},
		Employee{"Edwina Eager", "202-555-0105"},
		Employee{"Fred Franklin", "202-555-0106"},
		Employee{"Gina Gable", "202-555-0107"},
		Employee{"Herb Henshaw", "202-555-0108"},
		Employee{"Ida Iverson", "202-555-0109"},
		Employee{"Jeb Jacobs", "202-555-0110"},
	}

	hashTable := NewChainingHashTable(10)
	for _, employee := range employees {
		hashTable.set(employee.name, employee.phone)
	}
	hashTable.dump()

	fmt.Printf("Table contains Sally Owens: %t\n", hashTable.contains("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", hashTable.contains("Dan Deever"))
	fmt.Println("Deleting Dan Deever")
	hashTable.delete("Dan Deever")
	fmt.Printf("Table contains Dan Deever: %t\n", hashTable.contains("Dan Deever"))
	fmt.Printf("Sally Owens: %s\n", hashTable.get("Sally Owens"))
	fmt.Printf("Fred Franklin: %s\n", hashTable.get("Fred Franklin"))
	fmt.Println("Changing Fred Franklin")
	hashTable.set("Fred Franklin", "202-555-0100")
	fmt.Printf("Fred Franklin: %s\n", hashTable.get("Fred Franklin"))
}
