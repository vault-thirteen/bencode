package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/vault-thirteen/bencode"
)

// Settings.
const (
	ExampleFolder = "example"
	DataFolder    = "data"
	FileName      = "5942384.torrent"
)

func main() {
	var err = decodeFile()
	mustBeNoError(err)
}

func mustBeNoError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
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

	fmt.Println(decodedObject.FilePath)

	return nil
}
