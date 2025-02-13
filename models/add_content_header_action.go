package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddContentHeaderAction 
type AddContentHeaderAction struct {
    InformationProtectionAction
    // The alignment property
    alignment *ContentAlignment
    // Color of the font to use for the header.
    fontColor *string
    // Name of the font to use for the header.
    fontName *string
    // Font size to use for the header.
    fontSize *int32
    // The margin of the header from the top of the document.
    margin *int32
    // The contents of the header itself.
    text *string
    // The name of the UI element where the header should be placed.
    uiElementName *string
}
// NewAddContentHeaderAction instantiates a new AddContentHeaderAction and sets the default values.
func NewAddContentHeaderAction()(*AddContentHeaderAction) {
    m := &AddContentHeaderAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    return m
}
// CreateAddContentHeaderActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAddContentHeaderActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAddContentHeaderAction(), nil
}
// GetAlignment gets the alignment property value. The alignment property
func (m *AddContentHeaderAction) GetAlignment()(*ContentAlignment) {
    if m == nil {
        return nil
    } else {
        return m.alignment
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddContentHeaderAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
    res["alignment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseContentAlignment)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlignment(val.(*ContentAlignment))
        }
        return nil
    }
    res["fontColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFontColor(val)
        }
        return nil
    }
    res["fontName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFontName(val)
        }
        return nil
    }
    res["fontSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFontSize(val)
        }
        return nil
    }
    res["margin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMargin(val)
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    res["uiElementName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUiElementName(val)
        }
        return nil
    }
    return res
}
// GetFontColor gets the fontColor property value. Color of the font to use for the header.
func (m *AddContentHeaderAction) GetFontColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fontColor
    }
}
// GetFontName gets the fontName property value. Name of the font to use for the header.
func (m *AddContentHeaderAction) GetFontName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fontName
    }
}
// GetFontSize gets the fontSize property value. Font size to use for the header.
func (m *AddContentHeaderAction) GetFontSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.fontSize
    }
}
// GetMargin gets the margin property value. The margin of the header from the top of the document.
func (m *AddContentHeaderAction) GetMargin()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.margin
    }
}
// GetText gets the text property value. The contents of the header itself.
func (m *AddContentHeaderAction) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// GetUiElementName gets the uiElementName property value. The name of the UI element where the header should be placed.
func (m *AddContentHeaderAction) GetUiElementName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uiElementName
    }
}
// Serialize serializes information the current object
func (m *AddContentHeaderAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InformationProtectionAction.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlignment() != nil {
        cast := (*m.GetAlignment()).String()
        err = writer.WriteStringValue("alignment", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fontColor", m.GetFontColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fontName", m.GetFontName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("fontSize", m.GetFontSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("margin", m.GetMargin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uiElementName", m.GetUiElementName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlignment sets the alignment property value. The alignment property
func (m *AddContentHeaderAction) SetAlignment(value *ContentAlignment)() {
    if m != nil {
        m.alignment = value
    }
}
// SetFontColor sets the fontColor property value. Color of the font to use for the header.
func (m *AddContentHeaderAction) SetFontColor(value *string)() {
    if m != nil {
        m.fontColor = value
    }
}
// SetFontName sets the fontName property value. Name of the font to use for the header.
func (m *AddContentHeaderAction) SetFontName(value *string)() {
    if m != nil {
        m.fontName = value
    }
}
// SetFontSize sets the fontSize property value. Font size to use for the header.
func (m *AddContentHeaderAction) SetFontSize(value *int32)() {
    if m != nil {
        m.fontSize = value
    }
}
// SetMargin sets the margin property value. The margin of the header from the top of the document.
func (m *AddContentHeaderAction) SetMargin(value *int32)() {
    if m != nil {
        m.margin = value
    }
}
// SetText sets the text property value. The contents of the header itself.
func (m *AddContentHeaderAction) SetText(value *string)() {
    if m != nil {
        m.text = value
    }
}
// SetUiElementName sets the uiElementName property value. The name of the UI element where the header should be placed.
func (m *AddContentHeaderAction) SetUiElementName(value *string)() {
    if m != nil {
        m.uiElementName = value
    }
}
