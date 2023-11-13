package intermediate

import (
	"fmt"
)

type HashMap map[string]int

func (hash *HashMap) Insert(key string, data int) {
	(*hash)[key] = data
}

func (hash *HashMap) Retrieve(key string) {
	value := (*hash)[key]
	fmt.Println(value)
}

func HashMapChallenge() {

	// Challenge 7: Hash Map Implementation
	// Title: Challenge 7: Hash Map Implementation
	// Description: Create a basic hash map and perform insert, retrieve, or delete operations on key-value pairs.

	// Input:
	// hash_map := HashMap{}
	// hash_map.insert("a", 1)
	// hash_map.insert("b", 2)
	// hash_map.retrieve("a")

	// Desired Output:
	// 1

	// Incorrect Output:
	// 2

	hash_map := HashMap{}
	hash_map.Insert("a", 1)
	hash_map.Insert("b", 2)
	hash_map.Retrieve("a")

}
