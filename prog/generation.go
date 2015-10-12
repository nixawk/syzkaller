// Copyright 2015 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package prog

import (
	"math/rand"

	"github.com/google/syzkaller/sys"
)

// Generate generates a random program of length ~ncalls.
// calls is a set of allowed syscalls, if nil all syscalls are used.
func Generate(rs rand.Source, ncalls int, enabledCalls []*sys.Call) *Prog {
	p := new(Prog)
	r := newRand(rs)
	s := newState(enabledCalls)
	for len(p.Calls) < ncalls {
		calls := r.generateCall(s)
		for _, c := range calls {
			s.analyze(c)
			p.Calls = append(p.Calls, c)
		}
	}
	if err := p.validate(); err != nil {
		panic(err)
	}
	return p
}
