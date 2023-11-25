package main

import "fmt"

func main() {
	john := NewUser("John Doe")
	jane := NewUser("Jane Doe")
	alsoJane := NewUser("Jane Smith")
	fmt.Println(john.FullName)
	fmt.Println(jane.FullName)
	fmt.Println(alsoJane.FullName)
	fmt.Println("Memory taken by users:",
		len([]byte(john.FullName))+
			len([]byte(jane.FullName))+
			len([]byte(alsoJane.FullName)))

	john2 := NewUser2("John Doe")
	jane2 := NewUser2("Jane Doe")
	alsoJane2 := NewUser2("Jane Smith")
	totalMem := 0
	for _, a := range allNames {
		totalMem += len([]byte(a))
	}

	totalMem += len(john2.names)
	totalMem += len(jane2.names)
	totalMem += len(alsoJane2.names)
	fmt.Println("Memory taken by users2", totalMem)
}
