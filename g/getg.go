// Copyright 2016 Huan Du. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package g

import "unsafe"

func getg() *G

func GetG() *G {
	return getg()
}

type Panic struct {
	Argp      unsafe.Pointer // pointer to arguments of deferred call run during panic; cannot move - known to liblink
	Arg       interface{}    // argument to panic
	Link      *Panic         // link to earlier panic
	Recovered bool           // whether this panic is over
	Aborted   bool           // the panic was aborted
}

type stack struct {
	lo uintptr
	hi uintptr
}

// v1.9.2
//noinspection ALL
type G struct {
	stack       stack
	stackguard0 uintptr
	stackguard1 uintptr

	Panic *Panic
}
