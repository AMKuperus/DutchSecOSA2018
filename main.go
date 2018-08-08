// This version is only created as part of the OpenSourceAcademy Summerschool by
// DutchSec. When code is reused DutchSec or author cannot be hold accountable for
// any mistakes or harm done by inproper reuse of this code.
// Only for learning purpose.
// @author AMKuperus @company DutchSec
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AMKuperus/DutchSecOSA2018/headers"
	"github.com/fatih/color"
)

// Declare a function to run the program.
func main() {
	Run()
}

// Run reads in a .txt-file and performs a action on every line it reads
func Run() (string, error) { // This function has a string and error to return.
	// Open the file with the url's
	file, err := os.Open("alexa500.txt")
	// When a error occurs we catch it so we keep control of the program.
	// Declaring like this is common in go. First you store the value you want,
	// in this case the file in variable file, second we store the error in err.
	// Note that this function we build here works the same with (string, error)
	if err != nil {
		return "", err
	}
	// When there was no error we continue and scan the file to be able to read
	// the content line by line. First make a scanner "object" so we can use bufio
	// function's.
	scanner := bufio.NewScanner(file)
	// Declare a variable integer to function as a counter for coming loop
	var i int
	// With c as counter we loop over all lines existing in the file
	for scanner.Scan() {
		// Store the line we read as text in variable host
		host := scanner.Text()
		// Call function Check() with string host (which is a url)
		// -> To follow the code go to line 57 to proceed.
		result := Check(host)
		// with fmt we print the result to the terminal-window
		// Computers count from 0, people from 1, so for showing we do i + 1
		fmt.Printf("[%-5d] %s\n", (i + 1), result)
		// Make sure we add 1 to i, otherwise we would end up getting only the first
		// result over and over in a neverending loop.
		i++
	}
	// When we are at the end of the loop we end up here and print --Finished--
	// to the terminal window.
	return fmt.Sprintf("--%s--\n", "Finished"), nil
}

// Check creates struct and returns info gathered by calls made to the struct.
func Check(host string) string { // This function has a string to return.
	// Declare xframe as a headers.Xrame-struct
	var header headers.Header //-> Follow the code to file headers/headers.go line 17
	// Run New to set and perform checks.
	header.New(host) //-> Follow the code to headers/headers.go line 21
	// Store return string and boolean from Xframe.IsCorrect() in mistakes and iscorrect
	// Note that normally we use shorter names. The long names used in this code are
	// for better understanding but are not good practice.
	// Normally we would call these 2 variables something like:  str, boo
	mistakes, iscorrect := header.Xframe.IsCorrect()
	// Create a string with Sprintf, with this method we have a lot of control on
	// strings format.
	recordoutput := fmt.Sprintf("URL [%-25s] IsSet [%-5t] IsCorrect [%-5t] Mistakes: %s",
		host, header.Xframe.IsSet(), iscorrect, mistakes)

	// Use Record() to write the output to the file output.txt
	// -> To follow the code proceed to line 84
	Record("output.txt", recordoutput)

	// Build a return-string. -> Follow the code to headers/header.go line 58
	ret := color.CyanString(host) + "\n" + color.YellowString(header.ShowXframe())
	// Return the ret-string.
	// -> To follow the code return to line 49
	return ret
}

// Record writes input to a file
func Record(filedirectory, input string) { // This function has not return
	// Open a file to write to
	file, err := os.OpenFile(filedirectory, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		// If we can not open the file abort immediately with panic
		panic(err)
	}
	// In Go if we open something like a connection or in this case a file we can
	// directly after opening and catching all errors already tell the program to
	// close it when it is no longer used and needed by using defer xxx.Close()
	defer file.Close()

	// Write the string input to the file and end with a newline. If an error occurs
	// stop the program with a panic as we are unable to record.
	if _, err = file.WriteString(input + "\n"); err != nil {
		panic(err)
	}
	// -> To follow the code proceed to line 77
}
