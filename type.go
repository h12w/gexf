// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gexf

import (
	"encoding/xml"
)

const (
	Version    = "1.2"
	Namespace  = "http://www.gexf.net/1.2draft"
	TimeFormat = "2006-01-02+15:04"

	DefaultCreator = "Gexf Go Library " + Version
)

type Gexf struct {
	XMLName xml.Name
	Version string `xml:"version,attr"`
	Meta    Meta   `xml:"meta"`

	Graph Graph `xml:"graph"`
}

type Graph struct {
	Type     GraphType `xml:"type,attr"`
	IdType   IdType    `xml:"idtype,attr"`
	EdgeType EdgeType  `xml:"defaultedgetype,attr"`

	Nodes Nodes `xml:"nodes"`
	Edges Edges `xml:"edges"`
}

type EdgeType string

const (
	Directed        = EdgeType("directed")
	Undirected      = EdgeType("undirected")
	Mutual          = EdgeType("mutual")
	DefaultEdgeType = Directed
)

type IdType string

const (
	String        = IdType("string")
	Integer       = IdType("integer")
	DefaultIdType = String
)

type GraphType string

const (
	Static           = GraphType("static")
	Dynamic          = GraphType("dynamic")
	DefaultGraphType = Static
)

type Meta struct {
	LastModified string `xml:"lastmodifieddate,attr"`
	Creator      string `xml:"creator,omitempty"`
	Keywords     string `xml:"keywords,omitempty"`
	Description  string `xml:"description,omitempty"`
}

type Nodes struct {
	Node  []Node `xml:"node"`
	Count int    `xml:"count,attr"`
}

type Node struct {
	Id    string `xml:"id,attr"`
	Label string `xml:"label,attr"`
}

type Edges struct {
	Edge  []Edge `xml:"edge"`
	Count int    `xml:"count,attr"`
}

type Edge struct {
	Id     string  `xml:"id,attr"`
	Source string  `xml:"source,attr"`
	Target string  `xml:"target,attr"`
	Weight float64 `xml:"weight,attr"`
}
