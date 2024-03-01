// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_ShapePath struct {
	RemarkAttr *string
}

func NewCT_ShapePath() *CT_ShapePath {
	return &CT_ShapePath{}
}

func (m *CT_ShapePath) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.RemarkAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "remark"},
			Value: fmt.Sprintf("%v", *m.RemarkAttr)})
	}
	e.EncodeToken(start)
	return nil
}

func (m *CT_ShapePath) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "remark" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RemarkAttr = &parsed
			break
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_ShapePath: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Shape_Path and its children
func (m *CT_ShapePath) Validate() error {
	return m.ValidateWithPath("CT_Shape_Path")
}

// ValidateWithPath validates the CT_Shape_Path and its children, prefixing error messages with path
func (m *CT_ShapePath) ValidateWithPath(path string) error {

	return nil
}
