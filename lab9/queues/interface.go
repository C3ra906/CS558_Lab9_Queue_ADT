package queues

type adt [T any] interface {
	IsEmpty() bool
	Enqueue()
	Dequeue() T
}
