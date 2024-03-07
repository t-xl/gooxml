// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"

	"github.com/t-xl/gooxml"
	"github.com/t-xl/gooxml/schema/soo/dml"
)

type CT_DTable struct {
	ShowHorzBorder *CT_Boolean
	ShowVertBorder *CT_Boolean
	ShowOutline    *CT_Boolean
	ShowKeys       *CT_Boolean
	SpPr           *dml.CT_ShapeProperties
	TxPr           *dml.CT_TextBody
	ExtLst         *CT_ExtensionList
}

func NewCT_DTable() *CT_DTable {
	ret := &CT_DTable{}
	return ret
}

func (m *CT_DTable) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ShowHorzBorder != nil {
		seshowHorzBorder := xml.StartElement{Name: xml.Name{Local: "c:showHorzBorder"}}
		e.EncodeElement(m.ShowHorzBorder, seshowHorzBorder)
	}
	if m.ShowVertBorder != nil {
		seshowVertBorder := xml.StartElement{Name: xml.Name{Local: "c:showVertBorder"}}
		e.EncodeElement(m.ShowVertBorder, seshowVertBorder)
	}
	if m.ShowOutline != nil {
		seshowOutline := xml.StartElement{Name: xml.Name{Local: "c:showOutline"}}
		e.EncodeElement(m.ShowOutline, seshowOutline)
	}
	if m.ShowKeys != nil {
		seshowKeys := xml.StartElement{Name: xml.Name{Local: "c:showKeys"}}
		e.EncodeElement(m.ShowKeys, seshowKeys)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "c:spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.TxPr != nil {
		setxPr := xml.StartElement{Name: xml.Name{Local: "c:txPr"}}
		e.EncodeElement(m.TxPr, setxPr)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DTable) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DTable:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "showHorzBorder"}:
				m.ShowHorzBorder = NewCT_Boolean()
				if err := d.DecodeElement(m.ShowHorzBorder, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "showVertBorder"}:
				m.ShowVertBorder = NewCT_Boolean()
				if err := d.DecodeElement(m.ShowVertBorder, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "showOutline"}:
				m.ShowOutline = NewCT_Boolean()
				if err := d.DecodeElement(m.ShowOutline, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "showKeys"}:
				m.ShowKeys = NewCT_Boolean()
				if err := d.DecodeElement(m.ShowKeys, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "spPr"}:
				m.SpPr = dml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "txPr"}:
				m.TxPr = dml.NewCT_TextBody()
				if err := d.DecodeElement(m.TxPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_DTable %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DTable
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DTable and its children
func (m *CT_DTable) Validate() error {
	return m.ValidateWithPath("CT_DTable")
}

// ValidateWithPath validates the CT_DTable and its children, prefixing error messages with path
func (m *CT_DTable) ValidateWithPath(path string) error {
	if m.ShowHorzBorder != nil {
		if err := m.ShowHorzBorder.ValidateWithPath(path + "/ShowHorzBorder"); err != nil {
			return err
		}
	}
	if m.ShowVertBorder != nil {
		if err := m.ShowVertBorder.ValidateWithPath(path + "/ShowVertBorder"); err != nil {
			return err
		}
	}
	if m.ShowOutline != nil {
		if err := m.ShowOutline.ValidateWithPath(path + "/ShowOutline"); err != nil {
			return err
		}
	}
	if m.ShowKeys != nil {
		if err := m.ShowKeys.ValidateWithPath(path + "/ShowKeys"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	if m.TxPr != nil {
		if err := m.TxPr.ValidateWithPath(path + "/TxPr"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
