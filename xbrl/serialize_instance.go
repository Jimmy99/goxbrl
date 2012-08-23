package xbrl

import (
	"encoding/xml"
	"github.com/zeebo/goxbrl/marshal"
)

func serializeInstance(f Filing) *marshal.Node {
	return &marshal.Node{
		Name:  xml.Name{"", ""},
		Attrs: []xml.Attr{},
		Nodes: []*marshal.Node{},
		Value: "",
	}
}
