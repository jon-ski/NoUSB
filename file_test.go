package main

import (
	"testing"
)

func TestGetParentFolder(t *testing.T) {
	if getParentFolder() != "nousb\\" {
		t.Error("Get Parent Folder did not return \"nousb\\\"")
	}
}
