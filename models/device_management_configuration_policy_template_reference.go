package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationPolicyTemplateReference policy template reference information
type DeviceManagementConfigurationPolicyTemplateReference struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Template Display Name of the referenced template. This property is read-only.
    templateDisplayName *string
    // Template Display Version of the referenced Template. This property is read-only.
    templateDisplayVersion *string
    // Describes the TemplateFamily for the Template entity
    templateFamily *DeviceManagementConfigurationTemplateFamily
    // Template id
    templateId *string
}
// NewDeviceManagementConfigurationPolicyTemplateReference instantiates a new deviceManagementConfigurationPolicyTemplateReference and sets the default values.
func NewDeviceManagementConfigurationPolicyTemplateReference()(*DeviceManagementConfigurationPolicyTemplateReference) {
    m := &DeviceManagementConfigurationPolicyTemplateReference{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationPolicyTemplateReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationPolicyTemplateReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationPolicyTemplateReference(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["templateDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateDisplayName(val)
        }
        return nil
    }
    res["templateDisplayVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateDisplayVersion(val)
        }
        return nil
    }
    res["templateFamily"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationTemplateFamily)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateFamily(val.(*DeviceManagementConfigurationTemplateFamily))
        }
        return nil
    }
    res["templateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    return res
}
// GetTemplateDisplayName gets the templateDisplayName property value. Template Display Name of the referenced template. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetTemplateDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateDisplayName
    }
}
// GetTemplateDisplayVersion gets the templateDisplayVersion property value. Template Display Version of the referenced Template. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetTemplateDisplayVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateDisplayVersion
    }
}
// GetTemplateFamily gets the templateFamily property value. Describes the TemplateFamily for the Template entity
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetTemplateFamily()(*DeviceManagementConfigurationTemplateFamily) {
    if m == nil {
        return nil
    } else {
        return m.templateFamily
    }
}
// GetTemplateId gets the templateId property value. Template id
func (m *DeviceManagementConfigurationPolicyTemplateReference) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationPolicyTemplateReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("templateDisplayName", m.GetTemplateDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("templateDisplayVersion", m.GetTemplateDisplayVersion())
        if err != nil {
            return err
        }
    }
    if m.GetTemplateFamily() != nil {
        cast := (*m.GetTemplateFamily()).String()
        err := writer.WriteStringValue("templateFamily", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("templateId", m.GetTemplateId())
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
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTemplateDisplayName sets the templateDisplayName property value. Template Display Name of the referenced template. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetTemplateDisplayName(value *string)() {
    if m != nil {
        m.templateDisplayName = value
    }
}
// SetTemplateDisplayVersion sets the templateDisplayVersion property value. Template Display Version of the referenced Template. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetTemplateDisplayVersion(value *string)() {
    if m != nil {
        m.templateDisplayVersion = value
    }
}
// SetTemplateFamily sets the templateFamily property value. Describes the TemplateFamily for the Template entity
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetTemplateFamily(value *DeviceManagementConfigurationTemplateFamily)() {
    if m != nil {
        m.templateFamily = value
    }
}
// SetTemplateId sets the templateId property value. Template id
func (m *DeviceManagementConfigurationPolicyTemplateReference) SetTemplateId(value *string)() {
    if m != nil {
        m.templateId = value
    }
}
