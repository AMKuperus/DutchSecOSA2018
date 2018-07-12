// Package headers Xframe check. Checks if set, if set correct,
// and can return mistakes.
// This version is only created as part of the OpenSourceAcademy Summerschool by
// DutchSec. When code is reused DutchSec or author cannot be hold accountable for
// any mistakes or harm done by inproper reuse of this code.
// Only for learning purpose.
// @author AMKuperus @company DutchSec
package headers

import "strings"

// Xframe holds data about the X-frame header.
type Xframe struct {
	Config     string
	isset      bool
	deny       bool
	sameorigin bool
	allowall   bool
	allowfrom  bool
}

// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options
// https://www.owasp.org/index.php/Clickjacking_Defense_Cheat_Sheet#Defending_with_X-Frame-Options_Response_Headers/#Defending_with_X-Frame-Options_reponse_Headers

// set checks if there is input and if so sends it trough Xframe-methods for
// validation and filling Xframe struct.
func (x *Xframe) set(input string) Xframe {
	// Check if there is input, with no input the Xframe if not set.
	if len(input) > 0 {
		// Store input to Config.
		x.Config = input
		// set isset to true, as the Xframe header is set.
		x.isset = true
		// call sort() which contains the algorithm to filter everything and store
		// it in the right place.
		x.sort() // -> Follow the code to line 42
	}
	// Return the struct
	return *x
}

// sort is the sorting-algorithm for Xframe.
func (x *Xframe) sort() {
	// First make all characters uppercase. For the header-settings it does not
	// matter if it is upper or lower-case. Since developers all use different
	// writestyle easiest way to check is to set everything to either upper or
	// lower case.
	str := strings.ToUpper(x.Config)

	// A switch-statement on str, that we just created. We compare str to exact
	// matches of good Xframe settings
	switch str {
	case "DENY":
		// When we find a good match, like str == "DENY" we tell Xframe that deny is
		// true, which in this code means it is set.
		x.deny = true
	case "SAMEORIGIN":
		x.sameorigin = true
	case "ALLOWALL":
		// https://ipsec.pl/web-application-security/2013/mess-x-frame-options-allowall.html
		x.allowall = true
	default:
		// Since "ALLOW_FROM" is followed by a url we filter this one out by looking
		// if it contains "ALOW_FROM". Now it can contain more then taht, as long as
		// it contains "ALLOW_FROM"
		if strings.Contains(str, "ALLOW-FROM") {
			x.allowfrom = true
		}
	} // -> Follow  the code back to line 38
}

// IsSet returns if Xframe is set.
func (x Xframe) IsSet() bool {
	return x.isset // -> Following the code back to headers/headers.go line 64
}

// IsCorrect returns string wich is either empty or filled with wrong Xframe setup
// and bool telling setup was OK = true or NOT OK = false
func (x Xframe) IsCorrect() (string, bool) {
	if x.isset {
		// Find a x.var that is true, if found setup is OK
		// The \\ in this function means a OR which will turn true if atleas one is true.
		if x.deny || x.sameorigin || x.allowall || x.allowfrom {
			// OK
			return "", true
		}
		// wrong setup Xframe
		return x.Config, false
	}
	// Xframe not set
	return "", false
} // -> Following the code back to headers/headers.go line 62
