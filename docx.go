package main

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"strings"
)

type Docx struct {
	files   []*zip.File
	content string
}

func (d *Docx) AppendText(text string) {
	d.content += text
}

func (d *Docx) WriteToFile(path string) (err error) {
	var target *os.File
	target, err = os.Create(path)
	if err != nil {
		return
	}
	defer target.Close()
	err = d.Write(target)
	return
}

func (d *Docx) Write(ioWriter io.Writer) (err error) {
	w := zip.NewWriter(ioWriter)
	for _, file := range d.files {
		var writer io.Writer
		var readCloser io.ReadCloser

		writer, err = w.Create(file.Name)
		if err != nil {
			return err
		}
		readCloser, err = file.Open()
		if err != nil {
			return err
		}
		if file.Name == "word/document.xml" {
			writer.Write([]byte(d.content))
		} else {
			writer.Write(streamToByte(readCloser))
		}
	}
	w.Close()
	return
}

func streamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

func ReadDocxFromFS(file string) (*Docx, error) {
	zipReader, err := zip.OpenReader(file)
	if err != nil {
		return nil, err
	}

	content := ""
	for _, f := range zipReader.File {
		if f.Name == "word/document.xml" {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			buf := new(bytes.Buffer)
			buf.ReadFrom(rc)
			content = buf.String()
			break
		}
	}

	docx := &Docx{
		files:   zipReader.File,
		content: content,
	}

	return docx, nil
}
