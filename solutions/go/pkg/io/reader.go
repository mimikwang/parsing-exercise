package io

import (
	"io"
	"io/ioutil"
	"path"
)

// Reader iterates through a directory and returns the read bytes
type Reader struct {
	dirPath string
	files   []string
	index   int
}

// NewReader constructs a new Reader given a path to a directory
func NewReader(dirPath string) (Reader, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return Reader{}, err
	}
	filePaths := []string{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePaths = append(filePaths, file.Name())
	}
	return Reader{dirPath: dirPath, files: filePaths, index: 0}, nil
}

// Read reads a file and return bytes
func (r *Reader) Read() ([]byte, error) {
	if r.index == len(r.files) {
		return []byte{}, io.EOF
	}
	data, err := ioutil.ReadFile(path.Join(r.dirPath, r.files[r.index]))
	r.index++
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// FileName returns the file name of the current file
func (r *Reader) FileName() (string, error) {
	if r.index == len(r.files) {
		return "", io.EOF
	}
	return r.files[r.index], nil
}
