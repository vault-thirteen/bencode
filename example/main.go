package main

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/neverwinter-nights/bencode"
)

func main() {
	var err error = checkFileHashSum()
	if err == nil {
		fmt.Println("OK")
	} else {
		fmt.Println("Error.", err)
	}
}

func checkFileHashSum() (err error) {

	// Settings.
	const (
		ExampleFolder = "example"
		DataFolder    = "data"
		Btih          = "9ddf6a9b17b624991b39f8afd2edc64f673350e3"
		FileName      = "5942384.torrent"
	)

	// Parse the File.
	var f *bencode.File = bencode.NewFile(
		filepath.Join(ExampleFolder, DataFolder, FileName),
	)
	var decodedObject *bencode.DecodedObject
	decodedObject, err = f.Parse()
	if err != nil {
		return
	}

	// Check the BTIH.
	var ok bool = (decodedObject.BTIH.Text == Btih)
	var btihBytesExpected []byte
	btihBytesExpected, err = hex.DecodeString(Btih)
	if err != nil {
		return
	}
	ok = ok && (bytes.Equal(
		decodedObject.BTIH.Bytes[:],
		btihBytesExpected,
	))
	if !ok {
		var msg = fmt.Sprintf(
			"BTIH Mismatch. Expected:%v. Got:%v.",
			Btih,
			decodedObject.BTIH.Text,
		)
		err = errors.New(msg)
		return
	}

	return
}
