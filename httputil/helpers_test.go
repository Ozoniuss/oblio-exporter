package httputil

import (
	"testing"
)

func Test_GetFileNameFromHeader_CorrectFilename(t *testing.T) {
	filename, _ := GetFileNameFromHeader(`attachment; filename="NB0007_eFactura.xml"`)
	if filename != "NB0007_eFactura.xml" {
		t.Errorf("got invalid filename: %s", filename)
	}
}

func Test_GetFileNameFromHeader_MissingFilenameParameter(t *testing.T) {
	_, err := GetFileNameFromHeader(`attachment; "NB0007_eFactura.xml"`)
	if err == nil {
		t.Error("expected error when filename parameter is missing")
	}
}

func Test_GetFileNameFromHeader_EmptyFilenameParameter(t *testing.T) {
	_, err := GetFileNameFromHeader(`attachment; filename=`)
	if err == nil {
		t.Error("expected error when filename is empty")
	}
}
func Test_GetFileNameFromHeader_MissingFirstQuote(t *testing.T) {
	_, err := GetFileNameFromHeader(`attachment; filename=abc"`)
	if err == nil {
		t.Error("expected error when filename does not start with quote")
	}
}

func Test_GetFileNameFromHeader_MissingLastQuote(t *testing.T) {
	_, err := GetFileNameFromHeader(`attachment; filename="abc`)
	if err == nil {
		t.Error("expected error when filename does not end with quote")
	}
}
