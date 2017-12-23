package darkside_test

import (
	"testing"
	"github.com/aspcartman/darkside/g"
	"fmt"
	"github.com/aspcartman/darkside"
)

func TestPanicSafenet(t *testing.T) {
	darkside.SetUnrecoveredPanicHandler(func(p *g.Panic) {
		if p.Arg == "lol" {
			fmt.Println("I've actually succeed")
		} else {
			fmt.Println("I caught things, but arg is wrong")
		}

	})
	panic("lol")
}
