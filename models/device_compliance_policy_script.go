package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceCompliancePolicyScript 
type DeviceCompliancePolicyScript struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Device compliance script Id.
    deviceComplianceScriptId *string
    // Json of the rules.
    rulesContent []byte
}
// NewDeviceCompliancePolicyScript instantiates a new deviceCompliancePolicyScript and sets the default values.
func NewDeviceCompliancePolicyScript()(*DeviceCompliancePolicyScript) {
    m := &DeviceCompliancePolicyScript{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceCompliancePolicyScriptFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePolicyScriptFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePolicyScript(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceCompliancePolicyScript) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceComplianceScriptId gets the deviceComplianceScriptId property value. Device compliance script Id.
func (m *DeviceCompliancePolicyScript) GetDeviceComplianceScriptId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceScriptId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePolicyScript) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceComplianceScriptId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceScriptId(val)
        }
        return nil
    }
    res["rulesContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRulesContent(val)
        }
        return nil
    }
    return res
}
// GetRulesContent gets the rulesContent property value. Json of the rules.
func (m *DeviceCompliancePolicyScript) GetRulesContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.rulesContent
    }
}
// Serialize serializes information the current object
func (m *DeviceCompliancePolicyScript) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceComplianceScriptId", m.GetDeviceComplianceScriptId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("rulesContent", m.GetRulesContent())
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
func (m *DeviceCompliancePolicyScript) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceComplianceScriptId sets the deviceComplianceScriptId property value. Device compliance script Id.
func (m *DeviceCompliancePolicyScript) SetDeviceComplianceScriptId(value *string)() {
    if m != nil {
        m.deviceComplianceScriptId = value
    }
}
// SetRulesContent sets the rulesContent property value. Json of the rules.
func (m *DeviceCompliancePolicyScript) SetRulesContent(value []byte)() {
    if m != nil {
        m.rulesContent = value
    }
}
