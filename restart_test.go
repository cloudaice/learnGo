package useage

import (
	"testing"
)

func TestFoo(t *testing.T) {
	channel := make(chan string)
	go Foo(channel)
	var info string

	for i := 0; i < 10; i++ {
		info = <-channel
		if info != "Foo failed" {
			t.Error("test restart:Foo failed")
		}
	}
	t.Log("test restart:Foo pass")
}

func TestBar(t *testing.T) {
	channel := make(chan string)
	go Bar(channel)
	var info string
	for i := 0; i < 10; i++ {
		info = <-channel
		if info != "Bar failed" {
			t.Error("test restart:Bar failed")
		}
		go Bar(channel)
	}
	t.Log("test restart:Bar pass")
}
