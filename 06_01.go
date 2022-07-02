package main

import "fmt"

type node struct {
	val int
	nextNode *node
}

// ДЗ 06-01 Разворот односвязного списка
func main () {

	nodeThree := node{val:0,nextNode:nil}
	nodeTwo := node{val:0,nextNode:&nodeThree}
	nodeOne := node{val:0,nextNode:&nodeTwo}

	fmt.Println("Введите значение первого элемента")
	fmt.Scanln(&nodeOne.val)
	fmt.Println("Введите значение второго элемента")
	fmt.Scanln(&nodeTwo.val)
	fmt.Println("Введите значение третьего элемента")
	fmt.Scanln(&nodeThree.val)

	listPrint(&nodeOne)
	listSwap(&nodeOne)
	fmt.Println("Y ahora! El Swapped List!")
	listPrint(&nodeThree)

}

// Вывод списка
func listPrint (current *node) {
	if current == nil {
		return
	}

	fmt.Println(current.val)
	listPrint(current.nextNode)
}

// Разворот списка
func listSwap (current *node) {
	var prev, cur, next *node
	prev = nil
	cur = current
	next = cur.nextNode

	for ;; {
		next = cur.nextNode
		cur.nextNode = prev
		prev = cur
		cur = next
		if next == nil {
			return
		}
	}
}