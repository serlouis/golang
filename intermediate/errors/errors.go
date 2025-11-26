package main

import (
	"errors"
	"fmt"
)

func main() {

	/*result, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	result1, err1 := sqrt(-16)
	if err1 != nil {
		fmt.Println(err1)
		// if we return here program execution stops doh took me forever to track down why
		//return
	}
	fmt.Println(result1)
	*/
	/*data := []byte{}
	//if err := process(data);err != nil { // I hate this formatting the if is evaluating the err easier to read this other way
	err := process(data)
	if err != nil {

		fmt.Println("Error:", err)
		return
	}
	*/
	// this should never print since we return since we're trying to get an error
	//fmt.Println("Data Processed Successfully" )

	/*err1 := eprocess()
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	*/
	err := readData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data read successfully")
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math error: square root of negative number")
	}

	// compute the square root not going to calculate it just as an example
	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Empty data")
	}

	// Process data
	return nil
}

type myError struct {
	message string
}

// go has a built in interfaced called Error interface. It has a single method which is Error() and in go an error is representated by the error interface
// this interface returns a string that describes the error.
// using this interface we can implement our own custom error messages
func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func eprocess() error {
	return &myError{"Custom error message"}
}

func readConfig() error {
	return errors.New("config error")
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}

	return nil
}
