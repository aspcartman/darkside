package darkside

import (
	"fmt"

	"github.com/aspcartman/darkside/g"
	"github.com/alangpierce/go-forceexport"
	"github.com/bouk/monkey"
)

//noinspection SpellCheckingInspection,GoSnakeCaseUsage
func SetUnrecoveredPanicHandler(hanler func(p *g.Panic)) {
	var throw func(string)
	var startpanic func()
	var systemstack func(func())
	var startpanic_m func()
	var preprintpanics func(p *g.Panic)
	forceexport.GetFunc(&throw, "runtime.throw")
	forceexport.GetFunc(&startpanic, "runtime.startpanic")
	forceexport.GetFunc(&systemstack, "runtime.systemstack")
	forceexport.GetFunc(&startpanic_m, "runtime.startpanic_m")
	forceexport.GetFunc(&preprintpanics, "runtime.preprintpanics")

	monkey.Patch(startpanic, func() {
		defer func() {
			if recover() != nil {
				throw("panic while printing panic value")
			}
		}()

		p := g.GetG().Panic
		hanler(p)

		for p != nil {
			switch v := p.Arg.(type) {
			case error:
				p.Arg = v.Error()
			case fmt.Stringer:
				p.Arg = v.String()
			}
			p = p.Link
		}

		systemstack(startpanic_m)
	})
	monkey.Patch(preprintpanics, func(p *g.Panic) {
		// noop
	})
}
