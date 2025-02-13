package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSettingInstanceTemplateReference setting instance template reference information
type DeviceManagementConfigurationSettingInstanceTemplateReference struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Setting instance template id
    settingInstanceTemplateId *string
}
// NewDeviceManagementConfigurationSettingInstanceTemplateReference instantiates a new deviceManagementConfigurationSettingInstanceTemplateReference and sets the default values.
func NewDeviceManagementConfigurationSettingInstanceTemplateReference()(*DeviceManagementConfigurationSettingInstanceTemplateReference) {
    m := &DeviceManagementConfigurationSettingInstanceTemplateReference{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationSettingInstanceTemplateReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSettingInstanceTemplateReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSettingInstanceTemplateReference(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["settingInstanceTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingInstanceTemplateId(val)
        }
        return nil
    }
    return res
}
// GetSettingInstanceTemplateId gets the settingInstanceTemplateId property value. Setting instance template id
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) GetSettingInstanceTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingInstanceTemplateId
    }
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("settingInstanceTemplateId", m.GetSettingInstanceTemplateId())
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
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSettingInstanceTemplateId sets the settingInstanceTemplateId property value. Setting instance template id
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) SetSettingInstanceTemplateId(value *string)() {
    if m != nil {
        m.settingInstanceTemplateId = value
    }
}
