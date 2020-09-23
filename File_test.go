// File_test.go.

//============================================================================//
//
// Copyright © 2018..2020 by McArcher.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
//============================================================================//
//
// Web Site:		'https://github.com/neverwinter-nights'.
// Author:			McArcher.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

package bencode

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/tester"
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
	var aTest *tester.Test = tester.New(t)
	fmt.Println("Creating a Folder:", TestFolder)

	err := os.MkdirAll(TestFolder, os.ModeDir)
	aTest.MustBeNoError(err)
}

func deleteTestFolder(t *testing.T) {
	var aTest *tester.Test = tester.New(t)
	fmt.Println("Deleting a Folder:", TestFolder)

	err := os.RemoveAll(TestFolder)
	aTest.MustBeNoError(err)
}

func createTestFileA(t *testing.T) {
	var aTest *tester.Test = tester.New(t)

	filePath := filepath.Join(TestFolder, TestFileAName)
	fmt.Println("Creating a File:", filePath)
	file, err := os.Create(filePath)
	aTest.MustBeNoError(err)

	_, err = file.WriteString(TestFileAContents)
	aTest.MustBeNoError(err)

	err = file.Close()
	aTest.MustBeNoError(err)
}

func createTestFileB(t *testing.T) {
	var aTest *tester.Test = tester.New(t)

	filePath := filepath.Join(TestFolder, TestFileBName)
	fmt.Println("Creating a File:", filePath)
	file, err := os.Create(filePath)
	aTest.MustBeNoError(err)

	_, err = file.WriteString(TestFileBContents)
	aTest.MustBeNoError(err)

	err = file.Close()
	aTest.MustBeNoError(err)
}

func Test_getContents(t *testing.T) {
	var aTest *tester.Test = tester.New(t)

	// Test Initialization.
	createTestFolder(t)
	createTestFileA(t)
	filePath := filepath.Join(TestFolder, TestFileAName)
	var f *File = NewFile(filePath)

	// Test Finalization.
	defer func() {
		deleteTestFolder(t)
	}()

	// Open the File.
	var err error
	err = f.open()
	aTest.MustBeNoError(err)
	defer func() {
		// Close the File.
		var derr error
		derr = f.close()
		aTest.MustBeNoError(derr)
	}()

	// Test #1. Positive.
	var fileContents []byte
	fileContents, err = f.getContents()

	// Results Inspection.
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileContents, []byte(TestFileAContents))

	// Test #2. Negative.
	fOsFileOriginalValue := f.osFile
	f.osFile = nil
	fileContents, err = f.getContents()
	aTest.MustBeAnError(err)
	f.osFile = fOsFileOriginalValue
}

func Test_Parse(t *testing.T) {
	var aTest *tester.Test = tester.New(t)

	// Test Initialization.
	createTestFolder(t)
	createTestFileB(t)
	filePath := filepath.Join(TestFolder, TestFileBName)
	var f *File = NewFile(filePath)

	// Test Finalization.
	defer func() {
		deleteTestFolder(t)
	}()

	// Open the File.
	var err error
	err = f.open()
	aTest.MustBeNoError(err)
	defer func() {
		// Close the File.
		var derr error
		derr = f.close()
		aTest.MustBeNoError(derr)
	}()

	// Test #1. Positive.
	var do *DecodedObject
	do, err = f.Parse()

	// Results Inspection.
	aTest.MustBeDifferent(do, nil)
	const Sha1SumTextExpected = "beb11253d7cbb2eed50ee54e33218dc1829345a7"
	var doExpected *DecodedObject = &DecodedObject{
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
