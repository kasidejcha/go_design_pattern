package main

import (
	"fmt"
	"sync"
)

// Singleton is the structure that we want to instantiate only once
type Singleton struct {
	value string
}

var (
	instance *Singleton
	once     sync.Once
)

// GetInstance provides global access to the Singleton instance
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			value: "This is the Singleton instance",
		}
	})
	return instance
}

func main() {
	// Get the first instance
	s1 := GetInstance()
	fmt.Println(s1.value)

	// Get the second instance
	s2 := GetInstance()
	fmt.Println(s2.value)

	// Check if both instances point to the same memory address
	fmt.Printf("Are s1 and s2 the same instance? %v\n", s1 == s2)
}
