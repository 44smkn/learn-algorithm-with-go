package datastructure

import (
	"strconv"
	"strings"
)

type node struct {
	key        int
	prev, next *node
}

func printList(instructions []string) []int {
	for _, instruction := range instructions {
		commands := strings.Split(instruction, " ")
		kind := commands[0]
		key, _ := strconv.Atoi(commands[1])
		switch kind {
		case "insert":
			insert(key)
		case "delete":
			deleteKey(key)
		case "deleteFirst":
			deleteFirst()
		case "deleteLast":
			deleteLast()
		}
	}

	cur := null
	output := make([]int, 0, 10)
	for cur.next != null {
		cur = cur.next
		output = append(output, cur.key)
	}
	return output
}

// 番兵
var null = &node{}

func init() {
	null.prev = null
	null.next = null
}

func insert(key int) {
	x := &node{
		key:  key,
		next: null.next,
		prev: null,
	}
	null.next.prev = x
	null.next = x
}

func listSearch(key int) *node {
	cur := null.next
	for cur != null && cur.key != key {
		cur = cur.next
	}
	return cur
}

func deleteNode(t *node) {
	if t == null {
		return
	}
	t.next.prev = t.prev
	t.prev.next = t.next
}

func deleteFirst() {
	deleteNode(null.next)
}

func deleteLast() {
	deleteNode(null.prev)
}

func deleteKey(key int) {
	deleteNode(listSearch(key))
}
