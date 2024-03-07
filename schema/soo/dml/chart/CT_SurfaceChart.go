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
	"fmt"

	"github.com/t-xl/gooxml"
)

type CT_SurfaceChart struct {
	Wireframe *CT_Boolean
	Ser       []*CT_SurfaceSer
	BandFmts  *CT_BandFmts
	AxId      []*CT_UnsignedInt
	ExtLst    *CT_ExtensionList
}

func NewCT_SurfaceChart() *CT_SurfaceChart {
	ret := &CT_SurfaceChart{}
	return ret
}

func (m *CT_SurfaceChart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Wireframe != nil {
		sewireframe := xml.StartElement{Name: xml.Name{Local: "c:wireframe"}}
		e.EncodeElement(m.Wireframe, sewireframe)
	}
	if m.Ser != nil {
		seser := xml.StartElement{Name: xml.Name{Local: "c:ser"}}
		for _, c := range m.Ser {
			e.EncodeElement(c, seser)
		}
	}
	if m.BandFmts != nil {
		sebandFmts := xml.StartElement{Name: xml.Name{Local: "c:bandFmts"}}
		e.EncodeElement(m.BandFmts, sebandFmts)
	}
	seaxId := xml.StartElement{Name: xml.Name{Local: "c:axId"}}
	for _, c := range m.AxId {
		e.EncodeElement(c, seaxId)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SurfaceChart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SurfaceChart:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "wireframe"}:
				m.Wireframe = NewCT_Boolean()
				if err := d.DecodeElement(m.Wireframe, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "ser"}:
				tmp := NewCT_SurfaceSer()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ser = append(m.Ser, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "bandFmts"}:
				m.BandFmts = NewCT_BandFmts()
				if err := d.DecodeElement(m.BandFmts, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "axId"}:
				tmp := NewCT_UnsignedInt()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AxId = append(m.AxId, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_SurfaceChart %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SurfaceChart
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SurfaceChart and its children
func (m *CT_SurfaceChart) Validate() error {
	return m.ValidateWithPath("CT_SurfaceChart")
}

// ValidateWithPath validates the CT_SurfaceChart and its children, prefixing error messages with path
func (m *CT_SurfaceChart) ValidateWithPath(path string) error {
	if m.Wireframe != nil {
		if err := m.Wireframe.ValidateWithPath(path + "/Wireframe"); err != nil {
			return err
		}
	}
	for i, v := range m.Ser {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ser[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.BandFmts != nil {
		if err := m.BandFmts.ValidateWithPath(path + "/BandFmts"); err != nil {
			return err
		}
	}
	for i, v := range m.AxId {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AxId[%d]", path, i)); err != nil {
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
