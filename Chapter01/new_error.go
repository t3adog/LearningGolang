package main

import (
	"errors"
	"fmt"
)

var ERROR_MESSAGE = "Error in returnError() function!"

func return_error(a, b int) error {
	if a == b {
		err := errors.New(ERROR_MESSAGE)
		return err
	} else {
		return nil
	}
}

func main() {
	err := return_error(1, 2)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	err = return_error(10, 10)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	if err.Error() == ERROR_MESSAGE {
		fmt.Println("!!!")
	}
}
