package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type node struct {
	value int
	next  *node
}

func main() {
	fmt.Printf("Part 1: %v\n,", part1(initializeInput()))
	fmt.Printf("Part 2: %v\n", part2(initializeInput()))
}

func initializeInput() (*node, int, map[int]*node) {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	head := &node{}
	current := head
	nodesMap := make(map[int]*node)

	var length int
	for _, i := range string(data) {
		current.value = int(i - '0')
		nodesMap[current.value] = current
		current.next = &node{}
		current = current.next
		length++
	}
	current.value = head.value
	current.next = head.next
	head = current
	nodesMap[head.value] = head
	return head, length, nodesMap
}

func part2(head *node, length int, nodesMap map[int]*node) int {
	current := head
	for i := 1; i < length; i++ {
		current = current.next
	}
	current.next = &node{}
	current = current.next
	for i := length + 1; i <= 1000000; i++ {
		current.value = i
		nodesMap[current.value] = current
		current.next = &node{}
		current = current.next
		length++
	}
	current.value = head.value
	current.next = head.next
	head = current
	nodesMap[head.value] = head

	for i := 0; i < 10000000; i++ {
		head = move(head, length, nodesMap)
	}

	current = nodesMap[1].next
	return current.value * current.next.value
}

func move(current *node, length int, nodesMap map[int]*node) *node {
	picked := current.next
	current.next = picked.next.next.next
	pickedMap := make(map[int]bool)
	copyNode := picked
	for i := 0; i < 3; i++ {
		pickedMap[copyNode.value] = true
		copyNode = copyNode.next
	}
	putAt := current.value - 1
	for {
		if putAt == 0 {
			putAt = length
		}
		_, alreadyPicked := pickedMap[putAt]
		if alreadyPicked {
			putAt--
		} else {
			break
		}
	}
	currentNode := nodesMap[putAt]
	picked.next.next.next = currentNode.next
	currentNode.next = picked
	return current.next
}

func part1(head *node, length int, nodesMap map[int]*node) string {
	for i := 0; i < 100; i++ {
		head = move(head, length, nodesMap)
	}
	for head.value != 1 {
		head = head.next
	}
	head = head.next

	var ans string
	for i := 0; i < length-1; i++ {
		ans += fmt.Sprint(head.value)
		head = head.next
	}
	return ans
}
