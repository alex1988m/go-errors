package main

import (
	"fmt"
	"log"
	"os"
)

// | Error matching?  |	Error Message |	Guidance 											|
// | No 	 						|			static		| errors.New										|
// | No  					 		|			dynamic		| fmt.Errorf										|
// | Yes 							|			static		| top-level var with errors.New	|
// | Yes 							|			dynamic		| custom error type							|

func main() {
	// unwrapped error
	listing1()
	listing2()
	listing3()
	listing4()
	// wrapped error
	listing5()
}

// listing1 demonstrates error handling with a static error message.
//
// It calls the GetUser function with the name "Adam" and false values for the isStaticError and isMatchRequired parameters.
// The GetUser function returns an error with a static error message.
func listing1() {
	// | Error matching?  |	Error Message |	Guidance 											|
	// | No 	 						|			static		| errors.New										|

	// listing1: error is not matched as ErrNotFound or UserNotFoundError
	// listing1: get user: user not found

	logger := log.New(os.Stdout, "listing1: ", 0)
	_, err := GetUser("Adam", true, false)
	logError(err, logger)
	fmt.Println()
}

// listing2 demonstrates error handling with dynamic error messages.
//
// It calls the GetUser function with the name "Adam" and false values for the isStaticError and isMatchRequired parameters.
// The GetUser function returns an error with a dynamic error message.
func listing2() {
	// | Error matching?  |	Error Message |	Guidance 											|
	// | No  					 		|			dynamic		| fmt.Errorf										|

	// listing2: error is not matched as ErrNotFound or UserNotFoundError
	// listing2: get user: user with name "Adam" not found

	logger := log.New(os.Stdout, "listing2: ", 0)
	_, err := GetUser("Adam", false, false)
	logError(err, logger)
	fmt.Println()
}

// listing3 demonstrates error matching with a static error message using a top-level variable.
func listing3() {
	// | Error matching?  |	Error Message |	Guidance 											|
	// | Yes 							|			static		| top-level var with errors.New	|

	// listing3: error is matched as ErrNotFound
	// listing3: get user: user not found

	logger := log.New(os.Stdout, "listing3: ", 0)
	_, err := GetUser("Adam", true, true)
	logError(err, logger)
	fmt.Println()
}

// listing4 demonstrates error matching using a custom error type.
func listing4() {
	// | Error matching?  |	Error Message |	Guidance 											|
	// | Yes 							|			dynamic		| custom error type							|

	// listing4: error is matched as UserNotFoundError
	// listing4: get user: user with name "Adam" not found

	logger := log.New(os.Stdout, "listing4: ", 0)
	_, err := GetUser("Adam", false, true)
	logError(err, logger)
	fmt.Println()
}

// listing5 demonstrates handling wrapped errors and the propagation of errors up the call stack.
//
// This function calls the SaveUser function with different combinations of parameters to simulate various error scenarios.
// It then logs the errors using the logError function and prints a new line.
func listing5() {
	// wrapped error
	logger := log.New(os.Stdout, "listing5: ", 0)
	err := SaveUser("Adam", true, false)
	// listing5: error is not matched as ErrNotFound or UserNotFoundError
	// listing5: unexpected error: save user "Adam": user not found
	logError(err, logger)
	fmt.Println()

	err = SaveUser("Adam", false, false)
	// listing5: error is not matched as ErrNotFound or UserNotFoundError
	// listing5: unexpected error: save user "Adam": user with name "Adam" not found
	logError(err, logger)
	fmt.Println()

	err = SaveUser("Adam", true, true)
	// listing5: error is matched as ErrNotFound
	// listing5: unexpected error: save user "Adam": user not found
	logError(err, logger)
	fmt.Println()

	err = SaveUser("Adam", false, true)
	// listing5: error is matched as UserNotFoundError
	// listing5: unexpected error: save user "Adam": user with name "Adam" not found
	logError(err, logger)
	fmt.Println()
}
