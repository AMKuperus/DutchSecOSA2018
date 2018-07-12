// Package headers can check Xframe header. This is a small part of the original
// headers-package by @AMKuperus
// This version is only created as part of the OpenSourceAcademy Summerschool by
// DutchSec. When code is reused DutchSec or author cannot be hold accountable for
// any mistakes or harm done by inproper reuse of this code.
// Only for learning purpose.
// @author AMKuperus @company DutchSec
package headers

import (
	"fmt"
	"net/http"
	"time"
)

// Header holds headers
type Header struct { // -> Follow the code back to main.go line 61
	Xframe Xframe
}

// New makes a get-request, sends the info to the Xframe-struct which fills it
// with the information received by it's algorithm.
// Returns the header with Xframe-struct filled with data.
func (h *Header) New(host string) (Header, error) {
	// Create a url on which we can perform a http-request
	url := fmt.Sprintf("https://%s", host)

	// Set e timeout of 5 seconds, otherwise if the response is to slow we will
	// never get the response.
	timeout := time.Duration(5 * time.Second)
	// Setup http client for performing the get-request.
	client := http.Client{
		Timeout: timeout,
	}

	// Make the request and store the returned data in resp.
	resp, err := client.Get(url)
	// Handle the error
	if err != nil {
		// return  the header-object (empty) and a error containing the host and the
		// error that hopefully gives some usefull information about the error.
		return *h, fmt.Errorf("%s - %s", host, err.Error())
	}
	// Already tell program to close when no longer in use.
	defer resp.Body.Close()

	// From the response get the X-Frame-Options header
	// get := resp.Header.Get("X-Frame-Options")
	// Ask the Xframe object to perform function set() which performs it's algorithm
	// and fills the Xframe struct
	h.Xframe = h.Xframe.set(resp.Header.Get("X-Frame-Options")) // -> Follow the code
	// to headers/xframe.go line 25

	// Now return the filled header-structure with all it's new changes.
	return *h, nil
}

// ShowXframe shows all gathered data from Xframe.
func (h Header) ShowXframe() string {
	// Store IsCorrect() string and boolean.
	str, boo := h.Xframe.IsCorrect() // -> Follow the code to headers/xframe.go line 77
	// Build up a string to return. -> Follow the code to headers/xframe.go line 72
	ret := fmt.Sprintf("X-Frame isset [ %-5t ] IsCorrect [ %-5t ] %s", h.Xframe.IsSet(), boo, str)
	// Return the string we just build.
	return ret //-> Follow the code back to main.go line 79
}
