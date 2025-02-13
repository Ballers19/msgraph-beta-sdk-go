package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InformationProtectionAction 
type InformationProtectionAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The type property
    type_escaped *string
}
// NewInformationProtectionAction instantiates a new informationProtectionAction and sets the default values.
func NewInformationProtectionAction()(*InformationProtectionAction) {
    m := &InformationProtectionAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odatatypeValue := "#microsoft.graph.informationProtectionAction";
    m.SetType(&odatatypeValue);
    return m
}
// CreateInformationProtectionActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.addContentFooterAction":
                        return NewAddContentFooterAction(), nil
                    case "#microsoft.graph.addContentHeaderAction":
                        return NewAddContentHeaderAction(), nil
                    case "#microsoft.graph.addWatermarkAction":
                        return NewAddWatermarkAction(), nil
                    case "#microsoft.graph.applyLabelAction":
                        return NewApplyLabelAction(), nil
                    case "#microsoft.graph.customAction":
                        return NewCustomAction(), nil
                    case "#microsoft.graph.justifyAction":
                        return NewJustifyAction(), nil
                    case "#microsoft.graph.metadataAction":
                        return NewMetadataAction(), nil
                    case "#microsoft.graph.protectAdhocAction":
                        return NewProtectAdhocAction(), nil
                    case "#microsoft.graph.protectByTemplateAction":
                        return NewProtectByTemplateAction(), nil
                    case "#microsoft.graph.protectDoNotForwardAction":
                        return NewProtectDoNotForwardAction(), nil
                    case "#microsoft.graph.recommendLabelAction":
                        return NewRecommendLabelAction(), nil
                    case "#microsoft.graph.removeContentFooterAction":
                        return NewRemoveContentFooterAction(), nil
                    case "#microsoft.graph.removeContentHeaderAction":
                        return NewRemoveContentHeaderAction(), nil
                    case "#microsoft.graph.removeProtectionAction":
                        return NewRemoveProtectionAction(), nil
                    case "#microsoft.graph.removeWatermarkAction":
                        return NewRemoveWatermarkAction(), nil
                }
            }
        }
    }
    return NewInformationProtectionAction(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InformationProtectionAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtectionAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
func (m *InformationProtectionAction) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *InformationProtectionAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InformationProtectionAction) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetType sets the @odata.type property value. The type property
func (m *InformationProtectionAction) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
