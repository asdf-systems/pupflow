package golo

import (
	"testing"
	// "net"
	"time"
)

func TestServer(t *testing.T) {
	c := make(chan *Message)
	StartServer(7770, c)
	go func() {
		for {
			c <- nil
			time.Sleep(1 * time.Second)
		}
	}()
	for m := range c {
		if m == nil {
			continue
		}
		t.Logf("%s:\n", m.Path)
		for _, p := range m.Params {
			t.Logf("\t%f\n", p.(float32))
		}
	}
}
