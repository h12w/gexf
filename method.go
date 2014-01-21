// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gexf

import (
	"encoding/xml"
	"os"
	"time"
)

func Open(file string) (g *Gexf, err error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return g, xml.NewDecoder(f).Decode(&g)
}

func New() *Gexf {
	return &Gexf{
		XMLName: xml.Name{Namespace, "gexf"},
		Version: Version,
		Meta: Meta{
			LastModified: time.Now().Format(TimeFormat),
			Creator:      DefaultCreator,
		},
		Graph: Graph{
			EdgeType: DefaultEdgeType,
			IdType:   DefaultIdType,
			Type:     DefaultGraphType,
		},
	}
}

func (g *Gexf) AddNode(n Node) {
	g.Graph.Nodes.Add(n)
}

func (g *Gexf) AddEdge(e Edge) {
	g.Graph.Edges.Add(e)
}

func (g *Gexf) Save(file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write([]byte(xml.Header)); err != nil {
		return err
	}
	return xml.NewEncoder(f).Encode(g)
}

func (t EdgeType) String() string {
	switch t {
	case Directed:
		return "directed"
	case Undirected:
		return "undirected"
	case Mutual:
		return "mutual"
	}
	return ""
}

func (ns *Nodes) Add(n Node) {
	ns.Node = append(ns.Node, n)
	ns.Count++
}

func (es *Edges) Add(e Edge) {
	es.Edge = append(es.Edge, e)
	es.Count++
}
