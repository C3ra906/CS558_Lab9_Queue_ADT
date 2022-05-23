package main

import (
	"fmt"

	"ceoh.com/lab9/queues"
)

func main() {
	fmt.Println("Hello, Welcome to the program")
	fmt.Println("Creating a linked list queue of ints:")
	var li *queues.Queue[int] = queues.LEmpty[int]()
	fmt.Println("Done")
	fmt.Println("Adding: 5, -1, and 2 to the queue.")
	li.Enqueue(5)
	li.Enqueue(-1)
	li.Enqueue(2)
	fmt.Println("Done! Dequeuing all and printing values:")
	for (li.IsEmpty() != true){
			fmt.Println(li.Dequeue())
	}
	if (li.IsEmpty() == true){
		fmt.Println("The queue is now empty")
	}
    fmt.Println("\n")

	fmt.Println("Creating a linked list queue of strings:")
	var ls *queues.Queue[string] = queues.LEmpty[string]()
	fmt.Println("Done")
	fmt.Println("Adding: a, b, and c to the queue.")
	ls.Enqueue("a")
	ls.Enqueue("b")
	ls.Enqueue("c")
	fmt.Println("Done! Dequeuing all and printing values:")
	for (ls.IsEmpty() != true){
			fmt.Println(ls.Dequeue())
	}
	if (ls.IsEmpty() == true){
		fmt.Println("The queue is now empty")
	}
    fmt.Println("\n")

	fmt.Println("Creating a ring queue of ints:")
	var ri *queues.RQueue[int] = queues.REmpty[int]()
	fmt.Println("Done")
	fmt.Println("Adding: 5, 7, and 9 to the queue.")
	ri.Enqueue(5)
	ri.Enqueue(7)
	ri.Enqueue(9)
	fmt.Println("Done! Dequeuing all and printing values:")
	for (ri.IsEmpty() != true){
			fmt.Println(ri.Dequeue())
	}
	if (ri.IsEmpty() == true){
		fmt.Println("The queue is now empty")
	}
    fmt.Println("\n")
	
	fmt.Println("Creating a ring queue of strings:")
	var rs *queues.RQueue[string] = queues.REmpty[string]()
	fmt.Println("Done")
	fmt.Println("Adding: abe, dan, and frank to the queue.")
	rs.Enqueue("abe")
	rs.Enqueue("dan")
	rs.Enqueue("frank")
	fmt.Println("Done! Dequeuing all and printing values:")
	for (rs.IsEmpty() != true){
			fmt.Println(rs.Dequeue())
	}
	if (rs.IsEmpty() == true){
		fmt.Println("The queue is now empty")
	}
	
}
