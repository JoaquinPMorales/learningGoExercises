The simple calculator program doesn’t handle one error case: division by zero.

Change the function signature for the math operations to return both an int and an error. In the div function, if the divisor is 0, return errors.New("division by zero") for the error. In all other cases, return nil. Adjust the main function to check for this error.

Calculator example: https://github.com/learning-go-book-2e/ch05/blob/main/sample_code/calculator/main.go

