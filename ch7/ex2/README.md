Write a function called fileLen that has an input parameter of type string and returns an int and an error. The function takes in a filename and returns the number of bytes in the file. If there is an error reading the file, return the error.

Use defer to make sure the file is closed properly.