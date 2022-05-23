//Cera Oh. CS558 Winter 2022. Lab 9
package queues

type Node[T any] struct {
	value T
	next *Node[T]
}

type Queue [T any] struct {
	front  *Node[T]
	rear  *Node[T]
}

func LEmpty[T any]() *Queue[T] {
	q := new (Queue[T])
	q.front = nil
	q.rear = nil
	return q
}

func (q *Queue[T]) IsEmpty() bool {
	return q.front == nil
}

func (q *Queue[T]) Dequeue() T {
	if (q.front == nil){
		var result T
		return result
	} else {
	    n := q.front
		q.front = n.next
		if (q.front == nil){
		   	q.rear = nil
	        }
		v := n.value
		return v
	}
}

func (q *Queue[T]) Enqueue (v T){
	n := new (Node[T])
	n.value = v
	n.next = nil
	if (q.rear == nil){
		q.front = n
		q.rear = n
	} else{
		q.rear.next = n
		q.rear = n
	}
}

