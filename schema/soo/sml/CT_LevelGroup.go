// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/t-xl/gooxml"
)

type CT_LevelGroup struct {
	// Group Name
	NameAttr string
	// Unique Group Name
	UniqueNameAttr string
	// Group Caption
	CaptionAttr string
	// Parent Unique Name
	UniqueParentAttr *string
	// Group Id
	IdAttr *int32
	// OLAP Group Members
	GroupMembers *CT_GroupMembers
}

func NewCT_LevelGroup() *CT_LevelGroup {
	ret := &CT_LevelGroup{}
	ret.GroupMembers = NewCT_GroupMembers()
	return ret
}

func (m *CT_LevelGroup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uniqueName"},
		Value: fmt.Sprintf("%v", m.UniqueNameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "caption"},
		Value: fmt.Sprintf("%v", m.CaptionAttr)})
	if m.UniqueParentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uniqueParent"},
			Value: fmt.Sprintf("%v", *m.UniqueParentAttr)})
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	e.EncodeToken(start)
	segroupMembers := xml.StartElement{Name: xml.Name{Local: "ma:groupMembers"}}
	e.EncodeElement(m.GroupMembers, segroupMembers)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LevelGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.GroupMembers = NewCT_GroupMembers()
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
		if attr.Name.Local == "uniqueName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UniqueNameAttr = parsed
			continue
		}
		if attr.Name.Local == "caption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CaptionAttr = parsed
			continue
		}
		if attr.Name.Local == "uniqueParent" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UniqueParentAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.IdAttr = &pt
			continue
		}
	}
lCT_LevelGroup:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "groupMembers"}:
				if err := d.DecodeElement(m.GroupMembers, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_LevelGroup %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_LevelGroup
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_LevelGroup and its children
func (m *CT_LevelGroup) Validate() error {
	return m.ValidateWithPath("CT_LevelGroup")
}

// ValidateWithPath validates the CT_LevelGroup and its children, prefixing error messages with path
func (m *CT_LevelGroup) ValidateWithPath(path string) error {
	if err := m.GroupMembers.ValidateWithPath(path + "/GroupMembers"); err != nil {
		return err
	}
	return nil
}
