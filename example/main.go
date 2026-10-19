package main

import (
	"os"
	"path/filepath"

	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/errors"

	"github.com/vault-thirteen/bencode"
)

// Settings.
const (
	ExampleFolder = "example"
	DataFolder    = "data"
	FileName      = "5942384.torrent"
)

const (
	ExitCode_Failure = 1
)

func main() {
	var err = decodeFile()
	mustBeNoError(err)
}

func mustBeNoError(err error) {
	if err != nil {
		errors.LogError(err)
		os.Exit(ExitCode_Failure)
	}
}

func decodeFile() (err error) {
	var f = bencode.NewFile(
		filepath.Join(ExampleFolder, DataFolder, FileName),
	)

	var decodedObject *bencode.DecodedObject
	decodedObject, err = f.Parse(true)
	if err != nil {
		return err
	}

	cio.PrintMessage(decodedObject.FilePath)

	return nil
}
