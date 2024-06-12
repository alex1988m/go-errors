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
	listing1()
	listing2()
	listing3()
	listing4()
}

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
