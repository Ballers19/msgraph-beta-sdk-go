package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DlpActionInfo 
type DlpActionInfo struct {
    // The action property
    action *DlpAction
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The type property
    type_escaped *string
}
// NewDlpActionInfo instantiates a new dlpActionInfo and sets the default values.
func NewDlpActionInfo()(*DlpActionInfo) {
    m := &DlpActionInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odatatypeValue := "#microsoft.graph.dlpActionInfo";
    m.SetType(&odatatypeValue);
    return m
}
// CreateDlpActionInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDlpActionInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.blockAccessAction":
                        return NewBlockAccessAction(), nil
                    case "#microsoft.graph.deviceRestrictionAction":
                        return NewDeviceRestrictionAction(), nil
                    case "#microsoft.graph.notifyUserAction":
                        return NewNotifyUserAction(), nil
                }
            }
        }
    }
    return NewDlpActionInfo(), nil
}
// GetAction gets the action property value. The action property
func (m *DlpActionInfo) GetAction()(*DlpAction) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DlpActionInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DlpActionInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDlpAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*DlpAction))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetType gets the @odata.type property value. The type property
func (m *DlpActionInfo) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *DlpActionInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err := writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetType())
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
func (m *DlpActionInfo) SetAction(value *DlpAction)() {
    if m != nil {
        m.action = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DlpActionInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetType sets the @odata.type property value. The type property
func (m *DlpActionInfo) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
