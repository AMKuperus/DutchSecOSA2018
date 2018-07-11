package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AMKuperus/DSOSA/osa2018/headers"
	"github.com/fatih/color"
)

// Declare a function to run the program
func main() {
	Run()
}

// Run reads in a .txt-file and performs a action on every line it reads
func Run() (string, error) {
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
	var c int
	// With c as counter we loop over all lines existing in the file
	for scanner.Scan() {
		// Make sure we add 1 to c, otherwise we would end up getting only the first
		// result over and over in a neverending loop.
		c++
		// Store the line we read as text in host
		host := scanner.Text()
		// Call function Check() with string host (which is a url)
		result := Check(host)
		// with fmt we print the result to the terminal-window
		fmt.Printf("[%-5d] %s\n", c, result)
	}
	// When we are at the end of the loop we end up here and print --Finished--
	// to the terminal window.
	return fmt.Sprintf("--%s--\n", "Finished"), nil
}

// Check creates struct and returns info gathered by calls made to the struct.
func Check(host string) string {
	// Declare xframe as a headers.Xrame-struct
	var header headers.Header
	// Run New to set and perform checks.
	header.New(host)
	// Store return string and boolean from Xframe.IsCorrect() in mistakes and iscorrect
	// Note that normally we use shorter names. The long names used in this code are
	// for better understanding but are not good practice.
	// Normally we would call these 2 variables something like:  str, boo
	mistakes, iscorrect := header.Xframe.IsCorrect()

	recordoutput := fmt.Sprintf("URL [%-25s] IsSet [%-5t] IsCorrect [%-5t] Mistakes: %s",
		host, header.Xframe.IsSet(), iscorrect, mistakes)

	Record("output.txt", recordoutput)

	ret := color.CyanString(host) + "\n" + color.YellowString(header.ShowXframe())
	return ret
}

// Record writes input to a file
func Record(filedirectory, input string) {
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
}
