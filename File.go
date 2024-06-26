package bencode

import (
	"bufio"
	"errors"
	"io"
	"os"
	"time"

	ae "github.com/vault-thirteen/auxie/errors"
)

// File is a file.
type File struct {
	path   string
	osFile *os.File
}

// NewFile is a file's constructor.
func NewFile(filePath string) (f *File) {
	f = &File{
		path: filePath,
	}

	return f
}

// close closes a file.
func (f *File) close() (err error) {
	return f.osFile.Close()
}

// getContents reads the contents of an opened file.
func (f *File) getContents() (fileContents []byte, err error) {

	// Fool check.
	if f.osFile == nil {
		return nil, errors.New(ErrFileNotInitialized)
	}

	_, err = f.osFile.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	fileContents, err = io.ReadAll(f.osFile)
	if err != nil {
		return nil, err
	}

	return fileContents, nil
}

// open opens a file.
func (f *File) open() (err error) {
	f.osFile, err = os.Open(f.path)
	if err != nil {
		return err
	}

	return nil
}

// Parse parses an input file into an interface. It also stores some
// additional data, all packed into an object.
// If 'makeSelfCheck' flag is enabled, the self check is performed after
// decoding.
func (f *File) Parse(makeSelfCheck bool) (result *DecodedObject, err error) {

	// Open the file and prepare a stream reader.
	err = f.open()
	if err != nil {
		return nil, err
	}

	defer func() {
		// Close the file.
		derr := f.close()
		if derr != nil {
			err = ae.Combine(err, derr)
		}
	}()

	var bufioReader = bufio.NewReader(f.osFile)

	// Parse the file encoded with 'bencode' encoding into an object.
	var decoder = NewDecoder(bufioReader)
	var ifc any
	ifc, err = decoder.readBencodedValue()
	if err != nil {
		return nil, err
	}

	// Get the file contents.
	var fileContents []byte
	fileContents, err = f.getContents()
	if err != nil {
		return nil, err
	}

	// Prepare the result.
	var decodedObject *DecodedObject
	decodedObject = &DecodedObject{
		FilePath:        f.path,
		SourceData:      fileContents,
		RawObject:       ifc,
		DecodeTimestamp: time.Now().Unix(),
	}

	// Perform a self-check if needed.
	if makeSelfCheck {
		ok := decodedObject.MakeSelfCheck()
		if !ok {
			return nil, errors.New(ErrSelfCheck)
		}
	}

	return decodedObject, nil
}

// GetPath returns the path.
func (f *File) GetPath() (path string) {
	return f.path
}
