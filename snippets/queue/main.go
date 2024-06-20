package main

import (
	"fmt"
	"github.com/trigologiaa/codigo/data-structures/queue"
)

func main() {
	q := queue.New[int]()
	fmt.Println("q.IsEmpty() -> ", q.IsEmpty())
	fmt.Println("Encolando... 1, 2 y 3")
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println("q.IsEmpty() -> ", q.IsEmpty())
	fmt.Println("Desencolando...")
	for !q.IsEmpty() {
		item, _ := q.Dequeue()
		fmt.Println("Desencoló ->", item)
	}
	fmt.Println("q.IsEmpty() -> ", q.IsEmpty())
}