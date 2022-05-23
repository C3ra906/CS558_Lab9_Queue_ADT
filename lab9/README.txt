Cera Oh. CS558. Winter 2022. Lab 9
I chose to code in Go

a. How to Compile:
The zipped folder contains this README file, go.mod file, main.go (the driver file), and a folder called "queues" that serves as a package. 
The "queues" package contains LLqueues.go file, which is the Linearly Linked List Implementation of the queue, 
Rqueues.go file, which is the Ring Implementation of the queue, and interface.go file that serves as the interface. 
I'm new to Go, but I believe the go packages are compiled together as one, so the two implementations of the queues are compiled together
with the interface. I have the queues package imported into main.go and ready for use, and nothing needs to be done to run the ADT code 
but to run the main.go. The driver code as written excercises both types of queues using int and string types using the package. 

I used 1.18beta2 version of Go, which is an unstable version of Go. However, it allows use of generic types, which allowed me to write 
functions that take generic types instead of writing seperate overloaded functions for int and string types. You can install 1.18beta2 
by entering into your OS terminal:
go install golang.org/dl/go1.18beta2@latest
go1.18beta2 download

Once you have the 1.18beta2 of Go installed you can go the directory where you unzipped the folder and enter:
go1.18beta2 run main.go 
to run the program. 

b. How to modify driver code:
There is no need to modify the driver code. 

c. Possible Issues, Failures, and Comments:
I had some hard time implenting the interface. I was able to figure out how to make methods, which are functions with same name that 
operates on different types, work when the methods returned the same type of data but not when it was not returning the same type of data.
Empty() returns the queue, which could be a pointer to either the struct type RQueue for ring implementation or the struct type Queue for
linked list implementation. Since I could not declare a struct pointer type that matched for both types of data, I was unable to make a 
method that worked. I was able to include in the interface methods isEmpty(), Dequeue(), and Enqueue() because they return same type for 
both instances of the methods.

Structs are mutable in Go.

I used generic type T for values held by the queues so I can write functions using generic type instead of having multiple functions to
handle different types. While return type of Dequeue() should be elem + None, I either returned the element or nothing by declaring a
variable of the type without assigning a value. Go would not take nil for generic type T for returning or assignment. Both versions of 
Dequeue therefore returns Type T, either empty or with value. If Go version 1.18beta2 was not avaiable, this assignment would have been
very difficult since there would be no generic types. I'm glad it is getting added to 1.18.

For Ring Implementation, I declared a consant "const N int = 5" in the Rqueues.go file so the program knows the size of the array in the 
ring queue. I tried passing in an int but that did not help with function and method polymorphism so thought declaring a constant at the 
top of the file was the best way to go. I also could not fill a just created empty queue with arbitary value because the type was generic.
I filled it with empty value by assigning a variable of type T without a value. Otherwise I followed the pseudo-code in the assignment.

For Linked List Implementation, I also followed the pseudo-code in the assignment. Pointers could be declared nil, so it was pretty
straight forward to follow the pseudo-code. Go has garbage collection so I did not worry myself with memory allocation and deallocations. 

I have interface in its own file and the two implementations of queues and its functions in their respective files. The implementations 
and interface is parametrically polymorphic. I tried printing value of the queue directly in main.go and it didn't work for either 
implentation, so I think the queue is properly hidden from client code. There is no impact on client code to switch between the two
queues. Since we are importing packages, I do not think you can compile main.go seperately from queues packages from the command line.
However, technically the package and main.go are seperate and codes within packages are always compiled together. 