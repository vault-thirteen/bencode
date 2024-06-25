package bencode

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

// Test Settings.
const (
	TestFolder        = "test"
	TestFileAName     = "file-a.txt"
	TestFileAContents = "Test Contents."
	TestFileBName     = "file-b.txt"
	TestFileBContents = "d4:info3:Sune"
	TestFileCName     = "file-c.txt"
)

func createTestFolder(t *testing.T) {
	var aTest = tester.New(t)
	fmt.Println("Creating a Folder:", TestFolder)

	err := os.MkdirAll(TestFolder, os.ModeDir)
	aTest.MustBeNoError(err)
}

func deleteTestFolder(t *testing.T) {
	var aTest = tester.New(t)
	fmt.Println("Deleting a Folder:", TestFolder)

	err := os.RemoveAll(TestFolder)
	aTest.MustBeNoError(err)
}

func createTestFileA(t *testing.T) {
	var aTest = tester.New(t)

	filePath := filepath.Join(TestFolder, TestFileAName)
	fmt.Println("Creating a File:", filePath)
	file, err := os.Create(filePath)
	aTest.MustBeNoError(err)

	defer func() {
		derr := file.Close()
		aTest.MustBeNoError(derr)
	}()

	_, err = file.WriteString(TestFileAContents)
	aTest.MustBeNoError(err)
}

func createTestFileB(t *testing.T) {
	var aTest = tester.New(t)

	filePath := filepath.Join(TestFolder, TestFileBName)
	fmt.Println("Creating a File:", filePath)
	file, err := os.Create(filePath)
	aTest.MustBeNoError(err)

	defer func() {
		derr := file.Close()
		aTest.MustBeNoError(derr)
	}()

	_, err = file.WriteString(TestFileBContents)
	aTest.MustBeNoError(err)
}

func Test_NewFile(t *testing.T) {
	var aTest = tester.New(t)

	var file = NewFile("file_path")
	aTest.MustBeEqual(file.path, "file_path")
}

func Test_File_getContents(t *testing.T) {
	var aTest = tester.New(t)

	// Test Initialization.
	createTestFolder(t)
	createTestFileA(t)
	filePath := filepath.Join(TestFolder, TestFileAName)
	var f = NewFile(filePath)

	// Test Finalization.
	defer func() {
		deleteTestFolder(t)
	}()

	var fileContents []byte
	var err error

	// Test #1. Positive.
	{
		// Open the File.
		err = f.open()
		aTest.MustBeNoError(err)

		fileContents, err = f.getContents()
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(fileContents, []byte(TestFileAContents))

		// Close the File.
		err = f.close()
		aTest.MustBeNoError(err)
	}

	// Test #2. Negative.
	{
		fOsFileOriginalValue := f.osFile
		f.osFile = nil
		fileContents, err = f.getContents()
		aTest.MustBeAnError(err)
		f.osFile = fOsFileOriginalValue
	}
}

func Test_File_open(t *testing.T) {
	var aTest = tester.New(t)

	// Test Initialization.
	createTestFolder(t)
	createTestFileA(t)

	// Test Finalization.
	defer func() {
		deleteTestFolder(t)
	}()

	var f *File
	var err error

	// Test #1. Positive.
	{
		filePath := filepath.Join(TestFolder, TestFileAName)
		f = NewFile(filePath)

		// Open the File.
		err = f.open()
		aTest.MustBeNoError(err)

		// Close the File.
		err = f.close()
		aTest.MustBeNoError(err)
	}

	// Test #2. Negative.
	{
		filePath := filepath.Join(TestFolder, TestFileCName)
		f = NewFile(filePath)

		// Open the File.
		err = f.open()
		aTest.MustBeAnError(err)
	}
}

func Test_File_Parse(t *testing.T) {
	var aTest = tester.New(t)

	// Test Initialization.
	createTestFolder(t)
	createTestFileB(t)
	filePath := filepath.Join(TestFolder, TestFileBName)
	var f = NewFile(filePath)

	// Test Finalization.
	defer func() {
		deleteTestFolder(t)
	}()

	var do *DecodedObject
	var doExpected *DecodedObject
	var err error

	// Test #1. Positive without self-check.
	{
		do, err = f.Parse(false)
		aTest.MustBeNoError(err)
		aTest.MustBeDifferent(do, (*DecodedObject)(nil))

		doExpected = &DecodedObject{
			FilePath:   filePath,
			SourceData: []byte(TestFileBContents),
			RawObject: []DictionaryItem{
				{
					Key:      []byte("info"),
					Value:    []byte("Sun"),
					KeyStr:   "info",
					ValueStr: "Sun",
				},
			},
			DecodeTimestamp: do.DecodeTimestamp, // Synchronization with Test.
			//
			IsSelfChecked: false,
		}
		aTest.MustBeEqual(do, doExpected)
	}

	// Test #2. Positive with self-check.
	{
		do, err = f.Parse(true)
		aTest.MustBeNoError(err)
		aTest.MustBeDifferent(do, (*DecodedObject)(nil))

		doExpected = &DecodedObject{
			FilePath:   filePath,
			SourceData: []byte(TestFileBContents),
			RawObject: []DictionaryItem{
				{
					Key:      []byte("info"),
					Value:    []byte("Sun"),
					KeyStr:   "info",
					ValueStr: "Sun",
				},
			},
			DecodeTimestamp: do.DecodeTimestamp, // Synchronization with Test.
			//
			IsSelfChecked: true,
		}
		aTest.MustBeEqual(do, doExpected)
	}
}
