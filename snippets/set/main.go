package main

import (
	"fmt"

	sl "codigo/data-structures/set"
)

func main() {
	set := sl.NewListSet(1, 10, 5)
	fmt.Println(set)
	fmt.Println("Agregar 7")
	set.Add(7)
	fmt.Println(set)
	fmt.Println("set.Contains(7) -> ", set.Contains(7))
	fmt.Println("Eliminamos el 10")
	set.Remove(10)
	fmt.Println("set.Contains(10) -> ", set.Contains(10))
	fmt.Println(set)
	fmt.Println("set.Values() -> ", set.Values())
	fmt.Println("Agregar 4 y 12")
	set.Add(4, 12)
	fmt.Println(set)
}