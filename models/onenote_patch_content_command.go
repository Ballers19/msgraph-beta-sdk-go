package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnenotePatchContentCommand 
type OnenotePatchContentCommand struct {
    // The action property
    action *OnenotePatchActionType
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A string of well-formed HTML to add to the page, and any image or file binary data. If the content contains binary data, the request must be sent using the multipart/form-data content type with a 'Commands' part.
    content *string
    // The location to add the supplied content, relative to the target element. Possible values are: after (default) or before.
    position *OnenotePatchInsertPosition
    // The element to update. Must be the #<data-id> or the generated {id} of the element, or the body or title keyword.
    target *string
}
// NewOnenotePatchContentCommand instantiates a new onenotePatchContentCommand and sets the default values.
func NewOnenotePatchContentCommand()(*OnenotePatchContentCommand) {
    m := &OnenotePatchContentCommand{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnenotePatchContentCommandFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenotePatchContentCommandFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnenotePatchContentCommand(), nil
}
// GetAction gets the action property value. The action property
func (m *OnenotePatchContentCommand) GetAction()(*OnenotePatchActionType) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnenotePatchContentCommand) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContent gets the content property value. A string of well-formed HTML to add to the page, and any image or file binary data. If the content contains binary data, the request must be sent using the multipart/form-data content type with a 'Commands' part.
func (m *OnenotePatchContentCommand) GetContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenotePatchContentCommand) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnenotePatchActionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*OnenotePatchActionType))
        }
        return nil
    }
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["position"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnenotePatchInsertPosition)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosition(val.(*OnenotePatchInsertPosition))
        }
        return nil
    }
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val)
        }
        return nil
    }
    return res
}
// GetPosition gets the position property value. The location to add the supplied content, relative to the target element. Possible values are: after (default) or before.
func (m *OnenotePatchContentCommand) GetPosition()(*OnenotePatchInsertPosition) {
    if m == nil {
        return nil
    } else {
        return m.position
    }
}
// GetTarget gets the target property value. The element to update. Must be the #<data-id> or the generated {id} of the element, or the body or title keyword.
func (m *OnenotePatchContentCommand) GetTarget()(*string) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// Serialize serializes information the current object
func (m *OnenotePatchContentCommand) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err := writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    if m.GetPosition() != nil {
        cast := (*m.GetPosition()).String()
        err := writer.WriteStringValue("position", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. The action property
func (m *OnenotePatchContentCommand) SetAction(value *OnenotePatchActionType)() {
    if m != nil {
        m.action = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnenotePatchContentCommand) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContent sets the content property value. A string of well-formed HTML to add to the page, and any image or file binary data. If the content contains binary data, the request must be sent using the multipart/form-data content type with a 'Commands' part.
func (m *OnenotePatchContentCommand) SetContent(value *string)() {
    if m != nil {
        m.content = value
    }
}
// SetPosition sets the position property value. The location to add the supplied content, relative to the target element. Possible values are: after (default) or before.
func (m *OnenotePatchContentCommand) SetPosition(value *OnenotePatchInsertPosition)() {
    if m != nil {
        m.position = value
    }
}
// SetTarget sets the target property value. The element to update. Must be the #<data-id> or the generated {id} of the element, or the body or title keyword.
func (m *OnenotePatchContentCommand) SetTarget(value *string)() {
    if m != nil {
        m.target = value
    }
}
