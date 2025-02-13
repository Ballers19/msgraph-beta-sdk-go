package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppConfigurationSettingItem contains properties for App configuration setting item.
type AppConfigurationSettingItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // app configuration key.
    appConfigKey *string
    // App configuration key types.
    appConfigKeyType *MdmAppConfigKeyType
    // app configuration key value.
    appConfigKeyValue *string
}
// NewAppConfigurationSettingItem instantiates a new appConfigurationSettingItem and sets the default values.
func NewAppConfigurationSettingItem()(*AppConfigurationSettingItem) {
    m := &AppConfigurationSettingItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppConfigurationSettingItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppConfigurationSettingItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppConfigurationSettingItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppConfigurationSettingItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppConfigKey gets the appConfigKey property value. app configuration key.
func (m *AppConfigurationSettingItem) GetAppConfigKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appConfigKey
    }
}
// GetAppConfigKeyType gets the appConfigKeyType property value. App configuration key types.
func (m *AppConfigurationSettingItem) GetAppConfigKeyType()(*MdmAppConfigKeyType) {
    if m == nil {
        return nil
    } else {
        return m.appConfigKeyType
    }
}
// GetAppConfigKeyValue gets the appConfigKeyValue property value. app configuration key value.
func (m *AppConfigurationSettingItem) GetAppConfigKeyValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appConfigKeyValue
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppConfigurationSettingItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appConfigKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppConfigKey(val)
        }
        return nil
    }
    res["appConfigKeyType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMdmAppConfigKeyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppConfigKeyType(val.(*MdmAppConfigKeyType))
        }
        return nil
    }
    res["appConfigKeyValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppConfigKeyValue(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AppConfigurationSettingItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appConfigKey", m.GetAppConfigKey())
        if err != nil {
            return err
        }
    }
    if m.GetAppConfigKeyType() != nil {
        cast := (*m.GetAppConfigKeyType()).String()
        err := writer.WriteStringValue("appConfigKeyType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("appConfigKeyValue", m.GetAppConfigKeyValue())
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
func (m *AppConfigurationSettingItem) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppConfigKey sets the appConfigKey property value. app configuration key.
func (m *AppConfigurationSettingItem) SetAppConfigKey(value *string)() {
    if m != nil {
        m.appConfigKey = value
    }
}
// SetAppConfigKeyType sets the appConfigKeyType property value. App configuration key types.
func (m *AppConfigurationSettingItem) SetAppConfigKeyType(value *MdmAppConfigKeyType)() {
    if m != nil {
        m.appConfigKeyType = value
    }
}
// SetAppConfigKeyValue sets the appConfigKeyValue property value. app configuration key value.
func (m *AppConfigurationSettingItem) SetAppConfigKeyValue(value *string)() {
    if m != nil {
        m.appConfigKeyValue = value
    }
}
