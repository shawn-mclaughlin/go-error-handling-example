package main

import (
	goerrors "errors"
	"fmt"

	"error-example/errors"
)

func main() {
	basic := errors.NewNotFoundError("couldn't find foo in place")
	someThinlyWrappedError := errors.WrapAsNotFoundError(goerrors.New("some database error saying not found"))
	someDeeplyWrappedError := errors.WrapAsNotFoundError(
		fmt.Errorf("outermost context: %w",
			fmt.Errorf("middle context: %w",
				fmt.Errorf("inner context: %w",
					goerrors.New("some database error saying not found"),
				),
			),
		),
	)
	someOtherError := goerrors.New("some other error")

	fmt.Printf("\"basic\" is not found error: %t\n", errors.IsNotFound(basic))
	fmt.Println(basic.Error())
	fmt.Println()

	fmt.Printf("\"thinlyWrapped\" is not found error: %t\n", errors.IsNotFound(someThinlyWrappedError))
	fmt.Println(someThinlyWrappedError.Error())
	fmt.Println()

	fmt.Printf("\"deeplyWrapped\" is not found error: %t\n", errors.IsNotFound(someDeeplyWrappedError))
	fmt.Println(someDeeplyWrappedError.Error())
	fmt.Println()

	fmt.Printf("\"otherKindOfError\" is not found error: %t\n", errors.IsNotFound(someOtherError))
	fmt.Println(someOtherError.Error())
	fmt.Println()
}
