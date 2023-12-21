package bencode

import (
	"encoding/hex"
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

func Test_getContents(t *testing.T) {
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

	// Open the File.
	var err error
	err = f.open()
	aTest.MustBeNoError(err)

	// Test #1. Positive.
	var fileContents []byte
	fileContents, err = f.getContents()

	// Results Inspection.
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileContents, []byte(TestFileAContents))

	// Close the File.
	err = f.close()
	aTest.MustBeNoError(err)

	// Test #2. Negative.
	fOsFileOriginalValue := f.osFile
	f.osFile = nil
	fileContents, err = f.getContents()
	aTest.MustBeAnError(err)
	f.osFile = fOsFileOriginalValue
}

func Test_Parse(t *testing.T) {
	var aTest = tester.New(t)
	var err error

	// Test Initialization.
	createTestFolder(t)
	createTestFileB(t)
	filePath := filepath.Join(TestFolder, TestFileBName)
	var f = NewFile(filePath)

	// Test Finalization.
	defer func() {
		deleteTestFolder(t)
	}()

	// Test #1. Positive.
	var do *DecodedObject
	do, err = f.Parse()

	// Results Inspection.
	aTest.MustBeDifferent(do, nil)
	const Sha1SumTextExpected = "beb11253d7cbb2eed50ee54e33218dc1829345a7"
	var doExpected = &DecodedObject{
		FilePath:   filePath,
		SourceData: []byte(TestFileBContents),
		DecodedObject: []DictionaryItem{
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
		BTIH: BtihData{
			Bytes: Sha1Sum{}, // See below.
			Text:  Sha1SumTextExpected,
		},
	}
	var ba []byte
	ba, err = hex.DecodeString(Sha1SumTextExpected)
	aTest.MustBeNoError(err)
	copy(doExpected.BTIH.Bytes[:], ba)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(do, doExpected)
}
