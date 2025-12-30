Create a struct named Person with three fields: FirstName and LastName of type string and Age of type int. Write a function called MakePerson that
takes in firstName, lastName, and age and returns a Person. Write a second function MakePersonPointer that takes in firstName, lastName, and age and returns a *Person. Call both from main. Compile your program with go build
-gcflags="-m". This both compiles your code and prints out which values escape to the heap. Are you surprised about what escapes?

Solution:

The &Person{} returned from MakePersonPointer escapes to the heap. Any time a pointer is returned from a function, the pointer is returned on the stack, but the value it points to must be stored on the heap.

Surprisingly, it also says that p escapes to the heap. The reason for this is that it is passed into fmt.Println. This is because the parameter to fmt.Println are ...any. The current Go compiler moves to the heap any value that is passed in to a function via a parameter that is of an interface type. (Interfaces are covered in Chapter 7.)