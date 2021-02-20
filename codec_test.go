package jsonrpc2

import (
	"bufio"
	"strings"
	"testing"
)

func TestVSCodeObjectCodec_ReadObject(t *testing.T) {
	s := "Content-Type: foo\r\nContent-Length: 123\r\n\r\n789"
	var v int
	if err := (VSCodeObjectCodec{}).ReadObject(bufio.NewReader(strings.NewReader(s)), &v); err != nil {
		t.Fatal(err)
	}
	if want := 789; v != want {
		t.Errorf("got %v, want %v", v, want)
	}
}
