// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"

	"github.com/t-xl/gooxml"
	"github.com/t-xl/gooxml/schema/soo/dml"
)

type CT_CTStyleLabel struct {
	NameAttr       string
	FillClrLst     *CT_Colors
	LinClrLst      *CT_Colors
	EffectClrLst   *CT_Colors
	TxLinClrLst    *CT_Colors
	TxFillClrLst   *CT_Colors
	TxEffectClrLst *CT_Colors
	ExtLst         *dml.CT_OfficeArtExtensionList
}

func NewCT_CTStyleLabel() *CT_CTStyleLabel {
	ret := &CT_CTStyleLabel{}
	return ret
}

func (m *CT_CTStyleLabel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	e.EncodeToken(start)
	if m.FillClrLst != nil {
		sefillClrLst := xml.StartElement{Name: xml.Name{Local: "fillClrLst"}}
		e.EncodeElement(m.FillClrLst, sefillClrLst)
	}
	if m.LinClrLst != nil {
		selinClrLst := xml.StartElement{Name: xml.Name{Local: "linClrLst"}}
		e.EncodeElement(m.LinClrLst, selinClrLst)
	}
	if m.EffectClrLst != nil {
		seeffectClrLst := xml.StartElement{Name: xml.Name{Local: "effectClrLst"}}
		e.EncodeElement(m.EffectClrLst, seeffectClrLst)
	}
	if m.TxLinClrLst != nil {
		setxLinClrLst := xml.StartElement{Name: xml.Name{Local: "txLinClrLst"}}
		e.EncodeElement(m.TxLinClrLst, setxLinClrLst)
	}
	if m.TxFillClrLst != nil {
		setxFillClrLst := xml.StartElement{Name: xml.Name{Local: "txFillClrLst"}}
		e.EncodeElement(m.TxFillClrLst, setxFillClrLst)
	}
	if m.TxEffectClrLst != nil {
		setxEffectClrLst := xml.StartElement{Name: xml.Name{Local: "txEffectClrLst"}}
		e.EncodeElement(m.TxEffectClrLst, setxEffectClrLst)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CTStyleLabel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
	}
lCT_CTStyleLabel:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "fillClrLst"}:
				m.FillClrLst = NewCT_Colors()
				if err := d.DecodeElement(m.FillClrLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "linClrLst"}:
				m.LinClrLst = NewCT_Colors()
				if err := d.DecodeElement(m.LinClrLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "effectClrLst"}:
				m.EffectClrLst = NewCT_Colors()
				if err := d.DecodeElement(m.EffectClrLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "txLinClrLst"}:
				m.TxLinClrLst = NewCT_Colors()
				if err := d.DecodeElement(m.TxLinClrLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "txFillClrLst"}:
				m.TxFillClrLst = NewCT_Colors()
				if err := d.DecodeElement(m.TxFillClrLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "txEffectClrLst"}:
				m.TxEffectClrLst = NewCT_Colors()
				if err := d.DecodeElement(m.TxEffectClrLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "extLst"}:
				m.ExtLst = dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_CTStyleLabel %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CTStyleLabel
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CTStyleLabel and its children
func (m *CT_CTStyleLabel) Validate() error {
	return m.ValidateWithPath("CT_CTStyleLabel")
}

// ValidateWithPath validates the CT_CTStyleLabel and its children, prefixing error messages with path
func (m *CT_CTStyleLabel) ValidateWithPath(path string) error {
	if m.FillClrLst != nil {
		if err := m.FillClrLst.ValidateWithPath(path + "/FillClrLst"); err != nil {
			return err
		}
	}
	if m.LinClrLst != nil {
		if err := m.LinClrLst.ValidateWithPath(path + "/LinClrLst"); err != nil {
			return err
		}
	}
	if m.EffectClrLst != nil {
		if err := m.EffectClrLst.ValidateWithPath(path + "/EffectClrLst"); err != nil {
			return err
		}
	}
	if m.TxLinClrLst != nil {
		if err := m.TxLinClrLst.ValidateWithPath(path + "/TxLinClrLst"); err != nil {
			return err
		}
	}
	if m.TxFillClrLst != nil {
		if err := m.TxFillClrLst.ValidateWithPath(path + "/TxFillClrLst"); err != nil {
			return err
		}
	}
	if m.TxEffectClrLst != nil {
		if err := m.TxEffectClrLst.ValidateWithPath(path + "/TxEffectClrLst"); err != nil {
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
