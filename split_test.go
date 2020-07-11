package filesplit

import (
	"io/ioutil"
	"testing"
)

func TestSplit(t *testing.T) {
	// create test file
	file := "/tmp/file.dat"
	size := int64(2097152) // 2MB
	if err := CreateTestFile(file, size); err != nil {
		t.Errorf("Not able to create \"%s\" file", file)
	}

	// and chunk it!
	chunks, err := Split(file)
	if err != nil {
		t.Errorf("Not able to split file")
	}

	expected := 4
	if len(chunks) != expected {
		t.Errorf("Chunks number is %v and should be %v", len(chunks), expected)
	}
}

func TestSplitFromBytes(t *testing.T) {
	// create test file
	file := "/tmp/file.dat"
	size := int64(2097152) // 2MB
	if err := CreateTestFile(file, size); err != nil {
		t.Errorf("Not able to create \"%s\" file", file)
	}
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	// and chunk it!
	chunks, err := SplitFromBytes(fileBytes)
	if err != nil {
		t.Errorf("Not able to split file")
	}

	expected := 4
	if len(chunks) != expected {
		t.Errorf("Chunks number is %v and should be %v", len(chunks), expected)
	}
}

func TestSave(t *testing.T) {
	// create test file
	file := "/tmp/file.dat"
	size := int64(2097152) // 2MB
	if err := CreateTestFile(file, size); err != nil {
		t.Errorf("Not able to create \"%s\" file", file)
	}

	// chunk it
	chunks, err := Split(file)
	if err != nil {
		t.Errorf("Not able to split file")
	}

	// and save it!
	if err := Save(chunks, "/tmp/"); err != nil {
		t.Error(err)
	}
}
