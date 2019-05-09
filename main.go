package main

import "fmt"
import "./immutable_queue"

func main() {
	fmt.Println("start with {0, 10, 100}")

	items := []interface{}{0, 10, 100}
	iq := immutable_queue.New(items)

	fmt.Printf("iq.IsEmpty() is: %t\n", iq.IsEmpty())
	fmt.Printf("iq.Head() is: %#v\n", iq.Head())

	fmt.Println("iq.DeQueue()")
	iq = iq.DeQueue()
	fmt.Printf("iq.Head() is: %#v\n", iq.Head())

	fmt.Println("iq.DeQueue()")
	iq = iq.DeQueue()
	fmt.Printf("iq.Head() is: %#v\n", iq.Head())

	fmt.Println("iq.DeQueue()")
	iq = iq.DeQueue()
	fmt.Printf("iq.Head() is: %#v\n", iq.Head())
	fmt.Printf("iq.IsEmpty() is: %t\n", iq.IsEmpty())

	fmt.Println("iq.DeQueue()")
	iq = iq.DeQueue()
	fmt.Printf("iq.Head() is: %#v\n", iq.Head())

	fmt.Println("iq.EnQueue(2)")
	iq = iq.EnQueue(2)
	fmt.Printf("iq.Tail() is: %#v\n", iq.Tail())

	fmt.Println("iq.EnQueue(2)")
	iq = iq.EnQueue(22)
	fmt.Printf("iq.Tail() is: %#v\n", iq.Tail())
}
