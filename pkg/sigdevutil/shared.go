// Potentially shared functions that can be used for different types of
// signature generation.

package sigdevutil

import (
	"strings"
	"time"
)

// Now is highly changeable. This makes it difficult to test when tests
// are static. To enable mocking, we can alias this variable here, and
// then replace it during the tests.
var now = time.Now

// generateDate returns the current time, i.e. the time the signature
// file is being created.
func generateDate() string {
	const dateFormat = "2006-01-02T15:04:05"
	currentTime := now()
	return currentTime.Format(dateFormat)
}

// generateDateNoSpaces returns the current date with no spaces for
// container signatures.
func generateDateNoSpaces() string {
	const dateFormat = "20060102"
	currentTime := now()
	return currentTime.Format(dateFormat)
}

// formatFilenameString returns a normalized component of a filename to
// the caller.
func formatFilenameString(input string) string {
	niceName := strings.Replace(input, "/", "-", -1)
	niceName = strings.Replace(niceName, " ", "-", -1)
	niceName = strings.Replace(niceName, "*", "-", -1)
	return niceName
}
