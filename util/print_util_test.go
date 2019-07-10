package util

import (
	"testing"
)

func TestPrintInTextPlain(t *testing.T) {
	m := make(map[string]string)
	m["test1"] = "test1"
	m["test2"] = "test2"
	m["test3"] = "test3"

	PrintInTextPlain(m)
}

func TestPrintInMarkdownFormat(t *testing.T) {
	m := make(map[string]string)
	m["test1"] = "test1"
	m["test2"] = "test2"
	m["test3"] = "test3"

	PrintInMarkdownFormat(m)
}
