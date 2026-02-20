package uaparser

import (
	"fmt"
	"testing"
)

func TestDebugAIX(t *testing.T) {
	ua := "Mozilla/5.0 (X11; U; AIX 000138384C00; en-US; rv:1.0.1) Gecko/20030213 Netscape/7.0"
	os := New(ua).GetOS()
	fmt.Printf("AIX result: name=%q version=%q\n", os.Name, os.Version)
}
