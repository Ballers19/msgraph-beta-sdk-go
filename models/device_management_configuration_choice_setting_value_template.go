package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationChoiceSettingValueTemplate choice Setting Value Template
type DeviceManagementConfigurationChoiceSettingValueTemplate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Choice Setting Value Default Template.
    defaultValue DeviceManagementConfigurationChoiceSettingValueDefaultTemplateable
    // Recommended definition override.
    recommendedValueDefinition DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable
    // Required definition override.
    requiredValueDefinition DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable
    // Setting Value Template Id
    settingValueTemplateId *string
}
// NewDeviceManagementConfigurationChoiceSettingValueTemplate instantiates a new deviceManagementConfigurationChoiceSettingValueTemplate and sets the default values.
func NewDeviceManagementConfigurationChoiceSettingValueTemplate()(*DeviceManagementConfigurationChoiceSettingValueTemplate) {
    m := &DeviceManagementConfigurationChoiceSettingValueTemplate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationChoiceSettingValueTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationChoiceSettingValueTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationChoiceSettingValueTemplate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDefaultValue gets the defaultValue property value. Choice Setting Value Default Template.
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) GetDefaultValue()(DeviceManagementConfigurationChoiceSettingValueDefaultTemplateable) {
    if m == nil {
        return nil
    } else {
        return m.defaultValue
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationChoiceSettingValueDefaultTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val.(DeviceManagementConfigurationChoiceSettingValueDefaultTemplateable))
        }
        return nil
    }
    res["recommendedValueDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationChoiceSettingValueDefinitionTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendedValueDefinition(val.(DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable))
        }
        return nil
    }
    res["requiredValueDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationChoiceSettingValueDefinitionTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredValueDefinition(val.(DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable))
        }
        return nil
    }
    res["settingValueTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingValueTemplateId(val)
        }
        return nil
    }
    return res
}
// GetRecommendedValueDefinition gets the recommendedValueDefinition property value. Recommended definition override.
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) GetRecommendedValueDefinition()(DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable) {
    if m == nil {
        return nil
    } else {
        return m.recommendedValueDefinition
    }
}
// GetRequiredValueDefinition gets the requiredValueDefinition property value. Required definition override.
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) GetRequiredValueDefinition()(DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable) {
    if m == nil {
        return nil
    } else {
        return m.requiredValueDefinition
    }
}
// GetSettingValueTemplateId gets the settingValueTemplateId property value. Setting Value Template Id
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) GetSettingValueTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingValueTemplateId
    }
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("recommendedValueDefinition", m.GetRecommendedValueDefinition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("requiredValueDefinition", m.GetRequiredValueDefinition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingValueTemplateId", m.GetSettingValueTemplateId())
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
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDefaultValue sets the defaultValue property value. Choice Setting Value Default Template.
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) SetDefaultValue(value DeviceManagementConfigurationChoiceSettingValueDefaultTemplateable)() {
    if m != nil {
        m.defaultValue = value
    }
}
// SetRecommendedValueDefinition sets the recommendedValueDefinition property value. Recommended definition override.
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) SetRecommendedValueDefinition(value DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable)() {
    if m != nil {
        m.recommendedValueDefinition = value
    }
}
// SetRequiredValueDefinition sets the requiredValueDefinition property value. Required definition override.
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) SetRequiredValueDefinition(value DeviceManagementConfigurationChoiceSettingValueDefinitionTemplateable)() {
    if m != nil {
        m.requiredValueDefinition = value
    }
}
// SetSettingValueTemplateId sets the settingValueTemplateId property value. Setting Value Template Id
func (m *DeviceManagementConfigurationChoiceSettingValueTemplate) SetSettingValueTemplateId(value *string)() {
    if m != nil {
        m.settingValueTemplateId = value
    }
}
