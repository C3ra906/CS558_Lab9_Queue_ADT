//Cera Oh. CS558 Winter 2022. Lab 9
//Ring Implementation
package queues


import (
	"fmt"
)

const N int = 5 

type RQueue [T any] struct{
	values [N+1]T
	front int
	rear int
}

func REmpty[T any]() *RQueue[T] {
    q := new (RQueue[T])
    counter := 0

	var result T
    for counter < N+1 {
        q.values[counter] = result //fill with empty values
        counter++
	}
	
    q.front = 0
    q.rear = 0
    return q
}

func (q *RQueue[T]) IsEmpty()  bool {
	return q.front == q.rear
}

func (q *RQueue[T]) Dequeue() T {
	if (q.front == q.rear){
		var result T
		return result
	} else {
	    v := q.values[q.front]
	    q.front = (q.front+1)%(N+1)
	    return v
	}
}

func (q *RQueue[T]) Enqueue (v T){
	if ((q.rear + 1)%(N+1) == q.front) {
		fmt.Println("Cannot enqueue. Ran out of room!") //error: out of room
		return
	} else{
		q.values[q.rear] = v
 		q.rear = (q.rear+1)%(N+1)
	}
}



