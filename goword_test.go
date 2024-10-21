package goword

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestOpenWordFileInvalidFile(t *testing.T) {

	_, err := openWordFile("testfiles/text.txt")
	if err == nil {
		t.Errorf("text files are not zip, should fail to open.")
	}
	if !strings.Contains(fmt.Sprintf("%s", err), "not a valid zip file") {
		t.Errorf("error message should be not a valid zip file")
	}

}

func TestOpenWordFileValidFile(t *testing.T) {

	doc, err := openWordFile("testfiles/test.docx")
	if err != nil {
		t.Errorf("failed to open a word file.")
	}
	if !strings.Contains(doc, "This is a word file") {
		t.Errorf("Error reading document.xml %s ", doc)
	}
	fmt.Printf("%s", err)
}

func TestParseText(t *testing.T) {

	_, err := ParseText("testfiles/text.txt")
	if err == nil {
		t.Errorf("parse should fail \n %s", err)
	}

	doctext, err := ParseText("testfiles/test.docx")
	if err != nil {
		t.Errorf("parsing test.docx should work \n %s", err)
	}

	if !strings.Contains(doctext, "This is a word file") {
		t.Errorf("parsed text does not contain expected text \n %s", doctext)
	}
	if !strings.Contains(doctext, "What a lovely doc.") {
		t.Errorf("parsed text does not contain expected text \n %s", doctext)
	}

	doctext2, err := ParseText("testfiles/test2.docx")
	if err != nil {
		t.Errorf("parsing test.docx should work \n %s", err)
	}

	if !strings.Contains(doctext2, "before the table") {
		t.Errorf("parsed text does not contain expected text \n %s", doctext)
	}
	if !strings.Contains(doctext2, "This is text after the table.") {
		t.Errorf("parsed text does not contain expected text \n %s", doctext)
	}

	//fmt.Printf(doc)
}

func TestParseTextFromReader(t *testing.T) {

	dat, err := os.ReadFile("testfiles/text.txt")
	if err != nil {
		t.Errorf("read file error %s", err)
	}

	_, err = ParseTextFromReader(bytes.NewReader(dat), int64(len(dat)))
	if err == nil {
		t.Errorf("parse should fail \n %s", err)
	}

	dat, err = os.ReadFile("testfiles/test.docx")
	if err != nil {
		t.Errorf("read file error %s", err)
	}
	doctext, err := ParseTextFromReader(bytes.NewReader(dat), int64(len(dat)))
	if err != nil {
		t.Errorf("parsing test.docx should work \n %s", err)
	}

	if !strings.Contains(doctext, "This is a word file") {
		t.Errorf("parsed text does not contain expected text \n %s", doctext)
	}
	if !strings.Contains(doctext, "What a lovely doc.") {
		t.Errorf("parsed text does not contain expected text \n %s", doctext)
	}

	dat, err = os.ReadFile("testfiles/test2.docx")
	if err != nil {
		t.Errorf("read file error %s", err)
	}
	doctext2, err := ParseTextFromReader(bytes.NewReader(dat), int64(len(dat)))
	if err != nil {
		t.Errorf("parsing test.docx should work \n %s", err)
	}

	if !strings.Contains(doctext2, "before the table") {
		t.Errorf("parsed text does not contain expected text \n %s", doctext)
	}
	if !strings.Contains(doctext2, "This is text after the table.") {
		t.Errorf("parsed text does not contain expected text \n %s", doctext)
	}

	//fmt.Printf(doc)
}
