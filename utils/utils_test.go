package utils

import (
	"os"
	"testing"
	"time"
)

// TestExist calls utils.Exist with a path, checking
// for a valid return value.
func TestExistFile(t *testing.T) {
	path := "./temp_" + time.Now().Format("20060102_150405")
	f, err := os.Create(path)
	if err != nil {
		t.Fatalf("Cannot create a temp file!")
	}
	f.Close()

	want := true
	val := Exist(path)
	if want != val {
		t.Fatalf(`Exist("%s") = %t, want match for %t`, path, val, want)
	}

	err = os.Remove(path)
	if err != nil {
		t.Fatalf("Cannot create a temp file!")
	}
}

// TestExist calls utils.Exist with a non-existing path, checking
// for a valid return value.
func TestExistNot(t *testing.T) {
	path := "./temp_" + time.Now().Format("20060102_150405")
	want := false
	val := Exist(path)
	if want != val {
		t.Fatalf(`Exist("%s") = %t, want match for %t`, path, val, want)
	}
}

// Make a directory if the path does not exist
func TestMakeDirIfNotExist(t *testing.T) {
	path := "./temp_" + time.Now().Format("20060102_150405")
	err := MakeDirIfNotExist(path)
	if err != nil {
		t.Fatalf("Cannot create a new directory")
	}
	defer os.Remove(path)

	if !Exist(path) {
		t.Fatalf(`MakeDirIfNotExist("%s"), shall create a new directory, but it did not.`, path)
	}
	if err != nil {
		t.Fatalf(`MakeDirIfNotExist("%s"), want no error, but got:\n%s`, path, err)
	}
}
