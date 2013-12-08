package telnet

import (
	"bytes"
	"testing"
)

func TestUnserializeText(t *testing.T) {
	cmds := make(chan Command)
	go Unserialize(bytes.NewBufferString("Hello World\n"), cmds)
	cmd := <-cmds
	str, ok := cmd.(string)
	if !ok || str != "Hello World" {
		t.Error("Didn't parse text!")
	}
}

func TestCmdBreak(t *testing.T) {
	cmds := make(chan Command)
	go Unserialize(bytes.NewBuffer([]byte{IAC, BRK}), cmds)
	cmd := <-cmds
	_, ok := cmd.(Break)
	if !ok {
		t.Error("Didn't parse break!")
	}
}

func TestCmdEraseChar(t *testing.T) {
	cmds := make(chan Command)
	buf := bytes.NewBufferString("")
	buf.Write([]byte{IAC, EC})
	buf.WriteString("hellow")
	buf.Write([]byte{IAC, EC})
	buf.WriteString(" world\n")

	go Unserialize(buf, cmds)
	cmd := <-cmds
	str, ok := cmd.(string)
	if !ok || str != "hello world" {
		t.Error("Didn't handle erase char correctly!")
	}
}

func TestCmdEraseLine(t *testing.T) {
	cmds := make(chan Command)
	buf := bytes.NewBufferString("hellow")
	buf.Write([]byte{IAC, EL})
	buf.WriteString("world\n")

	go Unserialize(buf, cmds)
	cmd := <-cmds
	str, ok := cmd.(string)
	if !ok || str != "world" {
		t.Error("Didn't handle erase line correctly!")
	}
}
