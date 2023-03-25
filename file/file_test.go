package file

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func Test_GetFileLines(t *testing.T) {
	// create a txt file with 3 lines a b c

	// create a test to read the file and return the lines
	file, err := ioutil.TempFile("", "topics.txt")
	if err != nil {
		t.Error("Error is = ", err)
	}

	defer os.Remove(file.Name())

	content := []byte("a,b\nc, d")

	if _, err := file.Write(content); err != nil {
		t.Error("Error is = ", err)
	}

	lines, err := GetFileLines(file.Name())
	if err != nil {
		t.Error("Error is = ", err)
	}

	expected := []string{"'a',", "'c'"}

	if !reflect.DeepEqual(lines, expected) {
		t.Error("Expected ", expected, " got ", lines)
	}
}
