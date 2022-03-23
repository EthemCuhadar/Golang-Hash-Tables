package main

import (
	"fmt"
)

// Arrays size refers to bucket size in Hash table.
const ArraySize = 7

// HashTable struct which holds array for buckets.
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket is a linked list which stores the "values"
type bucket struct {
	head *bucketNode
}

// bucketNode -
type bucketNode struct {
	key  string
	next *bucketNode
}

// For Insert, Delete and Search operations firstly key will go in
// the hash function.

// Insert will take and append the element in the tash table.
func (ht *HashTable) Insert(key string) {
	index := hash(key)
	ht.array[index].insert(key)
}

// Delete will take a key and remove it from the hash table.
func (ht *HashTable) Delete(key string) {
	index := hash(key)
	ht.array[index].delete(key)
}

// Search will simply check whether the key is stored in the hash
// table or not. It returns true if the key is stored.
func (ht *HashTable) Search(key string) bool {
	index := hash(key)
	return ht.array[index].search(key)
}

// insert will take a key, create a node with that key and insert the
// node into the bucket.
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("Key already exists!")
	}

}

// search will take a key and returns true if it exits in the bucket.
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete will take in a key and delete the node from bucket.
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode != nil {
		if previousNode.next.key == k {
			// delete
			previousNode.next = previousNode.next.next
		}
	}
}

// hash function takes a key and returns a integer. The returned number
// should be modded with ArraySize.
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	myHashTable := Init()
	list := []string{
		"england",
		"france",
		"germany",
		"italy",
		"netherlands",
		"greece",
		"turkey",
	}

	for _, v := range list {
		myHashTable.Insert(v)
	}

	myHashTable.Delete("italy")
	myHashTable.Delete("germany")
	myHashTable.Delete("england")

	fmt.Println(myHashTable.Search("italy"))
	fmt.Println(myHashTable.Search("france"))
	fmt.Println(myHashTable.Search("england"))
	fmt.Println(myHashTable.Search("turkey"))

}
