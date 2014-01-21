// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gexf

import (
	"testing"
)

func Test_sample(t *testing.T) {
	g, err := Open("sample.gexf")
	checkError(err)

	//g := New(Directed)
	//g.AddNode(Node{"0", "a"})
	//g.AddNode(Node{"1", "b"})
	//g.AddEdge(Edge{"0", "0", "1", 1.0})

	p(g)
	checkError(g.Save("/dev/stdout"))
}
