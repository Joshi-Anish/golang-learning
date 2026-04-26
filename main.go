package main

import "fmt"

func main() {

	fmt.Println("====================================")
	fmt.Println("GO BACKEND JOURNEY")
	fmt.Println("====================================")

	fmt.Println("Started Learning Go: April 20")

	goals := []string{
		"Building Backend Applications with Golang",
		"Learning REST APIs & System Design",
		"Writing Clean and Efficient Code",
		"Understanding Secure Backend Practices",
	}

	fmt.Println("\nCurrent Learning Focus:")

	for _, goal := range goals {
		fmt.Printf("  • %s\n", goal)
	}

	fmt.Println("\nProgress Status:")
	fmt.Println("  → Learning by building projects")
	fmt.Println("  → Documenting everything on GitHub")
	fmt.Println("  → Improving every single day")

	fmt.Println("\nMission:")
	fmt.Println("  Build fast, scalable, and secure systems.")

	fmt.Println("\nGitHub Learning Journey:")
	fmt.Println("  github.com/Joshi-Anish/golang-learning")

	fmt.Println("\n====================================")
	fmt.Println("Keep Building. Keep Learning.")
	fmt.Println("====================================")
}
