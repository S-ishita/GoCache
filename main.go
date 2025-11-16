package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

const SIZE = 5

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Hash map[string]*Node

func main() {
	fmt.Println("START CACHE")
	cache := NewCache()
	//initial cache
	for _, word := range []string{"firewall", "encryption", "loadbalancer", "container","database", "microservice", "container", "endpoint"} {
		cache.Check(word)
		cache.Display()
	}

	// Interactive cache
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Do you want to input a new element? (yes/no): ")
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))

		if response == "no" {
			fmt.Println("Exiting...")
			break
		} else if response == "yes" {
			fmt.Print("Enter new element: ")
			word, _ := reader.ReadString('\n')
			word = strings.TrimSpace(word)
			cache.Check(word)
			cache.Display()
	} else {
			fmt.Println("Please enter 'yes' or 'no'.")
		}
	}
}
