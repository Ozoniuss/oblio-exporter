package httputil

import (
	"errors"
	"fmt"
	"strings"
)

// GetFileNameFromHeader extracts the filename from a header such as the
// following:
//
// Content-Disposition:[attachment; filename="NB0007_eFactura.xml"]
//
// The value in between the brackets is passed to this function.
func GetFileNameFromHeader(header string) (string, error) {
	start := strings.Index(header, "filename=") + len("filename=")
	if start == -1 {
		return "", errors.New("filename not found in Content-Disposition")
	}
	remainder := header[start:]
	if len(remainder) <= 2 {
		return "", fmt.Errorf("malformed filename in header: %s", header)
	}
	if remainder[0] != '"' || remainder[len(remainder)-1] != '"' {
		return "", fmt.Errorf("malformed filename in header (missing quotes): %s", header)
	}
	return remainder[1 : len(remainder)-1], nil
}
