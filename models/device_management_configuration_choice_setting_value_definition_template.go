package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationChoiceSettingValueDefinitionTemplate choice Setting Value Definition Template
type DeviceManagementConfigurationChoiceSettingValueDefinitionTemplate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Choice Setting Allowed Options
    allowedOptions []DeviceManagementConfigurationOptionDefinitionTemplateable
}
// NewDeviceManagementConfigurationChoiceSettingValueDefinitionTemplate instantiates a new deviceManagementConfigurationChoiceSettingValueDefinitionTemplate and sets the default values.
func NewDeviceManagementConfigurationChoiceSettingValueDefinitionTemplate()(*DeviceManagementConfigurationChoiceSettingValueDefinitionTemplate) {
    m := &DeviceManagementConfigurationChoiceSettingValueDefinitionTemplate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationChoiceSettingValueDefinitionTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationChoiceSettingValueDefinitionTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationChoiceSettingValueDefinitionTemplate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationChoiceSettingValueDefinitionTemplate) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowedOptions gets the allowedOptions property value. Choice Setting Allowed Options
func (m *DeviceManagementConfigurationChoiceSettingValueDefinitionTemplate) GetAllowedOptions()([]DeviceManagementConfigurationOptionDefinitionTemplateable) {
    if m == nil {
        return nil
    } else {
        return m.allowedOptions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationChoiceSettingValueDefinitionTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationOptionDefinitionTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationOptionDefinitionTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationOptionDefinitionTemplateable)
            }
            m.SetAllowedOptions(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationChoiceSettingValueDefinitionTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedOptions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAllowedOptions()))
        for i, v := range m.GetAllowedOptions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("allowedOptions", cast)
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
func (m *DeviceManagementConfigurationChoiceSettingValueDefinitionTemplate) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowedOptions sets the allowedOptions property value. Choice Setting Allowed Options
func (m *DeviceManagementConfigurationChoiceSettingValueDefinitionTemplate) SetAllowedOptions(value []DeviceManagementConfigurationOptionDefinitionTemplateable)() {
    if m != nil {
        m.allowedOptions = value
    }
}
