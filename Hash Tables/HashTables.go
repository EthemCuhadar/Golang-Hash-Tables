package main

import (
	"fmt"
)

// Arrays size refers to bucket size in Hash table.
const ArraySize = 7

// HashTable struct which holds array for buckets.
type HashTable struct {
	array	[ArraySize]*bucket
}

// bucket is a linked list which stores the "values"
type bucket struct{
	head	*bucketNode
}

// bucketNode - 
type bucketNode struct{
	key	string
	next	*bucketNode
}

// Insert

// Delete

// Search


// insert

// search

// delete

// hash function

// Init will create a bucket in each slot of the hash table
func Init() *HashTable{
	result := &HashTable{}
	for i := range result.array{
	result.array[i] = &bucket{}
	}
	return result
}

func main(){
	fmt.Println("OK Let's Go!")
	testHashTable := Init()
	fmt.Println(testHashTable)
}
