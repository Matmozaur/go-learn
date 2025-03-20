package main

import (
	"container/list"
	"fmt"
)

// Person struct with pointer field
type Person struct {
	Name string
	Age  int
}

// Function to modify age using pointer
func updateAge(p *Person, newAge int) {
	p.Age = newAge
}

// Stack implementation using slice
type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		panic("Stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Queue implementation using list package
type Queue struct {
	items *list.List
}

func NewQueue() *Queue {
	return &Queue{items: list.New()}
}

func (q *Queue) Enqueue(item int) {
	q.items.PushBack(item)
}

func (q *Queue) Dequeue() int {
	if q.items.Len() == 0 {
		panic("Queue is empty")
	}
	element := q.items.Front()
	q.items.Remove(element)
	return element.Value.(int)
}

func main() {
	// Working with Arrays
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)

	// Working with Slices
	slice := []int{10, 20, 30}
	slice = append(slice, 40)
	fmt.Println("Slice after append:", slice)

	// Working with Maps
	myMap := make(map[string]int)
	myMap["Alice"] = 25
	myMap["Bob"] = 30
	fmt.Println("Map:", myMap)

	// Accessing a value in map
	age, exists := myMap["Alice"]
	if exists {
		fmt.Println("Alice's age:", age)
	} else {
		fmt.Println("Alice not found in map")
	}

	// Using Structs and Pointers
	p1 := Person{Name: "John", Age: 28}
	fmt.Println("Before update:", p1)

	updateAge(&p1, 35)
	fmt.Println("After update:", p1)

	// Working with Pointers
	var num int = 42
	var ptr *int = &num
	fmt.Println("Value of num:", num)
	fmt.Println("Pointer to num:", ptr)
	fmt.Println("Value through pointer:", *ptr)

	// Using Stack
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println("Stack Pop:", stack.Pop())
	fmt.Println("Stack Pop:", stack.Pop())

	// Using Queue
	queue := NewQueue()
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	fmt.Println("Queue Dequeue:", queue.Dequeue())
	fmt.Println("Queue Dequeue:", queue.Dequeue())
}
